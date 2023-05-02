package wilayaindonesia

type Provinsi struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	NamaProvinsi string `json:"nama_provinsi"`
}
