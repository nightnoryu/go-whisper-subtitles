package infrastructure

import (
	"fmt"
	"io"
	"os"
	"time"

	"go-whisper-subtitles/pkg/go-whisper-subtitles/app"

	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/go-audio/wav"
)

func NewSubtitlesService(modelPath string) app.SubtitlesService {
	return &subtitlesService{
		modelPath: modelPath,
	}
}

type subtitlesService struct {
	modelPath string
}

func (s *subtitlesService) GenerateSubtitles(inputFilename string) (string, error) {
	model, err := whisper.New(s.modelPath)
	if err != nil {
		return "", err
	}
	defer model.Close()

	context, err := model.NewContext()
	if err != nil {
		return "", err
	}

	samples, err := s.loadSamples(inputFilename)
	if err != nil {
		return "", err
	}

	context.ResetTimings()
	err = context.Process(samples, nil, nil)
	if err != nil {
		return "", err
	}

	outputFile, err := os.Create("test.srt")
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	n := 1
	for {
		segment, err := context.NextSegment()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		fmt.Fprintln(outputFile, n)
		fmt.Fprintln(outputFile, srtTimestamp(segment.Start), " --> ", srtTimestamp(segment.End))
		fmt.Fprintln(outputFile, segment.Text)
		fmt.Fprintln(outputFile)
		n++
	}

	err = os.Remove(inputFilename)
	if err != nil {
		return "", err
	}

	return "test.srt", nil
}

func (s *subtitlesService) loadSamples(inputFilename string) ([]float32, error) {
	file, err := os.Open(inputFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, err
	}

	return buf.AsFloat32Buffer().Data, nil
}

func srtTimestamp(t time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d,%03d", t/time.Hour, (t%time.Hour)/time.Minute, (t%time.Minute)/time.Second, (t%time.Second)/time.Millisecond)
}
