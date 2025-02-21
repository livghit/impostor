package models

import "gorm.io/gorm"

type Lobby struct{
  gorm.Model
  Name string
}
