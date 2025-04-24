package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"renutri/internal/models"
)

// AppointmentRepository defines CRUD for appointment records.
type AppointmentRepository interface {
	CreateAppointment(app *models.Appointment) error
	GetAppointmentByID(id string) (*models.Appointment, error)
	UpdateAppointment(app *models.Appointment) error
	DeleteAppointment(id string) error
	ListAppointments() ([]models.Appointment, error)
	ListByPatient(patientID string) ([]models.Appointment, error)
}

type appointmentRepo struct {
	collection *mongo.Collection
}

func NewAppointmentRepository(collection *mongo.Collection) AppointmentRepository {
	return &appointmentRepo{collection}
}

func (r *appointmentRepo) CreateAppointment(app *models.Appointment) error {
	if app.ID == primitive.NilObjectID {
		app.ID = primitive.NewObjectID()
	}
	ctx := context.TODO()
	_, err := r.collection.InsertOne(ctx, app)
	return err
}

func (r *appointmentRepo) GetAppointmentByID(id string) (*models.Appointment, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var app models.Appointment
	ctx := context.TODO()
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&app)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("appointment not found")
		}
		return nil, err
	}
	return &app, nil
}

func (r *appointmentRepo) UpdateAppointment(app *models.Appointment) error {
	ctx := context.TODO()
	filter := bson.M{"_id": app.ID}
	update := bson.M{"$set": app}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *appointmentRepo) DeleteAppointment(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *appointmentRepo) ListAppointments() ([]models.Appointment, error) {
	ctx := context.TODO()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var apps []models.Appointment
	for cursor.Next(ctx) {
		var app models.Appointment
		if err := cursor.Decode(&app); err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}
	return apps, nil
}

func (r *appointmentRepo) ListByPatient(patientID string) ([]models.Appointment, error) {
	objID, err := primitive.ObjectIDFromHex(patientID)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	cursor, err := r.collection.Find(ctx, bson.M{"patient_id": objID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var apps []models.Appointment
	for cursor.Next(ctx) {
		var app models.Appointment
		if err := cursor.Decode(&app); err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}
	return apps, nil
}
