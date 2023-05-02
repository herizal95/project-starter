package entity

type DataSatuan struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	NamaSatuan string     `json:"nama_satuan"`
	DataUsaha  DataUsaha  `gorm:"foreignKey:IDUsaha"`
	DataOutlet DataOutlet `gorm:"foreignKey:IDOutlet"`
	IDUsaha    int        `json:"id_usaha"`
	IDOutlet   int        `json:"id_outlet"`
}
