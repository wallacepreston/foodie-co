// database/db.go

package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    DB = db
    return db, nil
}

type Recipe struct {
    gorm.Model
    Name        string
    Instructions string
}
