package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	collection *mongo.Collection
}

func NewHabitRepository(collection *mongo.Collection) HabitRepository {
	return &habitRepo{collection}
}

func (r *habitRepo) CreateHabit(a *models.Habit) error {
	if a.ID == primitive.NilObjectID {
		a.ID = primitive.NewObjectID()
	}
	ctx := context.TODO()
	_, err := r.collection.InsertOne(ctx, a)
	return err
}

func (r *habitRepo) GetHabitByID(id string) (*models.Habit, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var a models.Habit
	ctx := context.TODO()
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&a)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("habit not found")
		}
		return nil, err
	}
	return &a, nil
}

func (r *habitRepo) UpdateHabit(a *models.Habit) error {
	ctx := context.TODO()
	filter := bson.M{"_id": a.ID}
	update := bson.M{"$set": a}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *habitRepo) DeleteHabit(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *habitRepo) ListHabits() ([]models.Habit, error) {
	ctx := context.TODO()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var habits []models.Habit
	for cursor.Next(ctx) {
		var habit models.Habit
		if err := cursor.Decode(&habit); err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}
	return habits, nil
}
