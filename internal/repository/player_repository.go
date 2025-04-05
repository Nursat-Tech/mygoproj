package repository

import (
	"github.com/nursat/myproj/internal/models"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) CreatePlayer(player *models.Player) error {
	return r.db.Create(player).Error
}

func (r *PlayerRepository) GetPlayers() ([]models.Player, error) {
	var players []models.Player
	err := r.db.Find(&players).Error
	return players, err
}

func (r *PlayerRepository) GetPlayerByID(id uint) (*models.Player, error) {
	var player models.Player
	err := r.db.First(&player, id).Error
	return &player, err
}

func (r *PlayerRepository) UpdatePlayer(player *models.Player) error {
	return r.db.Save(player).Error
}

func (r *PlayerRepository) DeletePlayer(id uint) error {
	return r.db.Delete(&models.Player{}, id).Error
}
