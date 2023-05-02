package data_controller

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
	"github.com/skip2/go-qrcode"
)

func FindBarang(ctx *gin.Context) {
	data := new([]entity.DataBarang)
	if err := database.DB.Table("data_barangs").
		Preload("DataKategori").Preload("DataOutlet").Preload("DataUsaha").Preload("DataSupplier").Preload("DataSatuan").
		Find(&data).Error; err != nil {
		helper.Response(ctx.Writer, http.StatusBadRequest, "Bad Request, Error load data barang", nil)
		return
	}
	helper.Response(ctx.Writer, http.StatusOK, "Successfully fetch all data barang", data)
}
func FindIDBarang(ctx *gin.Context) {
	id := ctx.Param("id")
	data := new(entity.DataBarang)
	if err := database.DB.Table("data_barangs").Where("id = ?", id).First(&data).Error; err != nil {
		helper.Response(ctx.Writer, http.StatusNotFound, "Data barang is not found.", nil)
		return
	}
	helper.Response(ctx.Writer, http.StatusOK, "Successfully fetch data barang", data)
}
func SaveBarang(ctx *gin.Context) {
	var input entity.DataBarang
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helper.Response(ctx.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Generate the QR code
	png, err := qrcode.Encode(input.Barcode, qrcode.Medium, 256)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save file to device
	err = ioutil.WriteFile("public/files/barcode/"+input.NamaBarang, png, 0777)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Failed to save file to device",
		})
		return
	}
	// Save the URL to the QR code model
	url := "/public/files/" + filepath.ToSlash(filepath.Join("barcode", input.NamaBarang))

	data := entity.DataBarang{
		NamaBarang: input.NamaBarang,
		IDSatuan:   input.IDSatuan,
		HargaBeli:  input.HargaBeli,
		HargaJual:  input.HargaJual,
		Hpp:        input.Hpp,
		IDSupplier: input.IDSupplier,
		Status:     1,
		Barcode:    url,
		IDUsaha:    input.IDUsaha,
		IDOutlet:   input.IDOutlet,
		IDKategori: input.IDKategori,
	}

	if err := database.DB.Table("data_barangs").Create(&data).Error; err != nil {
		helper.Response(ctx.Writer, http.StatusInternalServerError, "Error create data barang", nil)
		return
	}

	helper.Response(ctx.Writer, http.StatusOK, "Successfully to created data barang", nil)
}
func UpdateBarang(ctx *gin.Context) {}
func DeleteBarang(ctx *gin.Context) {}
