package main

import (
	"log"

	"github.com/eduardohslfreire/animalia-api/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	api := new(api.Server)
	api.StartServer()
}

/*
package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   uint   `uri:"id" binding:"number"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}*/

/*
// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" query:"password" xml:"password" binding:"required"`
	//ID       int    `form:"id" query:"id" xml:"id" binding:"required,numeric"`
}

func main() {
	router := gin.Default()

	var trans ut.Translator

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt := pt.New()
		uni := ut.New(pt, pt)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, _ = uni.GetTranslator("pt")
		pt_translations.RegisterDefaultTranslations(v, trans)
	}

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.GET("/loginJSON/:id", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindQuery(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.(validator.ValidationErrors).Translate(trans)})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>manu</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
*/
