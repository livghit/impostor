package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/livghit/impostor/server/database"
	"github.com/livghit/impostor/server/models"
)


func main(){
  homeDir, err := os.UserHomeDir();if err != nil{
    log.Fatal(err)
  }
  db := database.New(filepath.Join(homeDir,"Development","OpenSource","impostor","test.db"))

  models := []interface{}{
    models.Player{},
    models.Game{},
    models.Lobby{},
  }
  db.Conn.AutoMigrate(models...)
}
