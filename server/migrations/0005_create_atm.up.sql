CREATE TABLE IF NOT EXISTS app.atm (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    position_id UUID REFERENCES app.position(id) ON DELETE RESTRICT NOT NULL,
    address VARCHAR(255) NOT NULL,
    all_day BOOLEAN NOT NULL
);