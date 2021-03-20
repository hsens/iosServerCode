package controllers

import (
	"net/http"
	"../structs"
	"github.com/gin-gonic/gin"
)

// CreateContact is create contact of the imei endpoint handler
// @Summary get contact from the endpoint
// @Description create a contact based on the endpoint
// @Tags contact
// @Accept  json
// @Produce  json
// @Param user body model.CreateUser true "create user"
// @Success 200 {object} model.UserDetail
// @Failure 400 {object} httputil.HTTPError
// @Router api/contact/ [post]
func (idb *InDB) CreateContact(c *gin.Context) {
	var (
		contact structs.Contact

		result gin.H
	)
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	phone := c.PostForm("phone_number")
	imei := c.PostForm("imei")
	
	
	if idb.GetPersonImei(imei) == "available"{
		contact.First_Name = firstName
		contact.Last_Name = lastName
		contact.Phone_Number = phone
		contact.Imei = imei

		idb.DB.Create(&contact)

		result = gin.H{
			"result": contact,
		}

		c.JSON(http.StatusOK, result)
	}else{
		c.JSON(http.StatusUnauthorized, "No registered Imei")
	}
	
}

//to get one data with {id}
func (idb *InDB) GetContact(c *gin.Context) {
	var (
		message structs.Contact
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
func (idb *InDB) GetContactImei(c *gin.Context) {
	var (
		message structs.Contact
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
func (idb *InDB) GetAllContact(c *gin.Context) {

	var (
		message []structs.Contact
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