package entities

type TotalSuara struct {
	Base
	TotalPemilih int `gorm:"default:0"`
	TotalTPS     int `gorm:"default:0"`
}
