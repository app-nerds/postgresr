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

## Testing

This library provides mock structures useful for unit tests. Here is an example of a test that mocks a Postgres variable passed to a function.

```go
func QueryForStuff(pg postgresr.Conn) ([]SomeStruct, error) {
	var (
		err error
		rows pgx.Rows
		result []SomeStruct
	)

	query := `SELECT * FROM that_table WHERE something='else'`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if rows, err = pg.Query(ctx, query); err != nil {
		// handle it
	}

	for rows.Next() {
		var (
			column1 string
			column2 int
		)

		if err = rows.Scan(&column1, &column2); err != nil {
			// handle it
		}

		result = append(result, SomeStruct{
			Column1: column1,
			Column2: column2,
		})
	}

	return result, nil
}

func TestQueryForStuff(t *testing.T) {
	testData := [][]interface{}{
		{
			"value1", // column1
			1,         // column2
		},
		{
			"value2",
			2,
		},
	}

	pg := &postgresr.MockConn{
		QueryFunc: postgresr.MockQuerySuccessHelper(testData),
	}

	want := []SomeStruct{
		{ Column1: "value1", Column2: 1 },
		{ Column1: "value2", Column2: 2 },
	}

	got, err := QueryForStuff(pg)

	if err != nil {
		t.Errorf("didn't expect an error!")
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted: %+v\ngot: %+v\n", want, got)
	}
}
```

If you need more control of what you return in query mocks, the `DataToRows` method might be useful.

```go
func TestQueryForStuff(t *testing.T) {
	var (
		data1CurrentRow *int
		data1TotalCount *int
		data2CurrentRow *int
		data2TotalCount *int
	)

	testData1 := [][]interface{}{
		{
			"value1", // column1
			1,         // column2
		},
		{
			"value2",
			2,
		},
	}

	testData2 := [][]interface{}{
		{ 1, "1" },
		{ 2, "2" },
	}

	data1Counter := postgresr.InitializeRowCounterFunc(data1CurrentRow, data1TotalCount, len(testData1))
	data2Counter := postgresr.InitializeRowCounterFunc(data2CurrentRow, data2TotalCount, len(testData2))

	pg := &postgresr.MockConn{
		QueryFunc: func(ctx context.Context, query string, arguments ...interface{}) (pgx.Rows, error) {
			if strings.Contains(query, "FROM table1") {
				return postgresr.DataToRows(testData1, data1Counter), nil
			}

			return postgresr.DataToRows(testData2, data2Counter), nil
		},
	}

	want := []SomeStruct{
		{ Column1: "value1", Column2: 1 },
		{ Column1: "value2", Column2: 2 },
	}

	got, err := QueryForStuff(pg)

	if err != nil {
		t.Errorf("didn't expect an error!")
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted: %+v\ngot: %+v\n", want, got)
	}
}
```
