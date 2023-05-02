package entity

type DataStok struct {
	ID         uint       `gorm:"primary_key"`
	IDBarang   int        `json:"id_barang"`
	Stok       int32      `json:"stok"`
	IDOutlet   int        `json:"id_outlet"`
	IDUsaha    int        `json:"id_usaha"`
	DataUsaha  DataUsaha  `gorm:"foreignKey:IDUsaha"`
	DataOutlet DataOutlet `gorm:"foreignKey:IDOutlet"`
	DataBarang DataBarang `gorm:"foreignKey:IDBarang"`
}
