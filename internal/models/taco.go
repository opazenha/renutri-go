package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// UnifiedFoodItem represents a food item with all available nutritional information
// combined from the various spreadsheets.
type UnifiedFoodItem struct {
	ID                  primitive.ObjectID `bson:"id" json:"id"`
	Descricao           string             `bson:"alimento_descricao" json:"alimento_descricao"`
	UmidadePercent    *float64             `bson:"umidade_percent,omitempty" json:"umidade_percent,omitempty"`
	EnergiaKcal       *int                 `bson:"energia_kcal,omitempty" json:"energia_kcal,omitempty"`
	EnergiaKj         *int                 `bson:"energia_kj,omitempty" json:"energia_kj,omitempty"`
	ProteinaG         *float64             `bson:"proteina_g,omitempty" json:"proteina_g,omitempty"`
	LipidiosG         *float64             `bson:"lipidios_g,omitempty" json:"lipidios_g,omitempty"`
	ColesterolMg      *float64             `bson:"colesterol_mg,omitempty" json:"colesterol_mg,omitempty"`
	CarboidratoG      *float64             `bson:"carboidrato_g,omitempty" json:"carboidrato_g,omitempty"`
	FibraAlimentarG   *float64             `bson:"fibra_alimentar_g,omitempty" json:"fibra_alimentar_g,omitempty"`
	CinzasG           *float64             `bson:"cinzas_g,omitempty" json:"cinzas_g,omitempty"`
	CalcioMg          *int                 `bson:"calcio_mg,omitempty" json:"calcio_mg,omitempty"`
	MagnesioMg        *int                 `bson:"magnesio_mg,omitempty" json:"magnesio_mg,omitempty"`
	ManganesMg        *float64             `bson:"manganes_mg,omitempty" json:"manganes_mg,omitempty"`
	FosforoMg         *int                 `bson:"fosforo_mg,omitempty" json:"fosforo_mg,omitempty"`
	FerroMg           *float64             `bson:"ferro_mg,omitempty" json:"ferro_mg,omitempty"`
	SodioMg           *float64             `bson:"sodio_mg,omitempty" json:"sodio_mg,omitempty"` // Using float64 and pointer for consistency and Tr/NA
	PotassioMg        *float64             `bson:"potassio_mg,omitempty" json:"potassio_mg,omitempty"` // Using float64 and pointer for consistency and Tr/NA
	CobreMg           *float64             `bson:"cobre_mg,omitempty" json:"cobre_mg,omitempty"`   // Using pointer for Tr/NA
	ZincoMg           *float64             `bson:"zinco_mg,omitempty" json:"zinco_mg,omitempty"`   // Using pointer for Tr/NA
	RetinolMcg        *float64             `bson:"retinol_mcg,omitempty" json:"retinol_mcg,omitempty"` // Using pointer for NA/Tr
	REMcg             *float64             `bson:"re_mcg,omitempty" json:"re_mcg,omitempty"`           // Using pointer for NA/Tr
	RAEMcg            *float64             `bson:"rae_mcg,omitempty" json:"rae_mcg,omitempty"`         // Using pointer for NA/Tr
	TiaminaMg         *float64             `bson:"tiamina_mg,omitempty" json:"tiamina_mg,omitempty"` // Using pointer for Tr/NA
	RiboflavinaMg     *float64             `bson:"riboflavina_mg,omitempty" json:"riboflavina_mg,omitempty"` // Using pointer for Tr/NA
	PiridoxinaMg      *float64             `bson:"piridoxina_mg,omitempty" json:"piridoxina_mg,omitempty"` // Using pointer for Tr/NA
	NiacinaMg         *float64             `bson:"niacina_mg,omitempty" json:"niacina_mg,omitempty"` // Using pointer for Tr/NA
	VitaminaCMg       *float64             `bson:"vitamina_c_mg,omitempty" json:"vitamina_c_mg,omitempty"` // Using pointer for NA/Tr
	TriptofanoG       *float64             `bson:"triptofano_g,omitempty" json:"triptofano_g,omitempty"`
	TreoninaG         *float64             `bson:"treonina_g,omitempty" json:"treonina_g,omitempty"`
	IsoleucinaG       *float64             `bson:"isoleucina_g,omitempty" json:"isoleucina_g,omitempty"`
	LeucinaG          *float64             `bson:"leucina_g,omitempty" json:"leucina_g,omitempty"`
	LisinaG           *float64             `bson:"lisina_g,omitempty" json:"lisina_g,omitempty"`
	MetioninaG        *float64             `bson:"metionina_g,omitempty" json:"metionina_g,omitempty"`
	CistinaG          *float64             `bson:"cistina_g,omitempty" json:"cistina_g,omitempty"`
	FenilalaninaG     *float64             `bson:"fenilalanina_g,omitempty" json:"fenilalanina_g,omitempty"`
	TirosinaG         *float64             `bson:"tirosina_g,omitempty" json:"tirosina_g,omitempty"`
	ValinaG           *float64             `bson:"valina_g,omitempty" json:"valina_g,omitempty"`
	ArgininaG         *float64             `bson:"arginina_g,omitempty" json:"arginina_g,omitempty"`
	HistidinaG        *float64             `bson:"histidina_g,omitempty" json:"histidina_g,omitempty"`
	AlaninaG          *float64             `bson:"alanina_g,omitempty" json:"alanina_g,omitempty"`
	AcidoAsparticoG   *float64             `bson:"acido_aspartico_g,omitempty" json:"acido_aspartico_g,omitempty"`
	AcidoGlutamicoG   *float64             `bson:"acido_glutamico_g,omitempty" json:"acido_glutamico_g,omitempty"`
	GlicinaG          *float64             `bson:"glicina_g,omitempty" json:"glicina_g,omitempty"`
	ProlinaG          *float64             `bson:"prolina_g,omitempty" json:"prolina_g,omitempty"`
	SerinaG           *float64             `bson:"serina_g,omitempty" json:"serina_g,omitempty"`
	MongoID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}