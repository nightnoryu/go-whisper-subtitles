package app

type MediaService interface {
	ExtractAudio(inputFilename string) (string, error)
	MergeSubtitles(inputFilename, subtitlesFilename string) error
}
