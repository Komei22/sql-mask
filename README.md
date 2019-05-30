# sql-mask [![Build Status](https://travis-ci.org/Komei22/sql-mask.svg?branch=master)](https://travis-ci.org/Komei22/sql-mask)
sql-mask is a tool to mask literal in SQL queries. sql-mask support MySQL and PostgreSQL.

## Installation

```
go get -u github.com/Komei22/sql-mask
```

## Usage
### Masking MySQL query

```go
package main

import (
	"fmt"
	"github.com/Komei22/sql-mask"
)

func main() {
	query := "SELECT * FROM `user` WHERE `id` = 1"

	m := &masker.MysqlMasker{}
	queryDigest, err := masker.Mask(m, query)
	if err != nil {
		panic(err)
	}

	fmt.Println(queryDigest)
}
```
Running will output the query digest which is masked the literal value of input MySQL query:

```sh
$ go run main.go
SELECT * FROM `user` WHERE `id` = ?
```

### Masking PostgreSQL query

```go
package main

import (
	"fmt"
	"github.com/Komei22/sql-mask"
)

func main() {
	query := `SELECT * FROM "user" WHERE "id" = 1`

	m := &masker.PgMasker{}
	queryDigest, err := masker.Mask(m, query)
	if err != nil {
		panic(err)
	}

	fmt.Println(queryDigest)
}
```
Running will output the query digest which is masked the literal value of input PostgreSQL query:

```sh
$ go run main.go
SELECT * FROM "user" WHERE "id" = ?
```
