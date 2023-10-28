-- Create endpoints table
CREATE TABLE IF NOT EXISTS endpoints (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL,
    method TEXT NOT NULL
);

-- Create models table
CREATE TABLE IF NOT EXISTS models (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Create actions table
CREATE TABLE IF NOT EXISTS actions (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Create branches table
CREATE TABLE IF NOT EXISTS branches (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    status TEXT DEFAULT 'active' NOT NULL -- for example: active, merged, etc.
);

-- Create commits table with reference to branches
CREATE TABLE IF NOT EXISTS commits (
    id SERIAL PRIMARY KEY,
    branch_id INTEGER REFERENCES branches(id),
    timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Junction table for commits and endpoints
CREATE TABLE IF NOT EXISTS commit_endpoints (
    commit_id INTEGER REFERENCES commits(id),
    endpoint_id INTEGER REFERENCES endpoints(id),
    PRIMARY KEY (commit_id, endpoint_id)
);

-- Junction table for commits and models
CREATE TABLE IF NOT EXISTS commit_models (
    commit_id INTEGER REFERENCES commits(id),
    model_id INTEGER REFERENCES models(id),
    PRIMARY KEY (commit_id, model_id)
);

-- Junction table for commits and actions
CREATE TABLE IF NOT EXISTS commit_actions (
    commit_id INTEGER REFERENCES commits(id),
    action_id INTEGER REFERENCES actions(id),
    PRIMARY KEY (commit_id, action_id)
);
