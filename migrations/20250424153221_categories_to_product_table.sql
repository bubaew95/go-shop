-- +goose Up
CREATE TABLE categories_to_product (
    product_id INTEGER DEFAULT NULL,
    category_id INTEGER DEFAULT NULL
);
ALTER TABLE categories_to_product ADD CONSTRAINT FK_PRODUCT_ID FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE SET NULL;
ALTER TABLE categories_to_product ADD CONSTRAINT FK_CATEGORY_ID FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE categories_to_product DROP CONSTRAINT FK_PRODUCT_ID;
ALTER TABLE categories_to_product DROP CONSTRAINT FK_CATEGORY_ID;
DROP TABLE categories_to_product;

