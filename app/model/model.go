package model

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type FamousPerson struct {
  gorm.Model
  Name   string `gorm:"unique" json:"name"`
  Bio    string `json:"bio"`
  Dob    string `json:"dob"`
}

// DBMigrate will create and migrate the tables, and then make some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
  db.AutoMigrate(&FamousPerson{})
  return db
}
