package repository

import "gorm.io/gorm"

type CharacterRepository interface {
}

type characterRepository struct {
    DB *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) CharacterRepository {
    return characterRepository{
        DB: db,
    }
}
