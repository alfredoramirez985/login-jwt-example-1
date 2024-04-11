CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL,
    first_name VARCHAR(55) NOT NULL,
    last_name VARCHAR(55) NOT NULL,
    phone VARCHAR(15) NOT NULL,
    email VARCHAR(55) NOT NULL,
    id_login_data UUID NOT NULL,
    CONSTRAINT fk_login_data FOREIGN KEY(id_login_data) REFERENCES login_data(id),
    PRIMARY KEY (id)
);