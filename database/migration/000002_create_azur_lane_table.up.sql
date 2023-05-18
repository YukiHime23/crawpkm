CREATE TABLE azur_lanes (
    id SERIAL PRIMARY KEY,
    file_name VARCHAR(255) NOT NULL,
    id_wallpaper INT NOT NULL,
    url VARCHAR(255) NOT NULL
);

CREATE INDEX idx_azur_lanes_id_wallpaper
    ON azur_lanes(id_wallpaper);