package data_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
)

func FindOutlets(c *gin.Context) {

	data := new([]entity.DataOutlet)

	if err := database.DB.Table("data_outlets").Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully fetch all data outlet.", data)
}

func FindIDOutlet(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataOutlet)

	if err := database.DB.Table("data_outlets").Where("id = ?", id).First(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, ID data not found.", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully Get data outlet.", data)
}

func SaveOutlet(c *gin.Context) {

	input := new(entity.DataOutlet)

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data := entity.DataOutlet{
		NamaOutlet:  input.NamaOutlet,
		Alamat:      input.Alamat,
		Contact:     input.Contact,
		Pic:         input.Pic,
		IDUsaha:     input.IDUsaha,
		IDProvinsi:  input.IDProvinsi,
		IDKabupaten: input.IDKabupaten,
		IDKecamatan: input.IDKecamatan,
		IDDesa:      input.IDDesa,
		IsPusat:     input.IsPusat,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "failed to create data", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to created data outlet.", data)
}

func UpdateOutlet(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataOutlet)
	dataInput := new(entity.DataOutlet)

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := database.DB.Table("data_outlets").Where("id = ?", id).Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data.NamaOutlet = dataInput.NamaOutlet
	data.Alamat = dataInput.Alamat
	data.Contact = dataInput.Contact
	data.IDProvinsi = dataInput.IDProvinsi
	data.IDKabupaten = dataInput.IDKabupaten
	data.IDKecamatan = dataInput.IDKecamatan
	data.IDDesa = dataInput.IDDesa

	if erruserUpdate := database.DB.Table("data_outlets").Where("id = ?", id).Updates(&data).Error; erruserUpdate != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, can't update data", nil)
		return
	}
	helper.Response(c.Writer, http.StatusOK, "Successfully to update data outlet!", data)

}

func DeleteOutlet(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataOutlet)

	errFind := database.DB.Table("data_outlets").Where("id = ?", id).Find(&data).Error
	if errFind != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, errFind.Error(), nil)
		return
	}

	if err := database.DB.Table("data_outlets").Unscoped().Where("id = ?", id).Delete(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, "can't delete data!", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to delete data!", nil)
}
