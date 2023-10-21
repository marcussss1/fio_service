CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE people (
    id          UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR(50),
    surname     VARCHAR(50),
    patronymic  VARCHAR(50),
    age         INTEGER,
    gender      VARCHAR(6),
    nationality VARCHAR(3)
);