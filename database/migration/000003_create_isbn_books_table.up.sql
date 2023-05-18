CREATE TABLE isbn_books (
    id SERIAL PRIMARY KEY,
    isbn VARCHAR(20) NOT NULL,
    book_title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    editor VARCHAR(255),
    publisher VARCHAR(255),
    partner VARCHAR(255),
    place_of_printing VARCHAR(255),
    submission_date_lc DATE
);
ALTER TABLE isbn_books
    ADD CONSTRAINT isbn_unique UNIQUE (isbn);

