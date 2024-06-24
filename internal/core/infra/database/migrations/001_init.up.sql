CREATE TABLE users (
                        id varchar(36) NOT NULL PRIMARY KEY,
                        email varchar(320) NOT NULL UNIQUE,
                        phone varchar(100) NOT NULL UNIQUE,
                        location varchar(200),
                        active boolean,
                        created timestamp
);