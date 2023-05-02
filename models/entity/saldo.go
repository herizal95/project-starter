package entity

import "google.golang.org/genproto/googleapis/type/decimal"

type Saldo struct {
	ID         uint            `json:"id" gorm:"primary_key"`
	NamaKas    string          `json:"nama_pelanggan"`
	Jumlah     decimal.Decimal `gorm:"type:decimal(16,2)"`
	DataOutlet DataOutlet      `gorm:"foreignKey:IDOutlet"`
	IDOutlet   int             `json:"id_outlet"`
}
