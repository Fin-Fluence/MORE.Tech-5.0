CREATE TABLE IF NOT EXISTS app.service (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(512) NOT NULL,
    capability BOOLEAN,
    activity BOOLEAN,
    atm_id UUID REFERENCES app.atm(id) ON DELETE CASCADE NOT NULL
);