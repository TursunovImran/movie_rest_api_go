CREATE TABLE films (
    id              SERIAL          PRIMARY KEY,
    title           VARCHAR(150)    NOT NULL,
    description     VARCHAR(1000)   NOT NULL,
    release_date    DATE            NOT NULL,
    rating          DECIMAL(3, 1)   CHECK (rating >= 0 AND rating <= 10)
);

CREATE TABLE actors (
    id              SERIAL          PRIMARY KEY,
    name            VARCHAR(255)    NOT NULL,
    gender          VARCHAR(10)     NOT NULL,
    birthdate    DATE            NOT NULL
);

CREATE TABLE films_actors (
    id              SERIAL PRIMARY KEY,
    film_id         INT,
    actor_id        INT,
    FOREIGN KEY     (film_id) REFERENCES films(id),
    FOREIGN KEY     (actor_id) REFERENCES actors(id)
);