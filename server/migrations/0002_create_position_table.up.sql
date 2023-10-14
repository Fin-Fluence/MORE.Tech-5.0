CREATE TABLE IF NOT EXISTS app.position (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    latitude NUMERIC(9, 6) NOT NULL,
    longitude NUMERIC(9, 6) NOT NULL
);