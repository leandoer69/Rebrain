// файл main.go
package main

import (
	"Rebrain/module08/internal/godb"
	_ "Rebrain/module08/internal/migrations"
	"Rebrain/module08/pkg/helpers/pg"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"os"
)

func main() {
	cfg := &pg.Config{
		Host:     "localhost",
		Port:     "54320",
		Username: "db_user",
		Password: "pwd123",
		DbName:   "db_test",
		Timeout:  5,
	}

	poolConfig, err := pg.NewPoolConfig(cfg)
	if err != nil {
		fmt.Printf("Pool config error: %v\n", err)
		os.Exit(1)
	}

	poolConfig.MaxConns = 5

	conn, err := pg.NewConnection(poolConfig)
	if err != nil {
		fmt.Printf("Connection to DB failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK!")

	mdb, _ := sql.Open("postgres", poolConfig.ConnString())
	err = mdb.Ping()
	if err != nil {
		panic(err)
	}

	err = goose.Up(mdb, "C:\\Users\\Глеб\\GolandProjects\\Rebrain\\module08\\internal\\migrations")
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(context.Background(), ";")
	if err != nil {
		fmt.Printf("Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Ping OK!")

	ins := &godb.Instance{Db: conn}
	ins.Start()
}
