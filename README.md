# Postgresr

**Postgresr** (pronounced post-gres-er) is a thin wrapper round the library [pgx](https://github.com/jackc/pgx). It provides interfaces and mocks that can be used for establishing real connections, as well as mocking your database interactions for unit testing. As of this writing this does not implement all methods on **pgx.Conn**, but it implements enough for my purposes.

## Install

`go get github.com/app-nerds/postgrer`

## Example

```go
package main

import (
  "context"

  "github.com/app-nerds/postgresr"
)

func main() {
  var (
    err error
	 db postgresr.Conn
  )

  if db, err = postgresr.Connect(context.Background(), "host=localhost dbname=example user=user password=password"); err != nil {
    panic("cannot connect to database!")
  }


}

```

