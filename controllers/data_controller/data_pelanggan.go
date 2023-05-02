package data_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
)

func FindPelanggan(c *gin.Context) {
	data := new([]entity.DataPelanggan)

	if err := database.DB.Table("data_pelanggans").Preload("DataHarga").Preload("DataUsaha").Preload("DataOutlet").Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully fetch all data.", data)
}

func FindIDPelanggan(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataPelanggan)

	if err := database.DB.Table("data_pelanggans").Where("id = ?", id).First(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, ID data not found.", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully Get data harga.", data)
}

func SavePelanggan(c *gin.Context) {

	input := new(entity.DataPelanggan)

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data := entity.DataPelanggan{
		NamaPelanggan: input.NamaPelanggan,
		Alamat:        input.Alamat,
		Contact:       input.Contact,
		IDHarga:       input.IDHarga,
		IDOutlet:      input.IDOutlet,
		IDUsaha:       input.IDUsaha,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "failed to create data", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to created data.", nil)
}

func UpdatePelanggan(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataPelanggan)
	dataInput := new(entity.DataPelanggan)

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := database.DB.Table("data_pelanggans").Where("id = ?", id).Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data.NamaPelanggan = dataInput.NamaPelanggan
	data.Alamat = dataInput.Alamat
	data.Contact = dataInput.Contact
	data.IDHarga = dataInput.IDHarga
	data.IDOutlet = dataInput.IDOutlet
	data.IDUsaha = dataInput.IDUsaha

	if erruserUpdate := database.DB.Table("data_pelanggans").Where("id = ?", id).Updates(&data).Error; erruserUpdate != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, can't update data", nil)
		return
	}
	helper.Response(c.Writer, http.StatusOK, "Successfully to update data!", nil)

}

func DeletePelanggan(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataPelanggan)

	errFind := database.DB.Table("data_pelanggans").Where("id = ?", id).Find(&data).Error
	if errFind != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, errFind.Error(), nil)
		return
	}

	if err := database.DB.Table("data_pelanggans").Unscoped().Where("id = ?", id).Delete(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, "can't delete data!", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to delete data!", nil)
}
