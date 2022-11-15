# restful-file-zip

## Changelog

- **v1**: developing

## Description

This is a personal side project to upload file and zip download service.

Using Technical:

- Golang
  - [Go Clean Architecture](https://github.com/bxcodec/go-clean-arch)
    - [Gin-Gonic](https://gin-gonic.com/)
  - [Mockery](https://github.com/vektra/mockery)
  - [Swaggo](https://github.com/swaggo/swag)
  - [Golangci-lint](https://github.com/golangci/golangci-lint)

## System Requirement

- Golang 1.16+

\*Make Sure you have set **GOROOT**/bin, **GOPATH**/bin into your env **PATH\***

## Install

```bash
$ make install
# Includes
  $ go get
  $ make mockery-install
  $ make swaggo-install
  $ make lint-install
```

## Test Golang

```bash
# quick test
$ make test
# test and open cover report
$ make test-cover
```

## Before Run

```bash
$ make before-run
# Includes
  $ make mockery
  $ make swaggo
  $ make lint
  $ make test
```

## Run Application

```bash
$ make run
```
