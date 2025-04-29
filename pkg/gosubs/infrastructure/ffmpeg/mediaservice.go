package ffmpeg

import (
	"github.com/nightnoryu/gosubs/pkg/gosubs/app"
	"github.com/pkg/errors"

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
	return errors.WithStack(err)
}

func (s *mediaService) MergeSubtitles(inputFilename, subtitlesFilename, outputFilename string) error {
	videoWithSubtitles := ffmpeg.Input(inputFilename).Get("v").Filter("subtitles", ffmpeg.Args{subtitlesFilename})
	originalSound := ffmpeg.Input(inputFilename).Get("a")

	input := []*ffmpeg.Stream{videoWithSubtitles, originalSound}
	err := ffmpeg.Output(input, outputFilename, ffmpeg.KwArgs{"c:a": "copy"}).OverWriteOutput().Run()

	return errors.WithStack(err)
}
