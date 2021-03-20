package controllers

import (
	"fmt"
	"net/http"
	//"github.com/dgrijalva/jwt-go"
	"strconv"
	"../structs"
	"github.com/gin-gonic/gin"
	
)

//create new data to the database
func (idb *InDB) SubmitPicture(c *gin.Context) {
	var (
		picture structs.Picture
		result gin.H
	)
	dateTime := c.PostForm("date_time")
	origin := c.PostForm("origin")
	imei := c.PostForm("imei")

	// Input Tipe File
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	test := idb.GetPersonImei(imei) 
	if test == "available"{
		date,err := strconv.ParseUint(dateTime, 10, 64)
		if err != nil {
			// handle error
		}
		// Set Folder untuk menyimpan filenya
		path := "/var/www/ios.id-successfactors.com/PurpleAdmin/imagesFromBot/" + file.Filename

		picture.DateTime = date
		picture.PictureName = file.Filename 
		picture.PicturePath = path
		picture.Origin = origin
		picture.Imei = imei
		
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}
		idb.DB.Create(&picture)

		result = gin.H{
			"result": picture,
		}

		c.JSON(http.StatusOK, result)
	}else{
		c.JSON(http.StatusBadRequest, "ERROR!")
	}
	
}

//to get one data with {id}
func (idb *InDB) GetPicture(c *gin.Context) {
	var (
		message structs.Picture
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&message).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": message,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)

}

//to get one data with {id}
func (idb *InDB) GetPictureImei(c *gin.Context) {
	var (
		message structs.Picture
		result gin.H
	)
	imei := c.Param("imei")
	err := idb.DB.Where("imei=?", imei).First(&message).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": message,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)

}

//get all data in person
func (idb *InDB) GetAllPicture(c *gin.Context) {

	var (
		message []structs.Picture
		result  gin.H
	)

	idb.DB.Find(&message)
	if len(message) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": message,
			"count":  len(message),
		}
	}

	c.JSON(http.StatusOK, result)
}