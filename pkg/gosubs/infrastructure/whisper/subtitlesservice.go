package whisper

import (
	stderrors "errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/nightnoryu/gosubs/pkg/gosubs/app"

	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/go-audio/wav"
	"github.com/pkg/errors"
)

var ErrModelNotFound = stderrors.New("model not found")

func NewSubtitlesService(modelPath string) (app.SubtitlesService, error) {
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		return nil, errors.WithStack(ErrModelNotFound)
	}

	return &subtitlesService{
		modelPath: modelPath,
	}, nil
}

type subtitlesService struct {
	modelPath string
}

func (s *subtitlesService) GenerateSubtitles(inputFile io.ReadSeeker, outputFile io.Writer) error {
	model, err := whisper.New(s.modelPath)
	if err != nil {
		return err
	}
	defer model.Close()

	context, err := model.NewContext()
	if err != nil {
		return err
	}
	err = context.SetLanguage("auto")
	if err != nil {
		return err
	}

	samples, err := s.loadSamples(inputFile)
	if err != nil {
		return err
	}

	context.ResetTimings()
	err = context.Process(samples, nil, nil, nil)
	if err != nil {
		return err
	}

	n := 1
	for {
		segment, err := context.NextSegment()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fmt.Fprintln(outputFile, n)
		fmt.Fprintln(outputFile, srtTimestamp(segment.Start), " --> ", srtTimestamp(segment.End))
		fmt.Fprintln(outputFile, segment.Text)
		fmt.Fprintln(outputFile)
		n++
	}

	return nil
}

func (s *subtitlesService) loadSamples(inputFile io.ReadSeeker) ([]float32, error) {
	decoder := wav.NewDecoder(inputFile)
	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, err
	}

	return buf.AsFloat32Buffer().Data, nil
}

func srtTimestamp(t time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d,%03d", t/time.Hour, (t%time.Hour)/time.Minute, (t%time.Minute)/time.Second, (t%time.Second)/time.Millisecond)
}
