CREATE TABLE products (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    mark_id CHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_products_mark_id 
        FOREIGN KEY (mark_id) 
        REFERENCES marks(id) 
        ON DELETE RESTRICT 
        ON UPDATE CASCADE
);

CREATE INDEX idx_products_mark_id ON products(mark_id);
CREATE INDEX idx_products_name ON products(name);