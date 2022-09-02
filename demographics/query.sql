-- SQLite

CREATE TABLE
    countries(
        id text NOT NULL PRIMARY KEY,
        name varchar(150) UNIQUE NOT NULL,
        continent varchar(150) NOT NULL,
        created_at datetime DEFAULT (datetime('now')),
        updated_at datetime DEFAULT (datetime('now'))
    );

INSERT INTO
    countries (id, name, continent)
VALUES (
        'd8fbb562-2ae0-11ed-8187-acde48001122 ',
        'Liberia',
        'Africa'
    ),
VALUES (
        '42c441a0-2ad5-11ed-ac47-acde48001122',
        'Sierra Leone',
        'Africa'
    ),
VALUES (
        '038cc4a8-2ad5-11ed-ab57-acde48001122',
        'Nigeria',
        'Africa'
    ),
VALUES (
        '4c47ca94-2ad5-11ed-804a-acde48001122',
        'Ghana',
        'Africa'
    )
VALUES (
        '4c47ca94-2ad5-11ed-804a-acde48001122',
        'Ghana',
        'Africa'
    );

CREATE TABLE
    regions (
        id TEXT NOT NULL PRIMARY KEY,
        name VARCHAR(150) NOT NULL UNIQUE,
        country_id TEXT NOT NULL,
        created_at datetime DEFAULT (datetime('now')),
        updated_at datetime DEFAULT (datetime('now')),
        FOREIGN KEY (country_id) REFERENCES countries (id)
    );

    