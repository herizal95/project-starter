package data_controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
)

func FindUsaha(c *gin.Context) {

	data := new([]entity.DataUsaha)

	if err := database.DB.Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully fetch all data usaha.", data)
}

func FindIDUsaha(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataUsaha)

	if err := database.DB.Table("data_usahas").Where("id = ?", id).First(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, ID data not found.", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully Get data usaha.", data)
}

func SaveUsaha(c *gin.Context) {

	input := new(entity.DataUsaha)

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data := entity.DataUsaha{
		NamaUsaha:    input.NamaUsaha,
		Alamat:       input.Alamat,
		Contact:      input.Contact,
		Endtime:      time.Now(),
		Subscription: "full",
		IsActive:     1,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "failed to create data", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to created data usaha.", data)
}

func UpdateUsaha(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataUsaha)
	dataInput := new(entity.DataUsaha)

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := database.DB.Table("data_usahas").Where("id = ?", id).Find(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data.NamaUsaha = dataInput.NamaUsaha
	data.Alamat = dataInput.Alamat
	data.Contact = dataInput.Contact
	data.Endtime = time.Now()

	if erruserUpdate := database.DB.Table("data_usahas").Where("id = ?", id).Updates(&data).Error; erruserUpdate != nil {
		helper.Response(c.Writer, http.StatusBadRequest, "Bad Request, can't update data", nil)
		return
	}
	helper.Response(c.Writer, http.StatusOK, "Successfully to update user!", data)

}

func DeleteUsaha(c *gin.Context) {

	id := c.Param("id")
	data := new(entity.DataUsaha)

	errFind := database.DB.Table("data_usahas").Where("id = ?", id).Find(&data).Error
	if errFind != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, errFind.Error(), nil)
		return
	}

	if err := database.DB.Table("data_usahas").Unscoped().Where("id = ?", id).Delete(&data).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, "can't delete data!", nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully to delete data!", nil)
}
