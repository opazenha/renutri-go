package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

type HabitService struct {
	repo repository.HabitRepository
}

func NewHabitService(repo repository.HabitRepository) *HabitService {
	return &HabitService{repo: repo}
}

func (s *HabitService) Create(a *models.Habit) error {
	return s.repo.CreateHabit(a)
}

func (s *HabitService) GetByID(id string) (*models.Habit, error) {
	return s.repo.GetHabitByID(id)
}

func (s *HabitService) Update(a *models.Habit) error {
	return s.repo.UpdateHabit(a)
}

func (s *HabitService) Delete(id string) error {
	return s.repo.DeleteHabit(id)
}

func (s *HabitService) List() ([]models.Habit, error) {
	return s.repo.ListHabits()
}
