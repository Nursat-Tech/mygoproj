package service

import (
	"github.com/nursat/myproj/internal/models"
	"gorm.io/gorm"
)

type PlayerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) *PlayerService {
	return &PlayerService{db: db}
}

func (s *PlayerService) CreatePlayer(player *models.Player) error {
	return s.db.Create(player).Error
}

func (s *PlayerService) GetPlayers() ([]models.Player, error) {
	var players []models.Player
	err := s.db.Find(&players).Error
	return players, err
}

func (s *PlayerService) GetPlayerByID(id uint) (*models.Player, error) {
	var player models.Player
	err := s.db.First(&player, id).Error
	return &player, err
}

func (s *PlayerService) UpdatePlayer(player *models.Player) error {
	return s.db.Save(player).Error
}

func (s *PlayerService) DeletePlayer(id uint) error {
	return s.db.Delete(&models.Player{}, id).Error
}
