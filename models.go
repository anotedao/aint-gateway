package main

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TxID      string `gorm:"size:255;uniqueIndex"`
	Type      string `gorm:"type:string"`
	Processed bool
}
