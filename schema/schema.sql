--- Database Schema
CREATE TABLE IF NOT EXISTS rules (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   description TEXT,
   rule TEXT NOT NULL,
   ast JSONB NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

--- Sample Data
/*INSERT INTO rules (
    name, description, rule, ast
) VALUES (
    "", "", "", ""
)*/;
