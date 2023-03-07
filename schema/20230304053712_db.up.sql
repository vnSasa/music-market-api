CREATE TABLE users
(
    id serial not null unique,
    login varchar(255) not null unique,
    first_name varchar(255),
    last_name varchar(255),
    password varchar(255) not null
);

CREATE TABLE artist
(
    id serial not null unique,
    name_artist varchar(255) not null,
    date_of_birth varchar(255),
    about_artist varchar(255)
);

CREATE TABLE song
(
    id serial not null unique,
    artist_id int references artist (id) on delete cascade,
    name_song varchar(255) not null,
    genre varchar(255) not null,
    second_genre varchar(255),
    year_of_release varchar(255) not null
);

CREATE TABLE user_library
(
    id serial not null unique,
    user_id int references users (id) on delete cascade,
    song_id int references song (id) on delete cascade
);