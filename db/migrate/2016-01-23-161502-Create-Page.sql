DROP TABLE IF EXISTS pages;
CREATE TABLE pages (
id SERIAL NOT NULL,
created_at timestamp,
updated_at timestamp,
status integer,
name text,
url text,
keywords text, 
summary text,
text text
);
ALTER TABLE pages OWNER TO sendto_server;
