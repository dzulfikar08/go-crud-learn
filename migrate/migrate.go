package main

import (
	"github.com/dzulfikar08/go-learning-crud/initializers"
	"github.com/dzulfikar08/go-learning-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
