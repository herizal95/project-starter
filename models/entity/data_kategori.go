package entity

type DataKategori struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	NamaKategori string     `json:"nama_kategori"`
	DataUsaha    DataUsaha  `gorm:"foreignKey:IDUsaha"`
	DataOutlet   DataOutlet `gorm:"foreignKey:IDOutlet"`
	IDUsaha      int        `json:"id_usaha"`
	IDOutlet     int        `json:"id_outlet"`
}
