## Used packages

#### For Api Creation

```
go get github.com/labstack/echo/v4
```

#### For database

```
go get github.com/jmoiron/sqlx              [https://github.com/jmoiron/sqlx]
go get -u github.com/go-sql-driver/mysql    [https://github.com/go-sql-driver/mysql]
go get -v github.com/rubenv/sql-migrate/... [https://github.com/rubenv/sql-migrate]
```

#### Other

```
go get github.com/google/uuid
go get -u github.com/golang-jwt/jwt/v5
go get github.com/labstack/echo-jwt/v4
go get github.com/sirupsen/logrus   [For logging]
go get github.com/rifflock/lfshook  [For writing log to file]
go get github.com/lestrrat-go/file-rotatelogs
go get github.com/gocolly/colly/v2  [https://github.com/gocolly/colly]
go get github.com/redis/go-redis/v9
```

#### Github Reference

https://github.com/thanhniencung/backend-github-trending

#### What I have learned

1. Build REST API by using Echo framework
2. Migrate database, use repository design pattern to work with db
3. Work with JWT
4. Write logs to file
5. Use a crawler to crawl data (html format) from websites
6. Implement job queue
7. Use middleware, custom validator
8. Use redis for caching

#### Command

1. For redis

```
- Start the docker
- Access redis server: docker exec -it [container_name] sh
- Run redis cli: redis-cli
- Get data by key: get [key]
- Set data by key: set [key] [value]
- Exit the cli: exit
```
