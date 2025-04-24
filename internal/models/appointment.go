package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceType defines the type of appointment/service provided.
type ServiceType string

const (
	ServiceConsultation ServiceType = "consulta"
	ServiceEvaluation   ServiceType = "avaliação"
	ServiceWeighing     ServiceType = "pesagem"
	ServiceFollowUp     ServiceType = "retorno"
	ServiceOther        ServiceType = "outro"
)

// PaymentMethod defines how an appointment was paid for.
type PaymentMethod string

const (
	PaymentCash         PaymentMethod = "dinheiro"
	PaymentCreditCard   PaymentMethod = "cartão de crédito"
	PaymentDebitCard    PaymentMethod = "cartão de débito"
	PaymentPIX          PaymentMethod = "pix"
	PaymentBankTransfer PaymentMethod = "transferência bancária"
	PaymentCheck        PaymentMethod = "cheque"
	PaymentInsurance    PaymentMethod = "convênio"
	PaymentOther        PaymentMethod = "outro"
)

// Appointment holds details about a single patient visit/interaction.
type Appointment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID     primitive.ObjectID `bson:"patient_id" json:"patientId"`
	Date          time.Time     `json:"date,omitempty"`
	Time          string        `json:"time,omitempty"`
	ServiceType   ServiceType   `json:"serviceType,omitempty"`
	Value         float64       `json:"value,omitempty"`
	PaymentMethod PaymentMethod `json:"paymentMethod,omitempty"`
	Notes         string        `json:"notes,omitempty"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}
