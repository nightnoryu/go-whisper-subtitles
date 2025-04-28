package main

import (
	"fmt"
	"log"
	"os"

	"gosubs/pkg/gosubs/infrastructure"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "test",
		Usage: "test",
		Action: func(ctx *cli.Context) error {
			inputFilename := ctx.Args().Get(0)
			fmt.Println("Transcoding file ", inputFilename)

			mediaService := infrastructure.NewMediaService()
			audioFile, err := mediaService.ExtractAudio(inputFilename)
			if err != nil {
				return err
			}

			subtitlesService := infrastructure.NewSubtitlesService("ggml-base.en.bin")
			subtitlesFile, err := subtitlesService.GenerateSubtitles(audioFile)
			if err != nil {
				return err
			}

			return mediaService.MergeSubtitles(inputFilename, subtitlesFile)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
