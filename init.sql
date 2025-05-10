-- Creating the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT,
    password TEXT NOT NULL
);

-- Creating the wines table
CREATE TABLE wines (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    year INTEGER,
    manufacturer TEXT,
    region TEXT,
    alcohol_content REAL,
    serving_temp REAL,
    serving_size REAL,
    serving_size_unit TEXT,
    serving_size_unit_abbreviation TEXT,
    serving_size_unit_description TEXT,
    serving_size_unit_description_abbreviation TEXT,
    serving_size_unit_description_plural TEXT,
    price REAL,
    rating REAL,
    review_count INTEGER DEFAULT 0,
    average_rating REAL,
    type TEXT,
    colour TEXT
);

-- Creating the reviews table
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    wine_id INTEGER NOT NULL,
    comment TEXT,
    review_date DATE,
    review_date_time TIMESTAMP WITH TIME ZONE,
    review_date_time_utc TIMESTAMP WITH TIME ZONE,
    title TEXT,
    description TEXT,
    rating INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (wine_id) REFERENCES wines(id) ON DELETE CASCADE
);

-- Creating the user_favorite_wines junction table for many-to-many relationship
CREATE TABLE user_favorite_wines (
    user_id INTEGER NOT NULL,
    wine_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, wine_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (wine_id) REFERENCES wines(id) ON DELETE CASCADE
);

-- Optional: Creating a view for UserProfile (if not a separate table)
CREATE VIEW user_profiles AS
SELECT 
    u.id,
    u.username,
    u.email,
    ARRAY_AGG(w.id) AS favorite_wines,
    ARRAY_AGG(r.id) AS reviews
FROM users u
LEFT JOIN user_favorite_wines ufw ON u.id = ufw.user_id
LEFT JOIN wines w ON ufw.wine_id = w.id
LEFT JOIN reviews r ON u.id = r.user_id
GROUP BY u.id, u.username, u.email;

-- Optional: Creating a view for WineProfile (if not a separate table)
CREATE VIEW wine_profiles AS
SELECT 
    w.id,
    w.name,
    w.year,
    w.manufacturer,
    w.type,
    w.colour,
    ARRAY_AGG(r.id) AS reviews,
    w.rating,
    w.price,
    w.region,
    w.alcohol_content,
    w.serving_temp,
    w.serving_size,
    w.serving_size_unit,
    w.serving_size_unit_abbreviation,
    w.serving_size_unit_description,
    w.serving_size_unit_description_abbreviation,
    w.serving_size_unit_description_plural
FROM wines w
LEFT JOIN reviews r ON w.id = r.wine_id
GROUP BY 
    w.id,
    w.name,
    w.year,
    w.manufacturer,
    w.type,
    w.colour,
    w.rating,
    w.price,
    w.region,
    w.alcohol_content,
    w.serving_temp,
    w.serving_size,
    w.serving_size_unit,
    w.serving_size_unit_abbreviation,
    w.serving_size_unit_description,
    w.serving_size_unit_description_abbreviation,
    w.serving_size_unit_description_plural;