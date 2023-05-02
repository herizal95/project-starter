package migration

import (
	"github.com/herizal95/project-starter/models/entity"
	"github.com/herizal95/project-starter/models/wilayaindonesia"
)

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: entity.User{}},
		{Model: entity.DataUsaha{}},
		{Model: entity.DataOutlet{}},
		{Model: entity.DataHarga{}},
		{Model: entity.DataKategori{}},
		{Model: entity.DataPelanggan{}},
		{Model: entity.DataSatuan{}},
		{Model: entity.DataSaldo{}},
		{Model: entity.DataStok{}},
		{Model: entity.DataSupplier{}},
		{Model: entity.Saldo{}},
		{Model: entity.DataBarang{}},
		{Model: entity.Transaksi{}},
		{Model: wilayaindonesia.Desa{}},
		{Model: wilayaindonesia.Kecamatan{}},
		{Model: wilayaindonesia.Kabupaten{}},
		{Model: wilayaindonesia.Provinsi{}},
	}
}
