CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username TEXT UNIQUE,
	hash_password TEXT,
	email TEXT,
	balance INTEGER DEFAULT 0,
	product_list INTEGER[] DEFAULT '{}'
);
	
CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	name TEXT,
	description TEXT,
	price INTEGER,
	quantity INTEGER,
	owner_id INTEGER
);
	
CREATE TABLE IF NOT EXISTS purchases (
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	product_id INTEGER,
	cost INTEGER,
	quantity INTEGER,
	timestamp TEXT
);