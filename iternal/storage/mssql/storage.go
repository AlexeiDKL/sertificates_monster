package mssql

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/url"
	"strconv"

	mssql "github.com/denisenkom/go-mssqldb"
)

const (
	selectNullFilter = "SELECT * FROM Partners;"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", "Hello2025_", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "DESKTOP-1LPEPB6", "the database server")
	user          = flag.String("user", "user", "the database user")
)

func makeConnURL() *url.URL {
	flag.Parse()
	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	var userInfo *url.Userinfo
	if *user != "" {
		userInfo = url.UserPassword(*user, *password)
	}
	return &url.URL{
		Scheme: "sqlserver",
		Host:   *server + ":" + strconv.Itoa(*port),
		User:   userInfo,
	}
}

func bd() {
	connString := makeConnURL().String()
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	connector, err := mssql.NewConnector(connString)
	if err != nil {
		log.Println(err)
		return
	}

	connector.SessionInitSQL = "SET ANSI_NULLS ON"
	db := sql.OpenDB(connector)
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var name, inn string
	var id int
	rows := db.QueryRowContext(ctx, selectNullFilter)
	err = rows.Scan(&id, &name, &inn)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("ID: %d, Name: %s, INN: %s\n", id, name, inn)
}
