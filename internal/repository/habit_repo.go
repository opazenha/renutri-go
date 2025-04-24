package repository

import (
	"gorm.io/gorm"
	"renutri/internal/models"
)

type HabitRepository interface {
	CreateHabit(a *models.Habit) error
	GetHabitByID(id string) (*models.Habit, error)
	UpdateHabit(a *models.Habit) error
	DeleteHabit(id string) error
	ListHabits() ([]models.Habit, error)
}

type habitRepo struct {
	db *gorm.DB
}

func NewHabitRepository(db *gorm.DB) HabitRepository {
	return &habitRepo{db: db}
}

func (r *habitRepo) CreateHabit(a *models.Habit) error {
	return r.db.Create(a).Error
}

func (r *habitRepo) GetHabitByID(id string) (*models.Habit, error) {
	var a models.Habit
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *habitRepo) UpdateHabit(a *models.Habit) error {
	return r.db.Save(a).Error
}

func (r *habitRepo) DeleteHabit(id string) error {
	return r.db.Delete(&models.Habit{}, "id = ?", id).Error
}

func (r *habitRepo) ListHabits() ([]models.Habit, error) {
	var items []models.Habit
	err := r.db.Find(&items).Error
	return items, err
}
