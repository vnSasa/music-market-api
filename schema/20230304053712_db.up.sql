CREATE TABLE users
(
    id serial not null unique,
    login varchar(255) not null unique,
    first_name varchar(255),
    last_name varchar(255),
    password varchar(255) not null
);

CREATE TABLE products
(
    id serial not null unique,
    name_product varchar(255) not null,
    description varchar(255) not null,
    price varchar(255) not null
);

CREATE TABLE bucket
(
    id serial not null unique,
    user_id int references users (id) on delete cascade,
    product_id int references products (id) on delete cascade
)