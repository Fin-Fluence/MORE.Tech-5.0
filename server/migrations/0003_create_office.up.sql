CREATE TABLE IF NOT EXISTS app.office (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sale_point_name VARCHAR(255) NOT NULL,
    address VARCHAR(400) NOT NULL,
    status VARCHAR(255) NOT NULL,
    rko BOOLEAN,
    type VARCHAR(255) NOT NULL,
    sale_point_format VARCHAR(255),
    suo_availability BOOLEAN,
    has_ramp BOOLEAN,
    position_id UUID REFERENCES app.position(id) ON DELETE RESTRICT NOT NULL,
    distance INT NOT NULL,
    kep BOOLEAN,
    my_branch BOOLEAN NOT NULL
);