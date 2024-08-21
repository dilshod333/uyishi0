create table books(
    id serial primary key,
    title varchar(64),
    author text,
    published_date text,
    isbn bigint 
);