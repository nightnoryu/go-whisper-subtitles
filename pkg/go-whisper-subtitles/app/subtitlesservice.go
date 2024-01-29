package app

type SubtitlesService interface {
	GenerateSubtitles(inputFilename string) (string, error)
}
