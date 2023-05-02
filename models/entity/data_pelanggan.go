package entity

type DataPelanggan struct {
	ID            uint       `json:"id" gorm:"primary_key"`
	NamaPelanggan string     `json:"nama_pelanggan"`
	Contact       string     `json:"contact"`
	Alamat        string     `json:"alamat"`
	DataHarga     DataHarga  `gorm:"foreignKey:IDHarga"`
	DataUsaha     DataUsaha  `gorm:"foreignKey:IDUsaha"`
	DataOutlet    DataOutlet `gorm:"foreignKey:IDOutlet"`
	IDHarga       int        `json:"id_harga"`
	IDUsaha       int        `json:"id_usaha"`
	IDOutlet      int        `json:"id_outlet"`
}
