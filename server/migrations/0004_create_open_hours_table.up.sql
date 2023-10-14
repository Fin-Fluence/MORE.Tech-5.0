CREATE TABLE IF NOT EXISTS app.open_hours (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    day VARCHAR(50) NOT NULL,
    hours VARCHAR(50) NOT NULL,
    is_individual BOOLEAN NOT NULL,
    office_id UUID REFERENCES app.office(id) ON DELETE CASCADE NOT NULL
);