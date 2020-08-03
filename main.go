package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var (
	router *gin.Engine
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func init() {
	router = gin.Default()

	router.Use(CORSMiddleware())

	router.Static("/static", "./files")
	router.Handle("GET", "/", Key)
	router.Handle("POST", "/", Print)
}

func Print(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(body))
}

func Key(c *gin.Context) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Length", "16")
	c.Writer.Write(decode("MyDH5i0pKq/J50xOaCZaHw=="))
}

func decode(s string) []byte {
	decodeString, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("a.key", decodeString, 0644)

	return decodeString
}

func main() {
	if err := router.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
