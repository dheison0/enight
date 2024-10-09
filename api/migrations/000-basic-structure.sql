CREATE TABLE IF NOT EXISTS locations(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  distance INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS clients(
  phone TEXT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  location_id INTEGER,
  FOREIGN KEY(location_id) REFERENCES locations(id)
);

CREATE TABLE IF NOT EXISTS products(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT,
  cover_url TEXT
);

CREATE TABLE IF NOT EXISTS product_sizes(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  price NUMBER NOT NULL,
  product_id INTEGER NOT NULL,
  FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
);


/* this table uses raw clients and products to avoid losing data when a client
   or product is deleted */
CREATE TABLE IF NOT EXISTS purchases(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  client TEXT NOT NULL,
  price NUMERIC NOT NULL,
  products TEXT NOT NULL,
  stage TEXT NOT NULL DEFAULT "added"
);