package entity

type DataSupplier struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	NamaSupplier string     `json:"nama_supplier"`
	Contact      string     `json:"contact"`
	Alamat       string     `json:"alamat"`
	DataUsaha    DataUsaha  `gorm:"foreignKey:IDUsaha"`
	DataOutlet   DataOutlet `gorm:"foreignKey:IDOutlet"`
	IDUsaha      int        `json:"id_usaha"`
	IDOutlet     int        `json:"id_outlet"`
}
