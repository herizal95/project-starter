package data_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
)

func FindSaldos(c *gin.Context) {
	data := new([]entity.Saldo)

	if err := database.DB.Table("data_Saldoss").Preload("DataOutlet").Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully fetch all data Saldos.", data)
}

func FindIDSaldos(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.Saldo)

	if err := database.DB.Table("data_Saldoss").Preload("DataOutlet").Where("id = ?", id).First(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, ID data not found.", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully Get data Saldos.", data)
}

func SaveSaldos(c *gin.Context) {

	input := new(entity.Saldo)

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data := entity.Saldo{
		NamaKas:  input.NamaKas,
		Jumlah:   input.Jumlah,
		IDOutlet: input.IDOutlet,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "failed to create data", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to created data Saldos.", nil)
}

func UpdateSaldos(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.Saldo)
	dataInput := new(entity.Saldo)

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := database.DB.Table("data_Saldoss").Where("id = ?", id).Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data.NamaKas = dataInput.NamaKas
	data.Jumlah = dataInput.Jumlah
	data.IDOutlet = dataInput.IDOutlet

	if erruserUpdate := database.DB.Table("data_Saldoss").Where("id = ?", id).Updates(&data).Error; erruserUpdate != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, can't update data", nil)
		return
	}
	helper.Response(c.Writer, http.StatusOK, "Successfully to update data satuan!", nil)

}

func DeleteSaldos(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.Saldo)

	errFind := database.DB.Table("data_Saldoss").Where("id = ?", id).Find(&data).Error
	if errFind != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, errFind.Error(), nil)
		return
	}

	if err := database.DB.Table("data_Saldoss").Unscoped().Where("id = ?", id).Delete(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, "can't delete data!", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to delete data!", nil)
}
