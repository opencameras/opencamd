package media

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Ffmpeg struct {
	Device string
	InputFrameRate float32
	OutputWidth uint
	OutputHeight uint
}

func NewFfmpeg(dev string, framerate float32, width, height uint) *Ffmpeg {
	return &Ffmpeg{Device: dev, InputFrameRate: framerate, OutputWidth: width, OutputHeight: height}
}

func (f *Ffmpeg) Start(ctx context.Context, dstAddr string){
	framerate := fmt.Sprintf("%f", f.InputFrameRate)
	cmd := exec.CommandContext(ctx, "ffmpeg",
		"-f", f.getInputFormat(),
		"-framerate", framerate,
		"-i", f.Device,
		"-r", framerate,
		"-bsf:v", "h264_mp4toannexb",
		"-vf", fmt.Sprintf("scale=%dx%d,format=yuv420p", f.OutputWidth, f.OutputHeight),
		//"-c:v", "h264",
		"-c:v", "libx264", "-intra", // test infra mode for sync
		"-f", "h264",
		dstAddr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func(cmd *exec.Cmd) {
		err := cmd.Run()
		if err != nil {
			fmt.Printf("ffmpeg run: %s", err)
		}
	} (cmd)
}