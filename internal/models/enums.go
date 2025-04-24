package models

// --- Enum Definitions ---

// Gender represents the sex of the patient.
type Gender string

const (
	Male   Gender = "masculino"
	Female Gender = "feminino"
	Other  Gender = "outro"
)

// MaritalStatus represents the patient's marital status.
type MaritalStatus string

const (
	Single         MaritalStatus = "solteiro(a)"
	Married        MaritalStatus = "casado(a)"
	Divorced       MaritalStatus = "divorciado(a)"
	Widowed        MaritalStatus = "viúvo(a)"
	StableUnion    MaritalStatus = "união estável"
	MaritalUnknown MaritalStatus = "desconhecido"
)

// EducationLevel represents the patient's education level.
type EducationLevel string

const (
	IncompleteElementary EducationLevel = "fundamental incompleto"
	CompleteElementary   EducationLevel = "fundamental completo"
	IncompleteHighSchool EducationLevel = "médio incompleto"
	CompleteHighSchool   EducationLevel = "médio completo"
	IncompleteHigherEd   EducationLevel = "superior incompleto"
	CompleteHigherEd     EducationLevel = "superior completo"
	PostGraduate         EducationLevel = "pós-graduação"
	EducationUnknown     EducationLevel = "desconhecido"
)

// BowelFunction describes the patient's typical bowel function.
type BowelFunction string

const (
	BowelNormal      BowelFunction = "normal"
	BowelConstipated BowelFunction = "obstipado"
	BowelDiarrheal   BowelFunction = "diarreico"
	BowelMixed       BowelFunction = "misto"
	BowelUnknown     BowelFunction = "desconhecido"
)

// UrineColor describes the typical urine color.
type UrineColor string

const (
	UrineClear       UrineColor = "clara"
	UrineLightYellow UrineColor = "amarela clara"
	UrineDarkYellow  UrineColor = "amarela escura"
	UrineOrange      UrineColor = "laranja"
	UrineOther       UrineColor = "outra"
	UrineUnknown     UrineColor = "desconhecida"
)

// StressLevel indicates the patient's perceived stress level.
type StressLevel string

const (
	StressNone     StressLevel = "nenhum"
	StressLow      StressLevel = "pouco"
	StressModerate StressLevel = "moderado"
	StressHigh     StressLevel = "alto"
	StressExtreme  StressLevel = "extremo"
	StressUnknown  StressLevel = "desconhecido"
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
