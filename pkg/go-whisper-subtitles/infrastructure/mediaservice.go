package infrastructure

import (
	"go-whisper-subtitles/pkg/go-whisper-subtitles/app"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func NewMediaService() app.MediaService {
	return &mediaService{}
}

type mediaService struct{}

func (s *mediaService) ExtractAudio(inputFilename string) (string, error) {
	err := ffmpeg.
		Input(inputFilename).
		Output("test.wav", ffmpeg.KwArgs{"ar": "16000", "ac": "1", "c:a": "pcm_s16le"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return "", err
	}

	return "test.wav", nil
}

func (s *mediaService) MergeSubtitles(inputFilename, subtitlesFilename string) error {
	input := []*ffmpeg.Stream{ffmpeg.Input(inputFilename), ffmpeg.Input(subtitlesFilename)}

	err := ffmpeg.Output(input, "result.mkv", ffmpeg.KwArgs{"c": "copy", "c:s": "srt"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return err
	}

	return nil
}
