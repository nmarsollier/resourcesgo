# SemVer Resource Server

A microservice for i18n json resources, providing a [semver](https://devhints.io/semver) strategy to 
store and retrieve them.

The idea is to store project json resource files by language, and provide a version for them,
so clients could retrieve documents for the desired project and language plus using semver semantic
to provide version compatibility (useful on mobile where app version is hard to ensure).

We need a mongoose database :

## MongoDB

```bash
docker run -d --name ec-mongo -p 27017:27017 mongo:4.0.18-xenial
```

## Go

Go 1.14  [golang.org](https://golang.org/doc/install)

```bash
export GO111MODULE=on
export GOFLAGS=-mod=vendor
```


```bash
git clone https://github.com/nmarsollier/resourcesgo $GOPATH/src/github.com/nmarsollier/resourcesgo
```

## Run

```bash
go mod download
go mod vendor
go install
resourcesgo
```

