package main

import (
	db "Assignment-2_AzmiFarisM/database"
	"Assignment-2_AzmiFarisM/routers"
)

func main() {
	db.Init()
	routers.ServerOn().Run(":8080")
}
