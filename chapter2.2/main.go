package main

import (
	"chapter2-2/database"
	"chapter2-2/routers"

	_ "github.com/lib/pq"
)

func main() {
	db := database.GetConnection()
	defer db.Close()

	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}
