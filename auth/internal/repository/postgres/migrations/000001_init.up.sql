CREATE TYPE role_type AS ENUM ('owner', 'manager', 'user');
CREATE TYPE gender_type AS ENUM ('male', 'female');

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       surname VARCHAR(100) NOT NULL,
                       img VARCHAR DEFAULT NULL,
                       position VARCHAR(255) DEFAULT NULL,
                       location VARCHAR(255) DEFAULT NULL,
                       birth_date TIMESTAMP DEFAULT NULL,
                       gender gender_type NOT NULL,
                       role role_type DEFAULT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       phone VARCHAR(20) UNIQUE DEFAULT NULL,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE company (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(50) NOT NULL,
                         owner_id INTEGER NOT NULL,
                         description VARCHAR DEFAULT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (owner_id) REFERENCES users(id)
);

CREATE TABLE company_employees (
                         company_id INTEGER,
                         user_id INTEGER,
                         PRIMARY KEY (company_id, user_id),
                         FOREIGN KEY (user_id) REFERENCES users (id),
                         FOREIGN KEY (company_id) REFERENCES company (id)
);
