package entity

import "google.golang.org/genproto/googleapis/type/decimal"

type DataSaldo struct {
	ID         uint            `gorm:"primary_key"`
	NamaSaldo  string          `json:"nama_saldo"`
	Keluar     decimal.Decimal `gorm:"type:decimal(16,2)"`
	Masuk      decimal.Decimal `gorm:"type:decimal(16,2)"`
	DataOutlet DataOutlet      `gorm:"foreignKey:IDOutlet"`
	IDOutlet   int             `json:"id_outlet"`
}
