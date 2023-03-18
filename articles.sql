DROP TABLE IF EXISTS articles;
CREATE TABLE articles (
    Id text,
    Title text,
    Description text,
    Content text
);
BEGIN;
INSERT INTO articles VALUES("1", "Hello 1", "Article 1 description", "Article 1 content");
INSERT INTO articles VALUES("1", "Hello 2", "Article 2 description", "Article 2 content");
COMMIT;
