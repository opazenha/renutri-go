package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// UnifiedFoodItem represents a food item with all available nutritional information
// combined from the various spreadsheets.
type UnifiedFoodItem struct {
	MongoID             primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	ID                  int                 `bson:"id" json:"id"`
	Descricao           string              `bson:"alimento_descricao" json:"alimento_descricao"`
	UmidadePercent      *float64            `bson:"umidade_percent,omitempty" json:"umidade_percent,omitempty"`
	EnergiaKcal         *int                `bson:"energia_kcal,omitempty" json:"energia_kcal,omitempty"`
	EnergiaKj           *int                `bson:"energia_kj,omitempty" json:"energia_kj,omitempty"`
	ProteinaG           *float64            `bson:"proteina_g,omitempty" json:"proteina_g,omitempty"`
	LipidiosG           *float64            `bson:"lipidios_g,omitempty" json:"lipidios_g,omitempty"`
	ColesterolMg        *float64            `bson:"colesterol_mg,omitempty" json:"colesterol_mg,omitempty"`
	CarboidratoG        *float64            `bson:"carboidrato_g,omitempty" json:"carboidrato_g,omitempty"`
	FibraAlimentarG     *float64            `bson:"fibra_alimentar_g,omitempty" json:"fibra_alimentar_g,omitempty"`
	CinzasG             *float64            `bson:"cinzas_g,omitempty" json:"cinzas_g,omitempty"`
	CalcioMg            *float64            `bson:"calcio_mg,omitempty" json:"calcio_mg,omitempty"`
	MagnesioMg          *float64            `bson:"magnesio_mg,omitempty" json:"magnesio_mg,omitempty"`
	ManganesMg          *float64            `bson:"manganes_mg,omitempty" json:"manganes_mg,omitempty"`
	FosforoMg           *float64            `bson:"fosforo_mg,omitempty" json:"fosforo_mg,omitempty"`
	FerroMg             *float64            `bson:"ferro_mg,omitempty" json:"ferro_mg,omitempty"`
	SodioMg             *float64            `bson:"sodio_mg,omitempty" json:"sodio_mg,omitempty"`
	PotassioMg          *float64            `bson:"potassio_mg,omitempty" json:"potassio_mg,omitempty"`
	CobreMg             *float64            `bson:"cobre_mg,omitempty" json:"cobre_mg,omitempty"`
	ZincoMg             *float64            `bson:"zinco_mg,omitempty" json:"zinco_mg,omitempty"`
	RetinolMcg          *float64            `bson:"retinol_mcg,omitempty" json:"retinol_mcg,omitempty"`
	REMcg               *float64            `bson:"re_mcg,omitempty" json:"re_mcg,omitempty"`
	RAEMcg              *float64            `bson:"rae_mcg,omitempty" json:"rae_mcg,omitempty"`
	TiaminaMg           *float64            `bson:"tiamina_mg,omitempty" json:"tiamina_mg,omitempty"`
	RiboflavinaMg       *float64            `bson:"riboflavina_mg,omitempty" json:"riboflavina_mg,omitempty"`
	PiridoxinaMg        *float64            `bson:"piridoxina_mg,omitempty" json:"piridoxina_mg,omitempty"`
	NiacinaMg           *float64            `bson:"niacina_mg,omitempty" json:"niacina_mg,omitempty"`
	VitaminaCMg         *float64            `bson:"vitamina_c_mg,omitempty" json:"vitamina_c_mg,omitempty"`
	Categoria           string              `bson:"categoria,omitempty" json:"categoria,omitempty"`
	// Fatty acids
	Fatty120            *float64            `bson:"12:0,omitempty" json:"12:0,omitempty"`
	Fatty140            *float64            `bson:"14:0,omitempty" json:"14:0,omitempty"`
	Fatty160            *float64            `bson:"16:0,omitempty" json:"16:0,omitempty"`
	Fatty180            *float64            `bson:"18:0,omitempty" json:"18:0,omitempty"`
	Fatty200            *float64            `bson:"20:0,omitempty" json:"20:0,omitempty"`
	Fatty220            *float64            `bson:"22:0,omitempty" json:"22:0,omitempty"`
	Fatty240            *float64            `bson:"24:0,omitempty" json:"24:0,omitempty"`
	Fatty141            *float64            `bson:"14:1,omitempty" json:"14:1,omitempty"`
	Fatty161            *float64            `bson:"16:1,omitempty" json:"16:1,omitempty"`
	Fatty181            *float64            `bson:"18:1,omitempty" json:"18:1,omitempty"`
	Fatty201            *float64            `bson:"20:1,omitempty" json:"20:1,omitempty"`
	Fatty186n6          *float64            `bson:"18:2 n-6,omitempty" json:"18:2 n-6,omitempty"`
	Fatty183n3          *float64            `bson:"18:3 n-3,omitempty" json:"18:3 n-3,omitempty"`
	Fatty204            *float64            `bson:"20:4,omitempty" json:"20:4,omitempty"`
	Fatty205            *float64            `bson:"20:5,omitempty" json:"20:5,omitempty"`
	Fatty225            *float64            `bson:"22:5,omitempty" json:"22:5,omitempty"`
	Fatty226            *float64            `bson:"22:6,omitempty" json:"22:6,omitempty"`
	Fatty181t           *float64            `bson:"18:1t,omitempty" json:"18:1t,omitempty"`
	Fatty182t           *float64            `bson:"18:2t,omitempty" json:"18:2t,omitempty"`
}