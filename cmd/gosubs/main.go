package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nightnoryu/gosubs/pkg/gosubs/app"
	"github.com/nightnoryu/gosubs/pkg/gosubs/infrastructure/ffmpeg"
	"github.com/nightnoryu/gosubs/pkg/gosubs/infrastructure/modelprovider"
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
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}

	modelsDir := filepath.Dir(ex)
	modelProvider := modelprovider.NewModelProvider(modelsDir)

	modelPath, err := modelProvider.DownloadModel(a.model)
	if err != nil {
		return nil, err
	}

	mediaService := ffmpeg.NewMediaService()
	subtitlesService, err := whisper.NewSubtitlesService(modelPath)
	if err != nil {
		return nil, err
	}
	return app.NewTranscribingService(mediaService, subtitlesService), nil
}
