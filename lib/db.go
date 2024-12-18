package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {

	//conn, err := pgx.Connect(context.Background(), "postgres://postgres:123@postgres:5432/konis_caffee?sslmode=disable")

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123@localhost:5433/konis_caffee?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	return conn
}
