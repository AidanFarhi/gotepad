package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func main() {
	dir, err := os.MkdirTemp("", "test-")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	fn := filepath.Join(dir, "db")
	db, err := sql.Open("sqlite", fn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Query("select 'hello';")
}
