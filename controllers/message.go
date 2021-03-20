package controllers

import (
	//"fmt"
	"net/http"

	//"github.com/dgrijalva/jwt-go"
	"strconv"
	"../structs"
	"github.com/gin-gonic/gin"
	
)

//to get one data with {id}
func (idb *InDB) GetMessageImei(c *gin.Context) {
	var (
		message structs.Message
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

//to get one data with {id}
func (idb *InDB) GetMessage(c *gin.Context) {
	var (
		message structs.Message
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


//get all data in person
func (idb *InDB) GetAllMessages(c *gin.Context) {

	var (
		message []structs.Message
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

//create new data to the database
func (idb *InDB) CreateMessage(c *gin.Context) {
	var (
		message structs.Message
		result gin.H
	)
	dateTime := c.PostForm("date_time")
	fromMe := c.PostForm("from_me")
	phone := c.PostForm("phone_number")
	textMessage := c.PostForm("text_message")
	
	imei := c.PostForm("imei")
	apiKey := authGenerator(imei)
	test := idb.GetPersonApi(apiKey) 
	if test == "available"{
		date,err := strconv.ParseUint(dateTime, 10, 64)
		if err != nil {
			// handle error
		}

		fMe,err := strconv.ParseBool(fromMe)
		if err != nil {
			// handle error
		}

		message.DateTime = date
		message.FromMe = fMe
		message.PhoneNumber = phone
		message.TextMessage = textMessage
		
		message.Imei = imei

		idb.DB.Create(&message)

		result = gin.H{
			"result": message,
		}

		c.JSON(http.StatusOK, result)
	}
	
}



// delete data with {id}
func (idb *InDB) DeleteMessage(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}

// func authGenerator(imei string) string {
// 	sign := jwt.New(jwt.GetSigningMethod("HS256"))
// 	token, err := sign.SignedString([]byte(imei))
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}

// 	fmt.Print(token)
// 	return token
// }