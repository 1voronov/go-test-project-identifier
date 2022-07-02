CREATE TABLE identifiers (
    id serial not null unique,
    string_identifier varchar(255) unique,
    description varchar(255) not null
);