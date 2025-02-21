package main

import (
	"github.com/livghit/impostor/server/database"
	"github.com/livghit/impostor/server/models"
	"github.com/livghit/impostor/utils"
)

func main() {
	databasePath := utils.DatabasePath()
	db := database.New(databasePath)

	models := []interface{}{
		models.Player{},
		models.Game{},
		models.Lobby{},
	}
	db.Conn.AutoMigrate(models...)
}
