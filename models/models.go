package models

type URL struct {
	ID      uint64
	Target  string
	URL     string
	Clicked uint64
	Random  bool
}
