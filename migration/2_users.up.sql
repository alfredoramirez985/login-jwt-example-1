CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL,
    name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    country VARCHAR NOT NULL,
    zip_code VARCHAR NOT NULL,
    id_login_data UUID NOT NULL,
    PRIMARY KEY (id)
    CONSTRAINT fk_login_data
      FOREIGN KEY(id_login_data) 
        REFERENCES logindata(id)
);