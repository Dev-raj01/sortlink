package models

type URL struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	TargetURL string `json:"target_url" gorm:"not null"`
	ShortURL  string `json:"short_url" gorm:"unique;not null"`
	Clicked   uint64
	Random    bool   `json:"random"`
}
