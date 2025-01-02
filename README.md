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
digraph DER {
    rankdir=LR;
    node [shape=record];
    projects [label="{ projects | { id %28PK%29 | name | created | enabled }}"];
    languages [label="{ languages | { id %28PK%29 | name | created | enabled }}"];
    resources [label="{ resources | { id %28PK%29 | project %28FK%29 | language %28FK%29 | sem_ver | values | created | enabled }}"];
    projects -> resources [label="fk_project"];
    languages -> resources [label="fk_language"];
    resources -> resources [label="unique_project_language_semver", style=dotted];
}
'>


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
