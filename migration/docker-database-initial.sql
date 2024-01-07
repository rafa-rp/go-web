create table products(
    id serial primary key,
    name varchar,
    description varchar,
    price decimal,
    quantity integer
);

INSERT INTO products(name, description, price, quantity) VALUES
('T-Shirt','Black',19,10),
('Phone','Head',99, 5);