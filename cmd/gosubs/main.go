package main

import (
	"log"

	"github.com/nightnoryu/gosubs/pkg/gosubs/app"
	"github.com/nightnoryu/gosubs/pkg/gosubs/infrastructure/ffmpeg"
	"github.com/nightnoryu/gosubs/pkg/gosubs/infrastructure/whisper"
)

func main() {
	args := parseArguments()

	transcribingService, err := buildTranscribingService(args)
	if err != nil {
		log.Fatal(err)
	}

	err = transcribingService.TranscribeVideo(args.inputFilename, args.outputFilename)
	if err != nil {
		log.Fatal(err)
	}
}

func buildTranscribingService(a *arguments) (app.TranscribingService, error) {
	mediaService := ffmpeg.NewMediaService()
	subtitlesService, err := whisper.NewSubtitlesService(a.model)
	if err != nil {
		return nil, err
	}
	return app.NewTranscribingService(mediaService, subtitlesService), nil
}
