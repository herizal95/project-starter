package wilayaindonesia

type Desa struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	NamaDesa    string `json:"nama_desa"`
	IDKecamatan int    `json:"id_kecamatan"`
}
