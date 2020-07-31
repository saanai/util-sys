CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    uuid       VARCHAR(255) NOT NULL,
    name       VARCHAR(255),
    email      VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL
)
    ENGINE = INNODB
    DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS sessions
(
    id         SERIAL PRIMARY KEY,
    uuid       VARCHAR(64) NOT NULL UNIQUE,
    email      VARCHAR(255),
    user_id    BIGINT UNSIGNED,
    created_at TIMESTAMP   NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)
    ENGINE = INNODB
    DEFAULT CHARSET = utf8;


CREATE TABLE IF NOT EXISTS threads
(
    id         SERIAL PRIMARY KEY,
    uuid       VARCHAR(64) NOT NULL UNIQUE,
    topic      TEXT,
    user_id    BIGINT UNSIGNED,
    created_at TIMESTAMP   NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)
    ENGINE = INNODB
    DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS posts
(
    id         SERIAL PRIMARY KEY,
    uuid       VARCHAR(64) NOT NULL UNIQUE,
    body       TEXT,
    user_id    BIGINT UNSIGNED,
    thread_id  BIGINT UNSIGNED,
    created_at TIMESTAMP   NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (thread_id) REFERENCES threads (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)
    ENGINE = INNODB
    DEFAULT CHARSET = utf8;