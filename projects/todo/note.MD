## Create a project

Create a folder as project root

## Create the module

Cd to the root folder
run

```
go mod init [module-name]
```

## Create the main file

```
main.go
```

## Used packages

1. for http server: gin [https://github.com/gin-gonic/gin]

```
go get -u github.com/gin-gonic/gin
```

2.for orm: GORM [https://gorm.io/]

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## Use env variables in cscode

https://200lab.io/blog/golang-environment-variable-bien-moi-truong-trong-golang/
