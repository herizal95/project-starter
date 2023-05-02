package data_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
)

func FindSaldo(c *gin.Context) {
	data := new([]entity.DataSaldo)

	if err := database.DB.Table("data_saldos").Preload("DataOutlet").Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully fetch all data saldo.", data)
}

func FindIDSaldo(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataSaldo)

	if err := database.DB.Table("data_saldos").Preload("DataOutlet").Where("id = ?", id).First(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, ID data not found.", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully Get data saldo.", data)
}

func SaveSaldo(c *gin.Context) {

	input := new(entity.DataSaldo)

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data := entity.DataSaldo{
		NamaSaldo: input.NamaSaldo,
		Keluar:    input.Keluar,
		Masuk:     input.Masuk,
		IDOutlet:  input.IDOutlet,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "failed to create data", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to created data saldo.", nil)
}

func UpdateSaldo(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataSaldo)
	dataInput := new(entity.DataSaldo)

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := database.DB.Table("data_saldos").Where("id = ?", id).Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data.NamaSaldo = dataInput.NamaSaldo
	data.Keluar = dataInput.Keluar
	data.Masuk = dataInput.Masuk
	data.IDOutlet = dataInput.IDOutlet

	if erruserUpdate := database.DB.Table("data_saldos").Where("id = ?", id).Updates(&data).Error; erruserUpdate != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, can't update data", nil)
		return
	}
	helper.Response(c.Writer, http.StatusOK, "Successfully to update data satuan!", nil)

}

func DeleteSaldo(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataSaldo)

	errFind := database.DB.Table("data_saldos").Where("id = ?", id).Find(&data).Error
	if errFind != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, errFind.Error(), nil)
		return
	}

	if err := database.DB.Table("data_saldos").Unscoped().Where("id = ?", id).Delete(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, "can't delete data!", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to delete data!", nil)
}
