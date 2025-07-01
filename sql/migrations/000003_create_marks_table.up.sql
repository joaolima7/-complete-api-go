CREATE TABLE marks (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
);

CREATE INDEX idx_marks_name ON marks(name);