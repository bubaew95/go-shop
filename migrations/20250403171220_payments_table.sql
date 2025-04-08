-- +goose Up
CREATE TABLE payments (
    id SERIAL PRIMARY KEY ,
    name VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose Down
DROP TABLE payments;
