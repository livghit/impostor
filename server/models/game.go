package models

import "gorm.io/gorm"

type Game struct{
   gorm.Model
   Name string
}

func CreateGame (name string) Game{
  return Game{
    Name: name,
  }
}
