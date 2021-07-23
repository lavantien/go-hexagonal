CREATE KEYSPACE candy_shop_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE candy_shop_db;

CREATE TABLE IF NOT EXISTS candies (
    candy_id uuid,
    category text,
    name text,
    price float,
    PRIMARY KEY (candy_id, category)
);

INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Chocolate', 'KitKat', 1.99);
INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Sour', 'Skittles', 1.89);

SELECT * FROM candies;