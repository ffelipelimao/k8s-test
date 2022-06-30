package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router := app.Group("/")

	router.GET("/", Hello)
	router.GET("/configmap", configMap)
	router.GET("/secret", secret)

	app.Run(":8080")

}

func Hello(c *gin.Context) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(c.Writer, "Hello i am %s. I am %s", name, age)
}

func configMap(c *gin.Context) {

	data, err := ioutil.ReadFile("/usr/src/app/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error:", err)
	}

	fmt.Fprintf(c.Writer, "My family %s", string(data))

}

func secret(c *gin.Context) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(c.Writer, "user %s and password %s", user, password)
}
