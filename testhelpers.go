package postgresr

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

type RowCounterFunc func() (*int, *int)

func InitializeRowCounterFunc(currentRowTracker *int, totalRowsTracker *int, totalRows int) RowCounterFunc {
	currentRowTracker = new(int)
	totalRowsTracker = new(int)

	*currentRowTracker = -1
	*totalRowsTracker = totalRows

	return func() (*int, *int) {
		return currentRowTracker, totalRowsTracker
	}
}

func DataToRows(rows [][]interface{}, rowCounterFunc RowCounterFunc) pgx.Rows {
	currentRow, _ := rowCounterFunc()

	if *currentRow > -1 {
		*currentRow = -1
	}

	return &MockRows{
		CloseFunc: func() {},
		GetTotalRowsFunc: func() uint64 {
			return uint64(len(rows))
		},
		NextFunc: func() bool {
			currentRow, totalRows := rowCounterFunc()
			*currentRow++

			result := *currentRow < *totalRows
			return result
		},
		ScanFunc: func(dest ...interface{}) error {
			currentRow, _ := rowCounterFunc()
			data := rows[*currentRow]

			for index, d := range dest {
				switch t := d.(type) {
				case *bool:
					p := d.(*bool)
					*p = data[index].(bool)

				case *string:
					p := d.(*string)
					*p = data[index].(string)

				case *int:
					p := d.(*int)
					*p = data[index].(int)

				case *int32:
					p := d.(*int32)
					*p = data[index].(int32)

				case *int64:
					p := d.(*int64)
					*p = data[index].(int64)

				case *float32:
					p := d.(*float32)
					*p = data[index].(float32)

				case *float64:
					p := d.(*float64)
					*p = data[index].(float64)

				case *time.Time:
					p := d.(*time.Time)
					*p = data[index].(time.Time)

				case *sql.NullBool:
					v, _ := data[index].(bool)
					result := sql.NullBool{Bool: v, Valid: true}
					p := d.(*sql.NullBool)
					*p = result

				case *sql.NullTime:
					v, _ := data[index].(time.Time)
					result := sql.NullTime{Time: v, Valid: true}
					p := d.(*sql.NullTime)
					*p = result

				case *sql.NullInt16:
					v, _ := data[index].(int16)
					result := sql.NullInt16{Int16: v, Valid: true}
					p := d.(*sql.NullInt16)
					*p = result

				case *sql.NullInt64:
					v, _ := data[index].(int64)
					result := sql.NullInt64{Int64: v, Valid: true}
					p := d.(*sql.NullInt64)
					*p = result

				case *sql.NullInt32:
					v, _ := data[index].(int32)
					result := sql.NullInt32{Int32: v, Valid: true}
					p := d.(*sql.NullInt32)
					*p = result

				case *sql.NullString:
					v, _ := data[index].(string)
					result := sql.NullString{String: v, Valid: true}
					p := d.(*sql.NullString)
					*p = result

				case *sql.NullFloat64:
					v, _ := data[index].(float64)
					result := sql.NullFloat64{Float64: v, Valid: true}
					p := d.(*sql.NullFloat64)
					*p = result

				default:
					fmt.Printf("undefined type '%T' for value '%v', skipping.", t, data[index])
				}
			}

			return nil
		},
		ValuesFunc: func() ([]interface{}, error) {
			currentRow, _ := rowCounterFunc()
			return rows[*currentRow], nil
		},
	}
}

func MockQuerySuccessHelper(rows [][]interface{}) func(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rowIndex := -1

	return func(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
		return &MockRows{
			CloseFunc: func() {},
			GetTotalRowsFunc: func() uint64 {
				return uint64(len(rows))
			},
			NextFunc: func() bool {
				rowIndex++
				return rowIndex < len(rows)
			},
			ScanFunc: func(dest ...interface{}) error {
				data := rows[rowIndex]

				for index, d := range dest {
					switch t := d.(type) {
					case *bool:
						p := d.(*bool)
						*p = data[index].(bool)

					case *string:
						p := d.(*string)
						*p = data[index].(string)

					case *int:
						p := d.(*int)
						*p = data[index].(int)

					case *int32:
						p := d.(*int32)
						*p = data[index].(int32)

					case *int64:
						p := d.(*int64)
						*p = data[index].(int64)

					case *float32:
						p := d.(*float32)
						*p = data[index].(float32)

					case *float64:
						p := d.(*float64)
						*p = data[index].(float64)

					case *time.Time:
						p := d.(*time.Time)
						*p = data[index].(time.Time)

					case *sql.NullBool:
						v, _ := data[index].(bool)
						result := sql.NullBool{Bool: v, Valid: true}
						p := d.(*sql.NullBool)
						*p = result

					case *sql.NullTime:
						v, _ := data[index].(time.Time)
						result := sql.NullTime{Time: v, Valid: true}
						p := d.(*sql.NullTime)
						*p = result

					case *sql.NullInt16:
						v, _ := data[index].(int16)
						result := sql.NullInt16{Int16: v, Valid: true}
						p := d.(*sql.NullInt16)
						*p = result

					case *sql.NullInt64:
						v, _ := data[index].(int64)
						result := sql.NullInt64{Int64: v, Valid: true}
						p := d.(*sql.NullInt64)
						*p = result

					case *sql.NullInt32:
						v, _ := data[index].(int32)
						result := sql.NullInt32{Int32: v, Valid: true}
						p := d.(*sql.NullInt32)
						*p = result

					case *sql.NullString:
						v, _ := data[index].(string)
						result := sql.NullString{String: v, Valid: true}
						p := d.(*sql.NullString)
						*p = result

					case *sql.NullFloat64:
						v, _ := data[index].(float64)
						result := sql.NullFloat64{Float64: v, Valid: true}
						p := d.(*sql.NullFloat64)
						*p = result

					default:
						fmt.Printf("undefined type '%T' for value '%v', skipping.", t, data[index])
					}
				}

				return nil
			},
			ValuesFunc: func() ([]interface{}, error) {
				return rows[rowIndex], nil
			},
		}, nil
	}
}
