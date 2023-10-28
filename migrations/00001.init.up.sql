CREATE TABLE IF NOT EXISTS components (
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    name TEXT NOT NULL,  
    data JSONB NOT NULL,
    UNIQUE(type, name),
    UNIQUE(data)
);

CREATE INDEX idx_components_data on components USING gin(data);

CREATE TABLE IF NOT EXISTS branches (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS commits (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    branch_id INTEGER REFERENCES branches(id),
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS commit_components (
    commit_id INTEGER REFERENCES commits(id),
    component_id INTEGER REFERENCES components(id),
    PRIMARY KEY (commit_id, component_id)
);
