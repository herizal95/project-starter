package entity

import "time"

type DataUsaha struct {
	ID           uint      `gorm:"size:36;not null;uniqueIndex;primary_key"`
	NamaUsaha    string    `json:"nama_usaha"`
	Alamat       string    `json:"alamat"`
	Contact      string    `json:"contact"`
	Endtime      time.Time `json:"endtime"`
	Subscription string    `json:"subcription"`
	IsActive     int       `json:"is_active"`
}
