CREATE TYPE status AS ENUM ('to do', 'doing', 'done');
CREATE TYPE priority AS ENUM ('high', 'medium','low');

CREATE TABLE projects (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       priority priority NOT NULL,
                       img VARCHAR DEFAULT NULL,
                       description VARCHAR DEFAULT NULL,
                       company_id INTEGER NOT NULL,
                       reporter_id INTEGER NOT NULL,
                       assignees_ids INTEGER[] NOT NULL,
                       starts_at TIMESTAMP default CURRENT_TIMESTAMP,
                       deadline_at TIMESTAMP DEFAULT NULL,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       project_id INTEGER NOT NULL,
                       name VARCHAR(50) UNIQUE NOT NULL,
                       "group" VARCHAR,
                       owner_id INTEGER NOT NULL,
                       description VARCHAR DEFAULT NULL,
                       priority priority NOT NULL,
                       assign INTEGER NOT NULL,
                       status status NOT NULL,
                       deadline_at TIMESTAMP DEFAULT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (project_id) REFERENCES projects(id)
);


CREATE TABLE project_participants (
                                   project_id INTEGER,
                                   user_id INTEGER,
                                   PRIMARY KEY (project_id, user_id),
                                   FOREIGN KEY (project_id) REFERENCES projects (id)
);
