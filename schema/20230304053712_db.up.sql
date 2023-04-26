CREATE TABLE users
(
    id serial not null unique,
    login varchar(255) not null unique,
    first_name varchar(255),
    last_name varchar(255),
    password varchar(255) not null,
    status varchar(255) not null
);

CREATE TABLE artists
(
    id serial not null unique,
    name_artist varchar(255) not null unique,
    date_of_birth varchar(255) not null,
    about_artist varchar(1000)
);

CREATE TABLE songs
(
    id serial not null unique,
    artist_id int references artist (id) on delete cascade,
    name_song varchar(255) not null,
    genre varchar(255) not null,
    second_genre varchar(255),
    year_of_release int not null,
    rating  int
);

CREATE TABLE user_library
(
    id serial not null unique,
    user_id int references users (id) on delete cascade,
    song_id int references song (id) on delete cascade,
);

CREATE TABLE user_top_ten
(
    id serial not null unique,
    user_id int references users (id) on delete cascade,
    song_id int references song (id) on delete cascade,
);