package wilayaindonesia

type Kecamatan struct {
	ID            uint   `gorm:"primary_key" json:"id"`
	NamaKecamatan string `json:"nama_kecamatan"`
	IDKabupaten   int    `json:"id_kabupaten"`
}
