package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"../structs"
	"github.com/gin-gonic/gin"
)

//to get one data with {id}
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)

}

func (idb *InDB) GetPersonImei(imei string) string {
	var (
		person structs.Person
	)
	//id := c.Param("id")
	err := idb.DB.Where("Imei=?", imei).First(&person).Error
	result := ""
	if err != nil {
		result = "error"
	} else {
		result = "available"
	}

	return result

}

func (idb *InDB) GetPersonApi(api string) string {
	var (
		person structs.Person
	)
	//id := c.Param("id")
	err := idb.DB.Where("Api=?", api).First(&person).Error
	result := ""
	if err != nil {
		result = "error"
	} else {
		result = "available"
	}

	return result

}

func (idb *InDB) GetAPI(tokenString string) string {
	var (
		person structs.Person
		result string
	)
	//tokenString := c.Request.Header.Get("Authorization")
	fmt.Printf("Token ada di con: %v", tokenString)
	err := idb.DB.Where("Api=?", tokenString).First(&person).Error
	if err != nil {
		result = "Error"
	} else {
		result = person.Imei
	}
	return result
}

// func (idb *InDB) GetAPIAdmin(tokenString string) string {
// 	var (
// 		person structs.Administrator
// 		result string
// 	)
// 	//tokenString := c.Request.Header.Get("Authorization")
// 	fmt.Printf("Token ada di con: %v", tokenString)
// 	err := idb.DB.Where("Password=?", tokenString).First(&person).Error
// 	if err != nil {
// 		result = "Error"
// 	} else {
// 		result = person.Imei
// 	}
// 	return result
// }

//get all data in person
func (idb *InDB) GetPersons(c *gin.Context) {

	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

//create new data to the database
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person

		result gin.H
	)
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	phone := c.PostForm("phone_number")
	imei := c.PostForm("imei")
	modelName := c.PostForm("model_name")
	modelNumber := c.PostForm("model_number")
	softwareVer := c.PostForm("software_ver")
	apiKey := authGenerator(imei)

	person.First_Name = firstName
	person.Last_Name = lastName
	person.Phone_Number = phone
	person.Model_Number = modelNumber
	person.Model_Name = modelName
	person.Software_Ver = softwareVer
	person.API = apiKey
	person.Imei = imei

	idb.DB.Create(&person)

	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func authGenerator(imei string) string {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte(imei))
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Print(token)
	return token
}
