package postgresr

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

/******************************************************************************
 * Structs
 *****************************************************************************/

type PgxConn struct {
	conn *pgx.Conn
}

/******************************************************************************
 * Methods
 *****************************************************************************/

func Connect(ctx context.Context, connString string) (Conn, error) {
	var (
		err    error
		conn   *pgx.Conn
		result *PgxConn
	)

	if conn, err = pgx.Connect(ctx, connString); err != nil {
		return result, err
	}

	return &PgxConn{
		conn: conn,
	}, nil
}

func ConnectConfig(ctx context.Context, connConfig *pgx.ConnConfig) (Conn, error) {
	var (
		err    error
		conn   *pgx.Conn
		result *PgxConn
	)

	if conn, err = pgx.ConnectConfig(ctx, connConfig); err != nil {
		return result, err
	}

	return &PgxConn{
		conn: conn,
	}, nil
}

func (c *PgxConn) Close(ctx context.Context) error {
	return c.conn.Close(ctx)
}

func (c *PgxConn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return c.conn.Exec(ctx, sql, arguments...)
}

func (c *PgxConn) Ping(ctx context.Context) error {
	return c.conn.Ping(ctx)
}

func (c *PgxConn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return c.conn.Query(ctx, sql, args...)
}
