package app

import (
	"os"

	"github.com/pkg/errors"
)

type TranscribingService interface {
	TranscribeFile(inputFilename, outputFilename string) error
}

func NewTranscribingService(
	mediaService MediaService,
	subtitlesService SubtitlesService,
) TranscribingService {
	return &transcribingService{
		mediaService:     mediaService,
		subtitlesService: subtitlesService,
	}
}

type transcribingService struct {
	mediaService     MediaService
	subtitlesService SubtitlesService
}

func (s *transcribingService) TranscribeFile(inputFilename, outputFilename string) error {
	tempAudioFile, err := os.CreateTemp("", "*.wav")
	if err != nil {
		return errors.WithStack(err)
	}
	defer tempAudioFile.Close()
	defer os.Remove(tempAudioFile.Name())

	tempSubtitlesFile, err := os.CreateTemp("", "*.srt")
	if err != nil {
		return errors.WithStack(err)
	}
	defer tempSubtitlesFile.Close()
	defer os.Remove(tempSubtitlesFile.Name())

	err = s.mediaService.ExtractAudio(inputFilename, tempAudioFile.Name())
	if err != nil {
		return err
	}

	err = s.subtitlesService.GenerateSubtitles(tempAudioFile, tempSubtitlesFile)
	if err != nil {
		return err
	}
	tempSubtitlesFile.Sync()

	return s.mediaService.MergeSubtitles(inputFilename, tempSubtitlesFile.Name(), outputFilename)
}
