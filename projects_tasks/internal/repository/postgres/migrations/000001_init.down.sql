-- migration_down.sql

-- Drop the "tasks" table
DROP TABLE tasks;

-- Drop the "project_participants" table
DROP TABLE project_participants;

-- Drop the "projects" table
DROP TABLE projects;

-- Drop the "priority" ENUM type
DROP TYPE priority;

-- Drop the "status" ENUM type
DROP TYPE status;
