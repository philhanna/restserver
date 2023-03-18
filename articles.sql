DROP TABLE IF EXISTS articles;
CREATE TABLE articles (
    id text,
    title text,
    description text,
    content text
);
BEGIN;
INSERT INTO articles VALUES("1", "Hello 1", "Article 1 description", "Article 1 content");
INSERT INTO articles VALUES("2", "Hello 2", "Article 2 description", "Article 2 content");
COMMIT;
