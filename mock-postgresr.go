package postgresr

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgx/v4"
)

type MockConn struct {
	CloseFunc    func(ctx context.Context) error
	ExecFunc     func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	PingFunc     func(ctx context.Context) error
	QueryFunc    func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRowFunc func(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type MockRows struct {
	CloseFunc             func()
	CommandTagFunc        func() pgconn.CommandTag
	ErrFunc               func() error
	FieldDescriptionsFunc func() []pgproto3.FieldDescription
	GetTotalRowsFunc      func() uint64
	NextFunc              func() bool
	ScanFunc              func(dest ...interface{}) error
	ValuesFunc            func() ([]interface{}, error)
	RawValuesFunc         func() [][]byte
}

type MockRow struct {
	ScanFunc func(dest ...interface{}) error
}

func (m *MockConn) Close(ctx context.Context) error {
	return m.CloseFunc(ctx)
}

func (m *MockConn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return m.ExecFunc(ctx, sql, arguments...)
}

func (m *MockConn) Ping(ctx context.Context) error {
	return m.PingFunc(ctx)
}

func (m *MockConn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return m.QueryFunc(ctx, sql, args...)
}

func (m *MockConn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return m.QueryRowFunc(ctx, sql, args...)
}

func (m *MockRows) Close() {
	m.CloseFunc()
}

func (m *MockRows) CommandTag() pgconn.CommandTag {
	return m.CommandTagFunc()
}

func (m *MockRows) Err() error {
	return m.ErrFunc()
}

func (m *MockRows) FieldDescriptions() []pgproto3.FieldDescription {
	return m.FieldDescriptionsFunc()
}

func (m *MockRows) GetTotalRows() uint64 {
	return m.GetTotalRowsFunc()
}

func (m *MockRows) Next() bool {
	return m.NextFunc()
}

func (m *MockRows) Scan(dest ...interface{}) error {
	return m.ScanFunc(dest...)
}

func (m *MockRows) Values() ([]interface{}, error) {
	return m.ValuesFunc()
}

func (m *MockRows) RawValues() [][]byte {
	return m.RawValuesFunc()
}

func (m *MockRow) Scan(dest ...interface{}) error {
	return m.ScanFunc(dest...)
}
