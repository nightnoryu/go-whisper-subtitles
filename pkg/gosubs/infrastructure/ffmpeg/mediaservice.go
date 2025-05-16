package ffmpeg

import (
	"fmt"

	"github.com/nightnoryu/gosubs/pkg/gosubs/app"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func NewMediaService() app.MediaService {
	return &mediaService{}
}

type mediaService struct{}

func (s *mediaService) ExtractAudio(inputFilename, outputFilename string) error {
	err := ffmpeg.
		Input(inputFilename).
		Output(outputFilename, ffmpeg.KwArgs{"ar": "16000", "ac": "1", "c:a": "pcm_s16le"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return fmt.Errorf("failed to extract audio: %w", err)
	}

	return nil
}

func (s *mediaService) MergeSubtitles(inputFilename, subtitlesFilename, outputFilename string) error {
	videoWithSubtitles := ffmpeg.Input(inputFilename).Get("v").Filter("subtitles", ffmpeg.Args{subtitlesFilename})
	originalSound := ffmpeg.Input(inputFilename).Get("a")

	input := []*ffmpeg.Stream{videoWithSubtitles, originalSound}
	err := ffmpeg.Output(input, outputFilename, ffmpeg.KwArgs{"c:a": "copy"}).OverWriteOutput().Run()
	if err != nil {
		return fmt.Errorf("failed to merge subtitles: %w", err)
	}

	return nil
}
