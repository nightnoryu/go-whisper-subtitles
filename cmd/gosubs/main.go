package main

import (
	"context"
	"log"
	"os"

	"github.com/nightnoryu/gosubs/pkg/gosubs/app"
	"github.com/nightnoryu/gosubs/pkg/gosubs/infrastructure/ffmpeg"
	"github.com/nightnoryu/gosubs/pkg/gosubs/infrastructure/whisper"

	"github.com/urfave/cli/v3"
)

const (
	argInputFile  = "input-file"
	argOutputFile = "output-file"
	argModelPath  = "model"
)

func main() {
	cmd := &cli.Command{
		Name:  "gosubs",
		Usage: "generate subtitles for a video",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      argInputFile,
				UsageText: "path to the input video",
			},
			&cli.StringArg{
				Name:      argOutputFile,
				UsageText: "path to the output video",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      argModelPath,
				Usage:     "path to the ggml whisper model",
				Required:  true,
				Value:     "model/ggml-base.bin",
				TakesFile: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			inputFilename := cmd.StringArg(argInputFile)
			outputFilename := cmd.StringArg(argOutputFile)
			modelPath := cmd.String(argModelPath)

			mediaService := ffmpeg.NewMediaService()
			subtitlesService, err := whisper.NewSubtitlesService(modelPath)
			if err != nil {
				return err
			}
			transcribingService := app.NewTranscribingService(mediaService, subtitlesService)

			return transcribingService.TranscribeFile(inputFilename, outputFilename)
		}}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
