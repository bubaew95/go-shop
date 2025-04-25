-- +goose Up
CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    firm_id INTEGER DEFAULT null,
    user_id INTEGER DEFAULT null,
    name VARCHAR(255) NOT NULL,
    anons VARCHAR(512) DEFAULT NULL,
    text TEXT DEFAULT NULL,
    stock INTEGER DEFAULT NULL,
    price DOUBLE PRECISION DEFAULT 0,
    discount SMALLINT DEFAULT NULL,

    seo_title VARCHAR(255) DEFAULT null,
    seo_description VARCHAR(512) DEFAULT null,
    seo_keywords VARCHAR(255) DEFAULT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS product;

