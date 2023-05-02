package wilayaindonesia

type Kabupaten struct {
	ID            uint   `gorm:"primary_key" json:"id"`
	NamaKabupaten string `json:"nama_kabupaten"`
	IDProvinsi    int    `json:"id_provinsi"`
}
