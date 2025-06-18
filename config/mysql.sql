CREATE TABLE IF NOT EXISTS product
(
    id             BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name           VARCHAR(255)   NOT NULL,
    price          DECIMAL(10, 2) NOT NULL,
    original_price DECIMAL(10, 2) NOT NULL,
    image          VARCHAR(512),
    brand          VARCHAR(100),
    brand_subtitle VARCHAR(255),
    created_at     DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at     DATETIME,
    INDEX idx_deleted_at (deleted_at)
);

INSERT INTO product (id, name, price, original_price, image, brand, brand_subtitle)
VALUES (1, '澳洲谷饲安格斯M3+雪花原切西冷牛排', 59.90, 8452.00,
        'https://images.unsplash.com/photo-1546833999-b9f581a1996d?w=400&h=300&fit=crop', 'HONG HUI', 'butcher life'),
       (2, '澳洲谷饲安格斯M3+雪花牛仔骨原切', 59.90, 6854.00,
        'https://images.unsplash.com/photo-1558030006-450675393462?w=400&h=300&fit=crop', 'HONG HUI', 'butcher life'),
       (3, '澳洲原切150天小西冷牛排（煎炒两用）', 19.90, 10360.00,
        'https://images.unsplash.com/photo-1529692236671-f1f6cf9683ba?w=400&h=300&fit=crop', 'HONG HUI',
        'butcher life'),
       (4, '美国谷饲Prime级佳级原切西冷牛排', 59.90, 6651.00,
        'https://images.unsplash.com/photo-1607623814075-e51df1bdc82f?w=400&h=300&fit=crop', 'HONG HUI',
        'butcher life');
