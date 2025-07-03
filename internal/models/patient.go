package models

type Patient struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Age   int    `json:"age"`
	Notes string `json:"notes"`
}
