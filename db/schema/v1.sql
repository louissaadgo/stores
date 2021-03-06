CREATE TABLE users (
    id varchar PRIMARY KEY,
    phone varchar,
    name varchar,
    token_id varchar,
    password varchar,
    sign_type varchar,
    sign_id varchar,
    status varchar,
    bday  TIMESTAMP WITH TIME ZONE,
    image varchar,
    country varchar,
    loyality_points int,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE addresses (
    id varchar PRIMARY KEY,
    user_id varchar,
    name varchar,
    region varchar,
    city varchar,
    address varchar,
    longitude varchar,
    latitude varchar,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE admins (
    id varchar PRIMARY KEY,
    name varchar,
    email varchar,
    token_id varchar,
    password varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE merchants (
    id varchar PRIMARY KEY,
    email varchar,
    status varchar,
    token_id varchar,
    password varchar,
    name varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE coupons (
    id varchar PRIMARY KEY,
    value float,
    type varchar,
    max_usage int,
    used int,
    code varchar,
    end_date TIMESTAMP WITH TIME ZONE,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE stores (
    id varchar PRIMARY KEY,
    merchant_id varchar,
    name varchar,
    description varchar,
    phone varchar,
    location varchar,
    country varchar,
    access_key varchar,
    cash_on_delivery BOOLEAN,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);

CREATE TABLE links (
    id varchar PRIMARY KEY,
    store_id varchar,
    name varchar,
    url varchar,
    logo varchar,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (store_id) REFERENCES stores(id)
);

CREATE TABLE categories (
    id varchar PRIMARY KEY,
    name varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE subcategories (
    id varchar PRIMARY KEY,
    name varchar,
    category_id varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE attributes (
    id varchar PRIMARY KEY,
    name varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE attribute_values (
    id varchar PRIMARY KEY,
    name varchar,
    attribute_id varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (attribute_id) REFERENCES attributes(id)
);

CREATE TABLE currencies (
    id varchar PRIMARY KEY,
    name varchar,
    symbol varchar,
    factor float
);

CREATE TABLE transactions (
    id varchar PRIMARY KEY,
    user_id varchar,
    currency_id varchar,
    amount float,
    created_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (currency_id) REFERENCES currencies(id)
);

CREATE TABLE items (
    id varchar PRIMARY KEY,
    sku varchar,
    name varchar,
    description varchar,
    long_description varchar,
    price float,
    store_id varchar,
    category_id varchar,
    subcategory_id varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    stock int,
    status varchar,
    FOREIGN KEY (store_id) REFERENCES stores(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (subcategory_id) REFERENCES subcategories(id)
);

CREATE TABLE item_images (
    id varchar PRIMARY KEY,
    item_id varchar,
    url varchar,
    FOREIGN KEY (item_id) REFERENCES items(id)
);

CREATE TABLE selected_attributes (
    id varchar PRIMARY KEY,
    item_id varchar,
    attribute_id varchar,
    attribute_value_id varchar,
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (attribute_id) REFERENCES attributes(id),
    FOREIGN KEY (attribute_value_id) REFERENCES attribute_values(id)
);

CREATE TABLE orders (
    id varchar PRIMARY KEY,
    status varchar,
    total float,
    total_discounted float,
    user_id varchar,
    coupon_id varchar,
    address_id varchar,
    currency_id varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE favorites (
    id varchar PRIMARY KEY,
    user_id varchar,
    item_id varchar,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE reviews (
    id varchar PRIMARY KEY,
    user_id varchar,
    item_id varchar,
    order_id varchar,
    rating int,
    content varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE carts (
    id varchar PRIMARY KEY,
    user_id varchar,
    item_id varchar,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE interests (
    id varchar PRIMARY KEY,
    user_id varchar,
    category_id varchar,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE items_order (
    id varchar PRIMARY KEY,
    order_id varchar,
    item_id varchar,
    store_id varchar,
    price float,
    discounted_price float,
    payment varchar,
    status varchar
);

CREATE TABLE item_images (
    id varchar PRIMARY KEY,
    source varchar,
    item_id varchar
);