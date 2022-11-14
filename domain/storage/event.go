package storage

type EventType string

const (
	EventCreate EventType = "create"
)

type FileCategory string

const (
	CategoryVideo FileCategory = "video"
	CategoryAudio FileCategory = "audio"
)

type Event struct {
	Category FileCategory
	Type     EventType
	FileName string
}
