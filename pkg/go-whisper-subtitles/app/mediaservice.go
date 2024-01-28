package app

type MediaService interface {
	ExtractAudio() error
	MergeSubtitles() error
}
