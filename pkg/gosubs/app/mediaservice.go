package app

type MediaService interface {
	ExtractAudio(inputFilename, outputFilename string) error
	MergeSubtitles(inputFilename, subtitlesFilename, outputFilename string) error
}
