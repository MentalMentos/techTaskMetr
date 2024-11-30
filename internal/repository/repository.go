package repository

import "gorm.io/gorm"

type RepoImpl struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *RepoImpl {
	return &RepoImpl{
		DB: DB,
	}
}
