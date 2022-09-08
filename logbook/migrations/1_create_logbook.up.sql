CREATE TABLE logbook
(
    id         BIGSERIAL PRIMARY KEY,
    date       DATE      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    text       TEXT      NOT NULL
);

CREATE INDEX ON logbook (date);