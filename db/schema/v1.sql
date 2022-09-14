CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar,
    phone varchar,
    verified_phone BOOLEAN,
    email varchar,
    verified_email BOOLEAN,
    password varchar,
    token_id varchar,
    country varchar,
    status varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    user_id int,
    name varchar,
    region varchar,
    city varchar,
    address varchar,
    longitude varchar,
    latitude varchar,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    name varchar,
    email varchar,
    token_id varchar,
    password varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE merchants (
    id SERIAL PRIMARY KEY,
    email varchar,
    status varchar,
    token_id varchar,
    password varchar,
    name varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE stores (
    id SERIAL PRIMARY KEY,
    merchant_id int,
    name varchar,
    description varchar,
    phone varchar,
    location varchar,
    country varchar,
    access_key varchar,
    cash_on_delivery BOOLEAN,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    store_id int,
    name varchar,
    url varchar,
    logo varchar,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE subcategories (
    id SERIAL PRIMARY KEY,
    name varchar,
    category_id int,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE attributes (
    id SERIAL PRIMARY KEY,
    name varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE attribute_values (
    id SERIAL PRIMARY KEY,
    name varchar,
    attribute_id int,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE currencies (
    id SERIAL PRIMARY KEY,
    name varchar,
    symbol varchar,
    factor float
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id int,
    currency_id int,
    amount float,
    created_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    sku varchar,
    name varchar,
    description varchar,
    long_description varchar,
    price float,
    store_id int,
    category_id int,
    subcategory_id int,
    stock int,
    status varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE selected_attributes (
    id SERIAL PRIMARY KEY,
    item_id int,
    attribute_id int,
    attribute_value_id int
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    status varchar,
    total float,
    total_discounted float,
    user_id int,
    address_id int,
    currency_id int,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE favorites (
    id SERIAL PRIMARY KEY,
    user_id int,
    item_id int
);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id int,
    order_id int,
    rating int,
    content varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id int,
    item_id int
);

CREATE TABLE items_order (
    id SERIAL PRIMARY KEY,
    order_id int,
    item_id int,
    store_id int,
    price float,
    discounted_price float,
    payment varchar,
    status varchar
);

CREATE TABLE item_images (
    id SERIAL PRIMARY KEY,
    source varchar,
    item_id int
);