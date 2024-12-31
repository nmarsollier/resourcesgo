# SemVer Resource Server

A microservice for i18n json resources, providing a [semver](https://devhints.io/semver) strategy to
store and retrieve them.

The idea is to store project json resource files by language, and provide a version for them,
so clients could retrieve documents for the desired project and language plus using semver semantic
to provide version compatibility (useful on mobile where app version is hard to ensure).

## PostgreSQL

```bash
docker run -d --name ec-postgres -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust postgres:13
```

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
