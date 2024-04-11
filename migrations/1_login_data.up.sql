CREATE TABLE IF NOT EXISTS login_data (
    id UUID DEFAULT uuid_generate_v4(),
    user_name VARCHAR(55) NOT NULL,
    password VARCHAR(100) NOT NULL,
    old_password VARCHAR(100) NULL,
    wrong_login_attempt INT NULL,
    today_login_attempt INT NULL,
    is_now_login BOOLEAN DEFAULT FALSE,
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    time TIME NOT NULL DEFAULT CURRENT_TIME,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
