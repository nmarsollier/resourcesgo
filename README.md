# SemVer Resource Server

A microservice for i18n json resources, providing a [semver](https://devhints.io/semver) strategy to
store and retrieve them.

The idea is to store project json resource files by language, and provide a version for them,
so clients could retrieve documents for the desired project and language plus using semver semantic
to provide version compatibility (useful on mobile where app version is hard to ensure).

This gives the client apps a retro compatibility with the texts that are using in different versions,
allowing developers change those strings in different app versions with the flexibility to change the
text ids, remove or arr new text strings over time.

## Architecture

Fully functional, no unwanted interfaces when there is no strategy pattern.

## PostgreSQL

```bash
docker run -d --name ec-postgres -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust postgres:13
```

To generate the database structure run the script DDL.sql.

<img src='https://g.gravizo.com/svg?
digraph%20DER%20%7B%0A%20%20%20%20rankdir%3DLR%3B%0A%20%20%20%20node%20%5Bshape%3Drecord%5D%3B%0A%0A%20%20%20%20projects%20%5Blabel%3D%22%7B%20projects%20%7C%20%7B%20id%20(PK)%20%7C%20name%20%7C%20created%20%7C%20enabled%20%7D%7D%22%5D%3B%0A%20%20%20%20languages%20%5Blabel%3D%22%7B%20languages%20%7C%20%7B%20id%20(PK)%20%7C%20name%20%7C%20created%20%7C%20enabled%20%7D%7D%22%5D%3B%0A%20%20%20%20resources%20%5Blabel%3D%22%7B%20resources%20%7C%20%7B%20id%20(PK)%20%7C%20project%20(FK)%20%7C%20language%20(FK)%20%7C%20sem_ver%20%7C%20values%20%7C%20created%20%7C%20enabled%20%7D%7D%22%5D%3B%0A%0A%20%20%20%20projects%20-%3E%20resources%20%5Blabel%3D%22fk_project%22%5D%3B%0A%20%20%20%20languages%20-%3E%20resources%20%5Blabel%3D%22fk_language%22%5D%3B%0A%20%20%20%20resources%20-%3E%20resources%20%5Blabel%3D%22unique_project_language_semver%22%2C%20style%3Ddotted%5D%3B%0A%7D'>

To setup the database just set the environment variable POSTGRES_URL.

## Go 1.22+

Go [golang.org](https://golang.org/doc/install)

```bash
git clone https://github.com/nmarsollier/resourcesgo $GOPATH/src/github.com/nmarsollier/resourcesgo
```

## Run

```bash
go install
resourcesgo
```

To see swagger docs navigate to [Swagger](http://localhost:3000/docs/index.html)

## Required libraries

```bash
go install github.com/swaggo/gin-swagger/swaggerFiles
go install github.com/swaggo/gin-swagger
go install github.com/swaggo/swag/cmd/swag
go install github.com/99designs/gqlgen@v0.17.56
```

## Environment vars

SERVER_NAME : Server Name for logs (resourcesgo)
POSTGRES_URL : Postgresql database (postgres://postgres@localhost:5432/postgres)
PORT : Server port (3000)
GQL_PORT : GraphQL Port (4000)
