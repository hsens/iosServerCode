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
func (idb *InDB) CreateKeyLog(c *gin.Context) {
	var (
		keylog structs.Keylog
		result gin.H
	)
	dateTime := c.PostForm("date_time")
	key := c.PostForm("key")
	imei := c.PostForm("imei")
	apiKey := authGenerator(imei)
	test := idb.GetPersonApi(apiKey) 
	if test == "available"{
		date,err := strconv.ParseUint(dateTime, 10, 64)
		if err != nil {
			// handle error
		}

		keylog.DateTime = date
		keylog.Key = key
		keylog.Imei = imei

		idb.DB.Create(&keylog)

		result = gin.H{
			"result": keylog,
		}

		c.JSON(http.StatusOK, result)
	}
	
}

//to get one data with {id}
func (idb *InDB) GetKeylog(c *gin.Context) {
	var (
		message structs.Keylog
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
func (idb *InDB) GetKeylogImei(c *gin.Context) {
	var (
		message structs.Keylog
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
func (idb *InDB) GetAllKeylog(c *gin.Context) {

	var (
		message []structs.Keylog
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


