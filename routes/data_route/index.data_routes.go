package data_route

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/controllers/data_controller"
)

func DataRoutes(app *gin.Engine) {

	router := app

	routes := router.Group(os.Getenv("BaseUrl"))

	// DATA USAHA
	routes.GET("/data-usahas", data_controller.FindUsaha)
	routes.POST("/data-usahas", data_controller.SaveUsaha)
	routes.GET("/data-usahas/:id", data_controller.FindIDUsaha)
	routes.PATCH("/data-usahas/:id", data_controller.UpdateUsaha)
	routes.DELETE("/data-usahas/:id", data_controller.DeleteUsaha)

	// DATA OUTLET
	routes.GET("/data-outlets", data_controller.FindOutlets)
	routes.POST("/data-outlets", data_controller.SaveOutlet)
	routes.GET("/data-outlets/:id", data_controller.FindIDOutlet)
	routes.PATCH("/data-outlets/:id", data_controller.UpdateOutlet)
	routes.DELETE("/data-outlets/:id", data_controller.DeleteOutlet)

	// DATA HARGA
	routes.GET("/data-hargas", data_controller.FindHarga)
	routes.POST("/data-hargas", data_controller.SaveHarga)
	routes.GET("/data-hargas/:id", data_controller.FindIDHarga)
	routes.PATCH("/data-hargas/:id", data_controller.UpdateHarga)
	routes.DELETE("/data-hargas/:id", data_controller.DeleteHarga)

	// DATA KATEGORI
	routes.GET("/data-kategoris", data_controller.FindKategori)
	routes.POST("/data-kategoris", data_controller.SaveKategori)
	routes.GET("/data-kategoris/:id", data_controller.FindIDKategori)
	routes.PATCH("/data-kategoris/:id", data_controller.UpdateKategori)
	routes.DELETE("/data-kategoris/:id", data_controller.DeleteKategori)

	// DATA PELANGGAN
	routes.GET("/data-pelanggans", data_controller.FindPelanggan)
	routes.POST("/data-pelanggans", data_controller.SavePelanggan)
	routes.GET("/data-pelanggans/:id", data_controller.FindIDPelanggan)
	routes.PATCH("/data-pelanggans/:id", data_controller.UpdatePelanggan)
	routes.DELETE("/data-pelanggans/:id", data_controller.DeletePelanggan)

	// DATA SATUAN
	routes.GET("/data-satuans", data_controller.FindSatuan)
	routes.POST("/data-satuans", data_controller.SaveSatuan)
	routes.GET("/data-satuans/:id", data_controller.FindIDSatuan)
	routes.PATCH("/data-satuans/:id", data_controller.UpdateSatuan)
	routes.DELETE("/data-satuans/:id", data_controller.DeleteSatuan)

	// DATA Saldo
	routes.GET("/data-saldos", data_controller.FindSaldo)
	routes.POST("/data-saldos", data_controller.SaveSaldo)
	routes.GET("/data-saldos/:id", data_controller.FindIDSaldo)
	routes.PATCH("/data-saldos/:id", data_controller.UpdateSaldo)
	routes.DELETE("/data-saldos/:id", data_controller.DeleteSaldo)

	// DATA SUPPLIERS
	routes.GET("/data-suppliers", data_controller.FindSupplier)
	routes.POST("/data-suppliers", data_controller.SaveSupplier)
	routes.GET("/data-suppliers/:id", data_controller.FindIDSupplier)
	routes.PATCH("/data-suppliers/:id", data_controller.UpdateSupplier)
	routes.DELETE("/data-suppliers/:id", data_controller.DeleteSupplier)

	// SALDO
	routes.GET("/saldos", data_controller.FindSaldos)
	routes.POST("/saldos", data_controller.SaveSaldos)
	routes.GET("/saldos/:id", data_controller.FindIDSaldos)
	routes.PATCH("/saldos/:id", data_controller.UpdateSaldos)
	routes.DELETE("/saldos/:id", data_controller.DeleteSaldos)

	// DATA BARANG
	routes.GET("/data-barangs", data_controller.FindBarang)
	routes.POST("/data-barangs", data_controller.SaveBarang)
	routes.GET("/data-barangs/:id", data_controller.FindIDBarang)
	routes.PATCH("/data-barangs/:id", data_controller.UpdateBarang)
	routes.DELETE("/data-barangs/:id", data_controller.DeleteBarang)
}
