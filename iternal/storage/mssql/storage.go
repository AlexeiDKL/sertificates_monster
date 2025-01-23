package mssql

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/url"
	"strconv"

	mssql "github.com/denisenkom/go-mssqldb"
)

const (
	selectNullFilter   = "SELECT * FROM Partners;"
	selectPhraseFilter = "select * from Phrases where Used = 'false';"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
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

func bd() (*sql.DB, error) {
	connString := makeConnURL().String()
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	connector, err := mssql.NewConnector(connString)
	if err != nil {
		return nil, err
	}

	connector.SessionInitSQL = "SET ANSI_NULLS ON"
	db := sql.OpenDB(connector)

	return db, nil
}

func GetPhrase() (string, error) {
	// подключаемся к бд
	// селект к таблице Phrases
	// возвращаем первую фразу у которой поле Used = false

	db, err := bd()
	if err != nil {
		return "", err
	}
	defer db.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var phrase, used string
	var id int
	rows := db.QueryRowContext(ctx, selectPhraseFilter)
	err = rows.Scan(&id, &phrase, &used)
	if err != nil {
		return "", err
	}
	fmt.Printf("ID: %d, Phrase: %s, Used: %s\n", id, phrase, used)
	return phrase, nil
}

func SaveCertificate(phrase, certificate string) error {
	// подключаемся к бд
	// селект к таблице Phrases в которой phrase = phrase
	// если в поле Used = false нет ни одной записи,
	// меняем поле Used = true у текущей записи
	// и сохраняем изменения

	return nil
}
