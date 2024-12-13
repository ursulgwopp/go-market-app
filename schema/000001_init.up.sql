CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username TEXT,
	password TEXT,
	email TEXT
);
	
CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	name TEXT,
	description TEXT,
	price INTEGER,
	quantity INTEGER
);
	
CREATE TABLE IF NOT EXISTS purchases (
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	product_id INTEGER,
	quantity INTEGER,
	timestamp TEXT
);