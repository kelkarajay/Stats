CREATE TABLE IF NOT EXISTS apps (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    "name" TEXT,
    "description" TEXT    
);

CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    app_id UUID NOT NULL,
    client_id VARCHAR(36),
    "namespace" TEXT,
    "name" TEXT,
    category TEXT,
    "data" JSON,
    CONSTRAINT fk_app FOREIGN KEY(app_id) REFERENCES apps(id)
);