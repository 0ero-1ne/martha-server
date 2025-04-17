DROP DATABASE IF EXISTS test;

CREATE USER test
WITH
    ENCRYPTED PASSWORD 'test';

CREATE DATABASE martha
WITH
    OWNER martha_admin ENCODING = 'UTF8' CONNECTION
LIMIT
    = -1 TABLESPACE = pg_default LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8' LOCALE_PROVIDER = 'libc' IS_TEMPLATE = FALSE;

DROP TABLE IF EXISTS comment_rates;

DROP TABLE IF EXISTS book_rates;

DROP TABLE IF EXISTS book_tag;

DROP TABLE IF EXISTS book_author;

DROP TABLE IF EXISTS chapters;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS books;

DROP TABLE IF EXISTS tags;

DROP TABLE IF EXISTS authors;

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    email VARCHAR(320) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    username VARCHAR(64) NOT NULL UNIQUE,
    ROLE VARCHAR(10) NOT NULL DEFAULT 'user',
    image VARCHAR(256) DEFAULT NULL,
    saved_books JSONB DEFAULT '[]'
) TABLESPACE pg_default;

CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title VARCHAR(256) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT '?',
    "year" SMALLINT NOT NULL DEFAULT 0,
    "views" BIGINT NOT NULL DEFAULT 0,
    "cover" VARCHAR(256) DEFAULT NULL
) TABLESPACE pg_default;

CREATE TABLE authors (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    fullname VARCHAR(256) NOT NULL DEFAULT 'unknown author',
    biography TEXT NOT NULL DEFAULT 'unknown',
    image VARCHAR(256) DEFAULT NULL
) TABLESPACE pg_default;

CREATE TABLE books_authors (
    book_id BIGINT NOT NULL,
    author_id BIGINT NOT NULL,
    CONSTRAINT books_authors_pkey PRIMARY KEY (book_id, author_id),
    CONSTRAINT fk_books_authors_book FOREIGN KEY (book_id) REFERENCES books (id) MATCH SIMPLE ON DELETE CASCADE,
    CONSTRAINT fk_books_authors_author FOREIGN KEY (author_id) REFERENCES authors (id) MATCH SIMPLE ON DELETE CASCADE
) TABLESPACE pg_default;

CREATE TABLE tags (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title VARCHAR(128) NOT NULL UNIQUE
) TABLESPACE pg_default;

CREATE TABLE books_tags (
    book_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    CONSTRAINT books_tags_pkey PRIMARY KEY (book_id, tag_id),
    CONSTRAINT fk_books_tags_book FOREIGN KEY (book_id) REFERENCES books (id) MATCH SIMPLE ON DELETE CASCADE,
    CONSTRAINT fk_books_tags_tag FOREIGN KEY (tag_id) REFERENCES tags (id) MATCH SIMPLE ON DELETE CASCADE
) TABLESPACE pg_default;

CREATE TABLE chapters (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title VARCHAR(256) NOT NULL,
    number SMALLINT NOT NULL,
    "text" VARCHAR(256) DEFAULT NULL,
    "audio" VARCHAR(256) DEFAULT NULL,
    book_id BIGINT NOT NULL,
    CONSTRAINT fk_chapters_book FOREIGN KEY (book_id) REFERENCES books (id) MATCH SIMPLE ON DELETE CASCADE
) TABLESPACE pg_default;

CREATE TABLE books_rates (
    book_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    rating SMALLINT NOT NULL CHECK (rating BETWEEN 0 AND 5),
    CONSTRAINT books_rates_pkey PRIMARY KEY (book_id, user_id),
    CONSTRAINT fk_books_rates_book FOREIGN KEY (book_id) REFERENCES books (id) MATCH SIMPLE ON DELETE CASCADE,
    CONSTRAINT fk_books_rates_user FOREIGN KEY (user_id) REFERENCES users (id) MATCH SIMPLE ON DELETE CASCADE
) TABLESPACE pg_default;

CREATE TABLE comments (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    "text" TEXT NOT NULL,
    book_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    CONSTRAINT fk_comments_book FOREIGN KEY (book_id) REFERENCES books (id) MATCH SIMPLE ON DELETE CASCADE,
    CONSTRAINT fk_comments_user FOREIGN KEY (user_id) REFERENCES users (id) MATCH SIMPLE ON DELETE CASCADE
) TABLESPACE pg_default;

CREATE TABLE comments_rates (
    comment_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    rating SMALLINT NOT NULL,
    CONSTRAINT comments_rates_pkey PRIMARY KEY (comment_id, user_id),
    CONSTRAINT fk_comments_rates_comment FOREIGN KEY (comment_id) REFERENCES comments (id) MATCH SIMPLE ON DELETE CASCADE,
    CONSTRAINT fk_comments_rates_user FOREIGN KEY (user_id) REFERENCES users (id) MATCH SIMPLE ON DELETE CASCADE
) TABLESPACE pg_default;

CREATE OR REPLACE FUNCTION update_modified_column () RETURNS TRIGGER AS $$
BEGIN
NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE OR REPLACE TRIGGER update_modified_time BEFORE
UPDATE ON tags FOR EACH ROW
EXECUTE PROCEDURE update_modified_column ();

CREATE OR REPLACE TRIGGER update_modified_time BEFORE
UPDATE ON books FOR EACH ROW
EXECUTE PROCEDURE update_modified_column ();

CREATE OR REPLACE TRIGGER update_modified_time BEFORE
UPDATE ON users FOR EACH ROW
EXECUTE PROCEDURE update_modified_column ();

CREATE OR REPLACE TRIGGER update_modified_time BEFORE
UPDATE ON chapters FOR EACH ROW
EXECUTE PROCEDURE update_modified_column ();

CREATE OR REPLACE TRIGGER update_modified_time BEFORE
UPDATE ON authors FOR EACH ROW
EXECUTE PROCEDURE update_modified_column ();

CREATE OR REPLACE TRIGGER update_modified_time BEFORE
UPDATE ON comments FOR EACH ROW
EXECUTE PROCEDURE update_modified_column ();