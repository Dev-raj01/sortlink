package models

type URL struct {
	ID      uint64 `json:"id" gorm:"primaryKey"`
	Target  string `json:"target" gorm:"not null"`
	URL     string `json:"url" gorm:"not null"`
	Clicked uint64 `json:"clicked"`
	Random  bool   `json:"random"`
}
