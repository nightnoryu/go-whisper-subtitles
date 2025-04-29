package app

import "io"

type SubtitlesService interface {
	GenerateSubtitles(inputFile io.ReadSeeker, outputFile io.Writer) error
}
