# Restserver
[![Go Report Card](https://goreportcard.com/badge/github.com/philhanna/restserver)][idGoReportCard]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/philhanna/restserver)][idPkgGoDev]

## Overview
This is a sample REST server application inspired by the tutorial
[here](https://tutorialedge.net/golang/creating-restful-api-with-golang).
I have added an SQLite3 database as a backing store.

## Installation
- Clone the git repository:
```bash
git clone git@github.com:philhanna/restserver.git
```
- Create a `config.yml` file in `$HOME/.config/restserver` (on Windows, `%appdata%\restserver`)
```yaml
host: localhost
port: 10000
dbname: /tmp/articles.db
dbsql: |
dbsql: |
  DROP TABLE IF EXISTS articles;
  CREATE TABLE articles (
      id          INTEGER PRIMARY KEY,
      title       TEXT,
      description TEXT,
      content     TEXT
  );
  BEGIN;
  INSERT INTO articles VALUES(null, "Hello 1", "Article 1 description", "Article 1 content");
  INSERT INTO articles VALUES(null, "Hello 2", "Article 2 description", "Article 2 content");
  COMMIT;
```
Adjust the `port` and `dbname` values as needed.

## References
- [Github repository](https://github.com/philhanna/restserver)

[idGoReportCard]: https://goreportcard.com/report/github.com/philhanna/restserver
[idPkgGoDev]: https://pkg.go.dev/github.com/philhanna/restserver
