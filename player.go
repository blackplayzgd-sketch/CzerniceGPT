package main

import (
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

var (
	sr beep.SampleRate = 48000
)

func playSongAndWaitForItToFinish(path string) time.Duration {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	speaker.Play(resampled)
	length := streamer.Len()

	return time.Duration(float64(length) / float64(format.SampleRate) * float64(time.Second))

}
