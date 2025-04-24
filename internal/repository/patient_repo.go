package repository

import (
    "context"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "renutri/internal/models"
)

// PatientRepository defines CRUD operations for patients.
type PatientRepository interface {
    CreatePatient(patient *models.Patient) error
    GetPatientByID(id string) (*models.Patient, error)
    UpdatePatient(patient *models.Patient) error
    DeletePatient(id string) error
    ListPatients() ([]models.Patient, error)
}

type patientRepo struct {
    collection *mongo.Collection
}

// NewPatientRepository returns a new PatientRepository using MongoDB.
func NewPatientRepository(collection *mongo.Collection) PatientRepository {
    return &patientRepo{collection}
}

func (r *patientRepo) CreatePatient(patient *models.Patient) error {
    if patient.ID == primitive.NilObjectID {
        patient.ID = primitive.NewObjectID()
    }
    ctx := context.TODO()
    _, err := r.collection.InsertOne(ctx, patient)
    return err
}

func (r *patientRepo) GetPatientByID(id string) (*models.Patient, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    var patient models.Patient
    ctx := context.TODO()
    err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&patient)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("patient not found")
        }
        return nil, err
    }
    return &patient, nil
}

func (r *patientRepo) UpdatePatient(patient *models.Patient) error {
    ctx := context.TODO()
    filter := bson.M{"_id": patient.ID}
    update := bson.M{"$set": patient}
    _, err := r.collection.UpdateOne(ctx, filter, update)
    return err
}

func (r *patientRepo) DeletePatient(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    ctx := context.TODO()
    _, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
    return err
}

func (r *patientRepo) ListPatients() ([]models.Patient, error) {
    ctx := context.TODO()
    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    var patients []models.Patient
    for cursor.Next(ctx) {
        var patient models.Patient
        if err := cursor.Decode(&patient); err != nil {
            return nil, err
        }
        patients = append(patients, patient)
    }
    return patients, nil
}