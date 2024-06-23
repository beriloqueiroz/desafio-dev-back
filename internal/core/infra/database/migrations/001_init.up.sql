CREATE TABLE users (
                        id varchar(36) NOT NULL PRIMARY KEY,
                        email varchar(320) NOT NULL,
                        phone varchar(100) NOT NULL,
                        location varchar(200),
                        active boolean,
                        created timestamp
);