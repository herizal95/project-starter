package entity

import (
	"time"

	"google.golang.org/genproto/googleapis/type/decimal"
)

type Transaksi struct {
	ID              uint            `json:"id" gorm:"primary_key"`
	KodeTransaksi   string          `json:"kode_transaksi"`
	Tanggal         time.Time       `json:"tanggal"`
	StatusTransaksi string          `json:"status_transaksi"`
	TotalHarga      decimal.Decimal `gorm:"type:decimal(16,2)"`
	Terhutang       decimal.Decimal `gorm:"type:decimal(16,2)"`
	IDPelanggan     int             `json:"id_pelanggan"`
	IDOutlet        int             `json:"id_outlet"`
	IDUsaha         int             `json:"id_usaha"`
	IDHarga         int             `json:"id_harga"`
	DataUsaha       DataUsaha       `gorm:"foreignKey:IDUsaha"`
	DataOutlet      DataOutlet      `gorm:"foreignKey:IDOutlet"`
	DataPelanggan   DataPelanggan   `gorm:"foreignKey:IDPelanggan"`
	DataHarga       DataHarga       `gorm:"foreignKey:IDHarga"`
}
