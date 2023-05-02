package entity

type DataBarang struct {
	ID           uint         `gorm:"size:36;not null;uniqueIndex;primary_key"`
	NamaBarang   string       `gorm:"size:100;not null" json:"nama_barang"`
	Barcode      string       `json:"barcode"`
	LinkBarcode  string       `json:"link_barcode"`
	HargaBeli    float32      `json:"harga_beli"`
	HargaJual    float32      `json:"harga_jual"`
	Hpp          float32      `json:"hpp"`
	IDKategori   int          `json:"id_kategori"`
	IDSatuan     int          `json:"id_satuan"`
	IDSupplier   int          `json:"id_supplier"`
	IDUsaha      int          `json:"id_usaha"`
	IDOutlet     int          `json:"id_outlet"`
	Status       int          `json:"status"`
	DataKategori DataKategori `gorm:"foreignKey:IDKategori"`
	DataSatuan   DataSatuan   `gorm:"foreignKey:IDSatuan"`
	DataSupplier DataSupplier `gorm:"foreignKey:IDSupplier"`
	DataOutlet   DataOutlet   `gorm:"foreignKey:IDOutlet"`
	DataUsaha    DataUsaha    `gorm:"foreignKey:IDUsaha"`
}
