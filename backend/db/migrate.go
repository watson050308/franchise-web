package main

import (
	model "franchise-web/app/models"
	"franchise-web/config/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.Migrator().DropTable(&model.User{})
}
