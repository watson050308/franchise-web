package main

import (
	"franchise-web/config/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	// initializers.DB.Migrator().DropTable(&model.User{})
}
