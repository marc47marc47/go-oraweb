<<<<<<< HEAD
# go-oraweb

Web-based Oracle database browser written in Go.

[![Release](https://img.shields.io/github/release/marc47marc47/go-oraweb.svg?label=Release)](https://github.com/marc47marc47/go-oraweb/releases)
[![Linux Build](https://img.shields.io/travis/marc47marc47/go-oraweb.svg?label=Linux)](https://travis-ci.org/marc47marc47/go-oraweb)
[![Windows Build](https://img.shields.io/appveyor/ci/marc47marc47/go-oraweb/master.svg?label=Windows)](https://ci.appveyor.com/project/marc47marc47/go-oraweb)

## Overview

go-oraweb is a web-based database browser for Oracle, written in Go and works
on OSX, Linux and Windows machines. Main idea behind using Go for backend development
is to utilize ability of the compiler to produce zero-dependency binaries for 
multiple platforms. go-oraweb was created as an attempt to build very simple and portable
application to work with local or remote Oracle databases.

[See application screenshots](SCREENS.md)

## Features

- Works on OSX, Linux and Windows
- Zero dependencies
- Simple installation (distributes as a single binary)
- Connect to local or remote servers
- Browse tables and table rows
- Get table details: structure, size, indeces, row count
- Run and analyze custom SQL queries
- Export table rows and query results as CSV
- Query history
- Server bookmarks

Visit [WIKI](https://github.com/marc47marc47/go-oraweb/wiki) for more details

## Installation

[Precompiled binaries](https://github.com/marc47marc47/go-oraweb/releases) for supported 
operating systems are available.

[More installation options](https://github.com/marc47marc47/go-oraweb/wiki/Installation)

## Usage

Start server:

```
go-oraweb
```

You can also provide connection flags:

```
go-oraweb --host localhost --user myuser --db mydb
```

Connection URL scheme is also supported:

```
go-oraweb --url oracle://user:password@host:port/database?sslmode=[mode]
```

## Deploy on Heroku

[![Heroku Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy?template=https://github.com/marc47marc47/go-oraweb)

## Testing

Before running tests, make sure you have Oracle server running on `localhost:1521`
interface. Also, you must have `oracle` user that could create new databases
in your local environment. go-oraweb server should not be running at the same time.

Execute test suite:

```
make test
```

## Contribute

- Fork this repository
- Create a new feature branch for a new functionality or bugfix
- Commit your changes
- Execute test suite
- Push your code and open a new pull request
- Use [issues](https://github.com/marc47marc47/go-oraweb/issues) for any questions
- Check [wiki](https://github.com/marc47marc47/go-oraweb/wiki) for extra documentation

# go-oraweb
Web Client for Oracle DBs
