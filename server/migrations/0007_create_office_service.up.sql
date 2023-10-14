CREATE TABLE IF NOT EXISTS app.office_service (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(512) NOT NULL,
    capability BOOLEAN,
    activity BOOLEAN,
    current_ticket VARCHAR(8) NOT NULL,
    last_ticket VARCHAR(8) NOT NULL,
    office_id UUID REFERENCES app.office(id) ON DELETE CASCADE NOT NULL
);