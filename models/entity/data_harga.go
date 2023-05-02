package entity

type DataHarga struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	NamaHarga  string     `json:"nama_harga"`
	DataUsaha  DataUsaha  `gorm:"foreignKey:IDUsaha"`
	DataOutlet DataOutlet `gorm:"foreignKey:IDOutlet"`
	IDUsaha    int        `json:"id_usaha"`
	IDOutlet   int        `json:"id_outlet"`
}
