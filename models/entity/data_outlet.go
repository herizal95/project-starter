package entity

type DataOutlet struct {
	ID          uint   `gorm:"primary_key"`
	NamaOutlet  string `json:"nama_outlet"`
	Alamat      string `json:"alamat"`
	Contact     string `json:"contact"`
	Pic         string `json:"pic"`
	IDUsaha     int    `json:"id_usaha"`
	IDProvinsi  int    `json:"id_provinsi"`
	IDKabupaten int    `json:"id_kabupaten"`
	IDKecamatan int    `json:"id_kecamatan"`
	IDDesa      int    `json:"id_desa"`
	IsPusat     int    `json:"is_pusat"`
}
