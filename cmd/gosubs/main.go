package main

import (
	"fmt"
	"os"

	"github.com/nightnoryu/gosubs/pkg/gosubs/app"
	"github.com/nightnoryu/gosubs/pkg/gosubs/infrastructure/ffmpeg"
	"github.com/nightnoryu/gosubs/pkg/gosubs/infrastructure/whisper"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	args, err := parseArguments()
	if err != nil {
		return err
	}

	transcribingService, err := buildTranscribingService(args)
	if err != nil {
		return err
	}

	return transcribingService.TranscribeVideo(args.inputFilename, args.outputFilename)
}

func buildTranscribingService(a *arguments) (app.TranscribingService, error) {
	mediaService := ffmpeg.NewMediaService()
	subtitlesService, err := whisper.NewSubtitlesService(a.model)
	if err != nil {
		return nil, err
	}
	return app.NewTranscribingService(mediaService, subtitlesService), nil
}
