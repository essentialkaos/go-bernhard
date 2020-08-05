<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-bernhard.svg"/></a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/bernhard.v2"><img src="https://godoc.org/pkg.re/essentialkaos/bernhard.v2?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/bernhard"><img src="https://goreportcard.com/badge/github.com/essentialkaos/bernhard"></a>
  <a href="https://travis-ci.com/essentialkaos/bernhard"><img src="https://travis-ci.com/essentialkaos/bernhard.svg"></a>
  <a href="https://github.com/essentialkaos/bernhard/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/bernhard/workflows/CodeQL/badge.svg" /></a>
  <a href='https://coveralls.io/github/essentialkaos/bernhard?branch=master'><img src='https://coveralls.io/repos/github/essentialkaos/bernhard/badge.svg?branch=master' alt='Coverage Status' /></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-bernhard-master"><img alt="codebeat badge" src="https://codebeat.co/badges/958c1200-21d8-4e14-964e-fdc88000520c" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`bernhard` is Go package for sending alerts to Bernhard service.

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.12+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/bernhard.v3
```

For update to the latest stable release, do:

```
go get -u pkg.re/essentialkaos/bernhard.v3
```

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![Build Status](https://travis-ci.com/essentialkaos/bernhard.svg?branch=master)](https://travis-ci.com/essentialkaos/bernhard) |
| `develop` | [![Build Status](https://travis-ci.com/essentialkaos/bernhard.svg?branch=develop)](https://travis-ci.com/essentialkaos/bernhard) |

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
