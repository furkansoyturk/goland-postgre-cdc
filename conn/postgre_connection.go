package conn

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func PostgresConnect() {
	databaseUrl := "postgres://postgres:postgres@localhost:5432/postgres"
	conn, err := pgx.Connect(context.Background(), databaseUrl)

	defer conn.Close(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// sample query to test connection
	var fakeTitle string
	var fakeDesc string
	err = conn.QueryRow(context.Background(), "select fake_title, fake_desc from fake_data").Scan(&fakeTitle, &fakeDesc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(fakeDesc, fakeDesc)

}
