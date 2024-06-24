CREATE TABLE schedules (
                        id varchar(36) NOT NULL PRIMARY KEY,
                        start_time timestamp NOT NULL,
                        executed boolean,
                        status text
);