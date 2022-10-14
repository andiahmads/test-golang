package game

import "gorm.io/gorm"

type Repository interface {
	Insert(data DataGame) (DataGame, error)
	ShowDataGame() ([]DataGame, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(data DataGame) (DataGame, error) {

	err := r.db.Debug().Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ShowDataGame() ([]DataGame, error) {
	var dataGame []DataGame

	err := r.db.Find(&dataGame).Error
	if err != nil {
		return dataGame, err
	}

	return dataGame, nil
}
