
-- Initial SQL schema for nutritionist app

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    profession TEXT,
    address TEXT,
    phones TEXT[],
    email TEXT
);

CREATE TABLE health_history (
    id SERIAL PRIMARY KEY,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    condition_name TEXT NOT NULL,
    diagnosis_date TIMESTAMP,
    status TEXT NOT NULL,
    medication TEXT
);

CREATE TABLE habits (
    id SERIAL PRIMARY KEY,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    usual_diet TEXT,
    dietary_recall TEXT,
    frequency_of_consumption TEXT,
    dietary_restrictions TEXT,
    supplementation TEXT,
    hydration DOUBLE PRECISION,
    alcohol_consumption TEXT,
    smoking_status BOOLEAN,
    sleep_pattern TEXT,
    physical_activity TEXT,
    daily_routine TEXT
);

CREATE TABLE clinical_evaluations (
    id SERIAL PRIMARY KEY,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    evaluation_date TIMESTAMP NOT NULL,
    symptoms TEXT,
    laboratory_results TEXT
);

CREATE TABLE anthropometric_assessments (
    id SERIAL PRIMARY KEY,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    assessment_date TIMESTAMP NOT NULL,
    weight DOUBLE PRECISION,
    height DOUBLE PRECISION,
    ideal_weight DOUBLE PRECISION,
    bmi DOUBLE PRECISION,
    bmi_classification TEXT,
    waist DOUBLE PRECISION,
    hip DOUBLE PRECISION,
    arm DOUBLE PRECISION,
    circumference_classification TEXT,
    triceps DOUBLE PRECISION,
    biceps DOUBLE PRECISION,
    subscapular DOUBLE PRECISION,
    suprailiac DOUBLE PRECISION,
    skinfold_classification TEXT,
    body_fat_percentage DOUBLE PRECISION
);

CREATE TABLE nutritional_needs_calculations (
    id SERIAL PRIMARY KEY,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    calculation_date TIMESTAMP NOT NULL,
    tmb DOUBLE PRECISION,
    geb DOUBLE PRECISION,
    get DOUBLE PRECISION,
    vet DOUBLE PRECISION,
    carbohydrates DOUBLE PRECISION,
    protein DOUBLE PRECISION,
    lipids DOUBLE PRECISION,
    micronutrient_needs TEXT,
    water_intake_ml DOUBLE PRECISION,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE diet_plans (
    id SERIAL PRIMARY KEY,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    plan_date TIMESTAMP NOT NULL,
    description TEXT,
    portions TEXT,
    food_groups TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE meals (
    id SERIAL PRIMARY KEY,
    diet_plan_id INTEGER NOT NULL REFERENCES diet_plans(id) ON DELETE CASCADE,
    meal_time TEXT,
    meal_description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE food_substitutions (
    id SERIAL PRIMARY KEY,
    meal_id INTEGER NOT NULL REFERENCES meals(id) ON DELETE CASCADE,
    original_food TEXT,
    substitute_food TEXT,
    household_measure TEXT,
    weight_grams DOUBLE PRECISION,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    date TIMESTAMP,
    time TEXT,
    service_type TEXT,
    value DOUBLE PRECISION,
    payment_method TEXT,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE guidance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    date TIMESTAMP NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE generated_documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    document_date TIMESTAMP NOT NULL,
    document_type TEXT,
    file_path TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);