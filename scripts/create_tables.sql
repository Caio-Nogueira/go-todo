DROP TABLE IF EXISTS "todo";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Todo (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        title VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        status INTEGER NOT NULL DEFAULT 0
    );
