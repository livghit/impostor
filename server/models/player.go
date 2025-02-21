package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name string
}

func CreatePlayer(name string) Player {
	return Player{
		Name: name,
	}
}
