package models

import "gorm.io/gorm"

type Lobby struct{
  gorm.Model
  Name string
}

func CreateLobby(name string) Lobby{
  return Lobby{
    Name: name,
  }
}
