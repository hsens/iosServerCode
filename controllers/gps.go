package controllers

import (
	//"fmt"
	"net/http"

	//"github.com/dgrijalva/jwt-go"
	"strconv"
	"../structs"
	"github.com/gin-gonic/gin"
	
)

//create new data to the database
func (idb *InDB) CreateGPS(c *gin.Context) {
	var (
		gps structs.GPS
		result gin.H
	)
	street := c.PostForm("street")
	city := c.PostForm("city")
	countryCode := c.PostForm("country_code")
	latitude := c.PostForm("latitude")
	longtitude := c.PostForm("longtitude")
	imei := c.PostForm("imei")
	//apiKey := authGenerator(imei)
	test := idb.GetPersonImei(imei) 
	if test == "available"{
		lat,err := strconv.ParseFloat(latitude, 64)
		if err != nil {
			// handle error
		}

		long,err := strconv.ParseFloat(longtitude, 64)
		if err != nil {
			// handle error
		}

		gps.Street = street
		gps.City = city 
		gps.CountryCode = countryCode
		gps.Latitude = lat
		gps.Longtitude = long
		gps.Imei = imei

		idb.DB.Create(&gps)

		result = gin.H{
			"result": gps,
		}

		c.JSON(http.StatusOK, result)
	}
	
}

//to get one data with {id}
func (idb *InDB) GetGPS(c *gin.Context) {
	var (
		message structs.GPS
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
func (idb *InDB) GetGPSImei(c *gin.Context) {
	var (
		message structs.GPS
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
func (idb *InDB) GetAllGPS(c *gin.Context) {

	var (
		message []structs.GPS
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

//update GPS