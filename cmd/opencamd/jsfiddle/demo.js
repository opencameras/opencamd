/* eslint-env browser */

let pc = new RTCPeerConnection({
  iceServers: [
    {
      urls: 'stun:stun.l.google.com:19302'
    }
  ]
})
let log = msg => {
  document.getElementById('div').innerHTML += '<br>' + msg + '<br>'
}

pc.ontrack = function (event) {
  var el = document.createElement(event.track.kind)
  el.srcObject = event.streams[0]
  el.autoplay = true
  el.controls = true

  document.getElementById('remoteVideos').appendChild(el)
}

pc.oniceconnectionstatechange = e => log(pc.iceConnectionState)
pc.onicecandidate = event => {
  if (event.candidate === null) {
    let data = JSON.stringify(pc.localDescription, undefined, 2)
    let sdpOrig = document.getElementById('localSessionDescriptionOrig');
    if (sdpOrig) {
      sdpOrig.value = data
      sdpOrig.style.height = 'auto';
      sdpOrig.style.height = sdpOrig.scrollHeight + 'px';
    }
    let sdp = document.getElementById('localSessionDescription');
    if (sdp) {
      sdp.value = data.replaceAll('\\r\\n', `
`)
      sdp.style.height = 'auto';
      sdp.style.height = sdp.scrollHeight + 'px';
    }

    document.getElementById('start').disabled = false;

    log("Browser side is ready, Click 'Start Session' button to view camera video")
  }
}

// Offer to receive 1 audio, and 1 video track
pc.addTransceiver('video', {'direction': 'sendrecv'})
pc.addTransceiver('audio', {'direction': 'sendrecv'})

pc.createOffer().then(d => pc.setLocalDescription(d)).catch(log)

window.startSession = () => {
  fetch("http://localhost:8000/api/v1/media/live/session", {
    method: "POST",
    headers: {
      'Access-Control-Request-Method': 'POST',
      'Access-Control-Allow-Origin': '*',
      'Content-Type': 'application/json',
      'x-token': '123456'},
    body: document.getElementById('localSessionDescriptionOrig').value
  })
      .then(response => response.json())
      .then(data => {
        let sdp = document.getElementById('remoteSessionDescription');
        if (sdp) {
          sdp.value = JSON.stringify(data, undefined, 2).replaceAll('\\r\\n',`
`)
          sdp.style.height = 'auto';
          sdp.style.height = sdp.scrollHeight + 'px';
        }

        try {
          pc.setRemoteDescription(new RTCSessionDescription(data))
        } catch (e) {
          alert(e)
        }
      });
}
