package events

type Fetcher interface {
	Fetch(limit int) ([]Events, error)
}

type Processor interface {
	Process(e Events) error
}

type Type int

const (
	Unknown Type = iota
	Messaage
)

type Events struct {
	Type Type
	Text string
}
