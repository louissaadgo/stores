CREATE TABLE users (
    id varchar PRIMARY KEY,
    phone varchar,
    password varchar,
    sign_type varchar,
    sign_id varchar,
    bday  TIMESTAMP WITH TIME ZONE,
    image varchar,
    country varchar,
    location varchar,
    loyality_points int,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE admins (
    id varchar PRIMARY KEY,
    name varchar,
    email varchar,
    password varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE stores (
    id varchar PRIMARY KEY,
    email varchar,
    name varchar,
    password varchar,
    description varchar,
    phone varchar,
    location varchar,
    country varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
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

CREATE TABLE wallets (
    id varchar PRIMARY KEY,
    user_id varchar,
    amount float,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE transactions (
    id varchar PRIMARY KEY,
    wallet_id varchar,
    amount float,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);

CREATE TABLE currencies (
    id varchar PRIMARY KEY,
    name varchar,
    symbol varchar,
    factor float
);

CREATE TABLE items (
    id varchar PRIMARY KEY,
    sku varchar,
    name varchar,
    description varchar,
    long_description varchar,
    price float,
    images varchar ARRAY,
    store_id varchar,
    attribute_ids varchar ARRAY,
    category_id varchar,
    subcategory_id varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (store_id) REFERENCES stores(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (subcategory_id) REFERENCES subcategories(id)
);

CREATE TABLE orders (
    id varchar PRIMARY KEY,
    status varchar,
    item_ids varchar ARRAY,
    total float,
    coupon_id varchar,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (coupon_id) REFERENCES coupons(id)
);

CREATE TABLE coupons (
    id varchar PRIMARY KEY,
    value float,
    max_usage int,
    used int,
    code varchar,
    end_date TIMESTAMP WITH TIME ZONE,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE favorites (
    id varchar PRIMARY KEY,
    user_id varchar,
    item_ids varchar ARRAY,
    updated_at  TIMESTAMP WITH TIME ZONE,
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
    item_ids varchar ARRAY,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE interests (
    id varchar PRIMARY KEY,
    user_id varchar,
    category_ids varchar ARRAY,
    updated_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);