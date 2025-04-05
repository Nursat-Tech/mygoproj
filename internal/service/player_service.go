package service

import (
	"github.com/nursat/myproj/internal/models"
	"github.com/nursat/myproj/internal/repository"
)

type PlayerService struct {
	repo *repository.PlayerRepository
}

func NewPlayerService(repo *repository.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) CreatePlayer(player *models.Player) error {
	return s.repo.CreatePlayer(player)
}

func (s *PlayerService) GetPlayers() ([]models.Player, error) {
	return s.repo.GetPlayers()
}

func (s *PlayerService) GetPlayerByID(id uint) (*models.Player, error) {
	return s.repo.GetPlayerByID(id)
}

func (s *PlayerService) UpdatePlayer(player *models.Player) error {
	return s.repo.UpdatePlayer(player)
}

func (s *PlayerService) DeletePlayer(id uint) error {
	return s.repo.DeletePlayer(id)
}
