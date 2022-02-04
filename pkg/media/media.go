package media

import (
	"context"
	"errors"
	"fmt"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"io"
	"time"
)

const (
	oggPageDuration   = time.Millisecond * 20
	h264FrameDuration = time.Millisecond * 33
)


type WebrtcSession struct {
	VideoSource Datasource
	AudioSource Datasource
	StunServers []string
}

func (ws *WebrtcSession) Start(offer webrtc.SessionDescription, closeChan chan struct{}) (*webrtc.SessionDescription, error) {
	// Assert that we have an audio or video file
	var haveVideoFile, haveAudioFile bool
	var videoSourceName, audioSourceName string
	if ws.VideoSource == nil {
		haveVideoFile = false
	} else {
		vcodecs, err := ws.VideoSource.GetCodecs()
		if err != nil {
			return nil, err
		}
		haveVideoFile = vcodecs != nil
		videoSourceName = ws.VideoSource.String()
	}

	if ws.AudioSource == nil {
		haveAudioFile = false
	} else {
		acodecs, err := ws.AudioSource.GetCodecs()
		if err != nil {
			return nil, err
		}
		haveAudioFile = acodecs != nil
		audioSourceName = ws.AudioSource.String()
	}

	if !haveAudioFile && !haveVideoFile {
		return nil, errors.New("Could not find `" + videoSourceName + "` or `" + audioSourceName + "`")
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{{URLs: ws.StunServers,},
		},
	})
	if err != nil {
		return nil, err
	}
	go func() {
		<- closeChan
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()

	iceConnectedCtx, iceConnectedCtxCancel := context.WithCancel(context.Background())

	if haveVideoFile {
		err = ws.addVideoTrack(iceConnectedCtx, peerConnection)
		if err != nil {
			return nil, err
		}
	}

	if haveAudioFile {
		err = ws.addAudioTrack(iceConnectedCtx, peerConnection)
		if err != nil {
			return nil, err
		}
	}

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
		if connectionState == webrtc.ICEConnectionStateConnected {
			iceConnectedCtxCancel()
		}
	})

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			fmt.Println("Peer Connection has gone to failed exiting")
		}
	})

	// Set the remote SessionDescription
	if err = peerConnection.SetRemoteDescription(offer); err != nil {
		return nil, err
	}

	// Create answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		return nil, err
	}

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	if err = peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	return peerConnection.LocalDescription(), nil
}

func (ws *WebrtcSession) addVideoTrack(iceConnectedCtx context.Context, peerConnection *webrtc.PeerConnection) error {
	// Create a video track
	videoTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeH264}, "video", "pion")
	if err != nil {
		return err
	}

	rtpSender, err := peerConnection.AddTrack(videoTrack)
	if err != nil {
		return err
	}

	// Read incoming RTCP packets
	// Before these packets are returned they are processed by interceptors. For things
	// like NACK this needs to be called.
	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	go func() {
		// Open a H264 file and start reading using our IVFReader
		h264, h264Err := ws.VideoSource.OpenH264()
		if h264Err != nil {
			fmt.Printf("open h264 video stream: %s\n", h264Err)
			return
		}

		// Wait for connection established
		<-iceConnectedCtx.Done()

		// Send our video file frame at a time. Pace our sending so we send it at the same speed it should be played back as.
		// This isn't required since the video is timestamped, but we will such much higher loss if we send all at once.
		//
		// It is important to use a time.Ticker instead of time.Sleep because
		// * avoids accumulating skew, just calling time.Sleep didn't compensate for the time spent parsing the data
		// * works around latency issues with Sleep (see https://github.com/golang/go/issues/44343)
		ticker := time.NewTicker(h264FrameDuration)
		for ; true; <-ticker.C {
			nal, h264Err := h264.NextNAL()
			if h264Err == io.EOF {
				fmt.Println("All video frames parsed and sent")
				return
			}
			if h264Err != nil {
				fmt.Printf("read h264 video stream: %s\n", h264Err)
				return
			}

			if h264Err = videoTrack.WriteSample(media.Sample{Data: nal.Data, Duration: time.Second}); h264Err != nil {
				fmt.Printf("write h264 video stream: %s\n", h264Err)
				return
			}
		}
	}()
	return nil
}

func (ws *WebrtcSession) addAudioTrack(iceConnectedCtx context.Context, peerConnection *webrtc.PeerConnection) error {
	// Create a audio track
	audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, "audio", "pion")
	if err != nil {
		return err
	}

	rtpSender, err := peerConnection.AddTrack(audioTrack)
	if err != nil {
		return err
	}

	// Read incoming RTCP packets
	// Before these packets are returned they are processed by interceptors. For things
	// like NACK this needs to be called.
	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	go func() {
		// Open on oggfile in non-checksum mode.
		ogg, oggErr := ws.AudioSource.OpenOgg()
		if oggErr != nil {
			fmt.Printf("open ogg audio stream: %s\n", oggErr)
			return
		}

		// Wait for connection established
		<-iceConnectedCtx.Done()

		// Keep track of last granule, the difference is the amount of samples in the buffer
		var lastGranule uint64

		// It is important to use a time.Ticker instead of time.Sleep because
		// * avoids accumulating skew, just calling time.Sleep didn't compensate for the time spent parsing the data
		// * works around latency issues with Sleep (see https://github.com/golang/go/issues/44343)
		ticker := time.NewTicker(oggPageDuration)
		for ; true; <-ticker.C {
			pageData, pageHeader, oggErr := ogg.ParseNextPage()
			if oggErr == io.EOF {
				fmt.Println("All audio pages parsed and sent")
				return
			}

			if oggErr != nil {
				fmt.Printf("parse ogg audio stream: %s\n", oggErr)
				return
			}

			// The amount of samples is the difference between the last and current timestamp
			sampleCount := float64(pageHeader.GranulePosition - lastGranule)
			lastGranule = pageHeader.GranulePosition
			sampleDuration := time.Duration((sampleCount/48000)*1000) * time.Millisecond

			if oggErr = audioTrack.WriteSample(media.Sample{Data: pageData, Duration: sampleDuration}); oggErr != nil {
				fmt.Printf("write ogg audio stream: %s\n", oggErr)
				return
			}
		}
	}()
	return nil
}