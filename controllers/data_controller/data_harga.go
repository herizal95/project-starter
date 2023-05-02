package data_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
)

func FindHarga(c *gin.Context) {
	data := new([]entity.DataHarga)

	if err := database.DB.Table("data_hargas").Preload("DataUsaha").Preload("DataOutlet").Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully fetch all data harga.", data)
}

func FindIDHarga(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataHarga)

	if err := database.DB.Table("data_hargas").Preload("DataUsaha").Preload("DataOutlet").Where("id = ?", id).First(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, ID data not found.", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully Get data harga.", data)
}

func SaveHarga(c *gin.Context) {

	input := new(entity.DataHarga)

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data := entity.DataHarga{
		NamaHarga: input.NamaHarga,
		IDOutlet:  input.IDOutlet,
		IDUsaha:   input.IDUsaha,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "failed to create data", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to created data harga.", data)
}

func UpdateHarga(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataHarga)
	dataInput := new(entity.DataHarga)

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := database.DB.Table("data_hargas").Where("id = ?", id).Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data.NamaHarga = dataInput.NamaHarga
	data.IDOutlet = dataInput.IDOutlet
	data.IDUsaha = dataInput.IDUsaha

	if erruserUpdate := database.DB.Table("data_hargas").Where("id = ?", id).Updates(&data).Error; erruserUpdate != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, can't update data", nil)
		return
	}
	helper.Response(c.Writer, http.StatusOK, "Successfully to update data harga!", nil)

}

func DeleteHarga(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataHarga)

	errFind := database.DB.Table("data_hargas").Where("id = ?", id).Find(&data).Error
	if errFind != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, errFind.Error(), nil)
		return
	}

	if err := database.DB.Table("data_hargas").Unscoped().Where("id = ?", id).Delete(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, "can't delete data!", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to delete data!", nil)
}
