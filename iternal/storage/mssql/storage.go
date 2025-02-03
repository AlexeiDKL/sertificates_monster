package mssql

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"flag"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strconv"
	"sync"
	"time"

	config "dkl.dklsa.certificates_monster/iternal/config"
	mssql "github.com/denisenkom/go-mssqldb"
)

type Phrase struct {
	Id        int    `json:"id"`
	Phrase    string `json:"phrase"`
	Used      bool   `json:"used"`
	TotalRows int    `json:"total_rows"`
}

const (
	selectNullFilter   = "SELECT * FROM Partners;"
	selectPhraseFilter = "SELECT Id, Phrase, Used, (select COUNT(*) FROM Phrases where Used='false') as total_rows FROM Phrases where Used='false';"
	insertPhraseFilter = "INSERT INTO Phrases (Phrase, Used) VALUES ('%s','%s')%s"
)

var WG sync.WaitGroup

func makeConnURL() *url.URL {
	flag.Parse()
	if config.Config.Logger.Level == "DEBUG" {
		fmt.Printf(" password:%s\n", config.Config.Storages.Password)
		fmt.Printf(" port:%d\n", config.Config.Storages.Port)
		fmt.Printf(" server:%s\n", config.Config.Storages.Server)
		fmt.Printf(" user:%s\n", config.Config.Storages.User)
	}

	var userInfo *url.Userinfo
	if config.Config.Storages.User != "" {
		userInfo = url.UserPassword(config.Config.Storages.User, config.Config.Storages.Password)
	}
	return &url.URL{
		Scheme: "sqlserver",
		Host:   config.Config.Storages.Server + ":" + strconv.Itoa(config.Config.Storages.Port),
		User:   userInfo,
	}
}

func BD() (*sql.DB, error) {
	connString := makeConnURL().String()
	if config.Config.Logger.Level == "DEBUG" {
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
	db, err := BD()
	if err != nil {
		return "", err
	}
	defer db.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var phraseStruct Phrase

	rows := db.QueryRowContext(ctx, selectPhraseFilter)
	err = rows.Scan(&phraseStruct.Id, &phraseStruct.Phrase, &phraseStruct.Used, &phraseStruct.TotalRows)
	if err != nil {
		if err == sql.ErrNoRows {
			// InsertPhrases(count)
			phrase := CreatePhrase()
			SavePhrase(phrase)
			InsertPhrases(config.Config.Storages.NumberOfSparePhrases * 2)
			return phrase, nil
		}
		return "", err
	}
	fmt.Printf("ID: %d, Phrase: %s, Used: %t, Count: %d \n", phraseStruct.Id, phraseStruct.Phrase, phraseStruct.Used, phraseStruct.TotalRows)
	if phraseStruct.TotalRows < config.Config.Storages.NumberOfSparePhrases {
		InsertPhrases(config.Config.Storages.NumberOfSparePhrases * 2)
	}
	return phraseStruct.Phrase, nil
}

func SavePhrase(phrase string) error {
	db, err := BD()
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}
	defer db.Close()

	stmt, err := db.PrepareContext(context.Background(), fmt.Sprintf(insertPhraseFilter, phrase, "false", ""))
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.Background())
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}

	return nil
}

func InsertPhrases(q int) {
	WG.Add(1)
	go func(n int) {
		fmt.Println(n)
		err := InsertPhrase(n)

		defer WG.Done()
		if err != nil {
			fmt.Printf("Error inserting phrase: %s\n", err.Error())
		} else {
			fmt.Printf("New phrase created and saved!\n")
		}

	}(q)
}

func InsertPhrase(n int) error {
	if n <= 0 {
		return fmt.Errorf("empty phrase")
	}

	db, err := BD()
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}
	defer db.Close()
	prepare := insertPhraseFilter
	for i := 0; i < n; i++ {

		phrase := CreatePhrase()
		used := "false"
		fmt.Println(phrase)
		if i == n-1 {
			prepare = fmt.Sprintf(prepare, phrase, used, "")
		} else {
			prepare = fmt.Sprintf(prepare, phrase, used, ",('%s','%s')%s")
		}
	}

	stmt, err := db.Prepare(prepare)
	if err != nil {
		return fmt.Errorf("prepare context failed: %v", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("insert phrase failed: %v", err)
	}
	return nil
}

func randRange(max int) int {
	return rand.IntN(max-0) + 0
}

func CreatePhrase() string {
	// генерируем фразу
	h := sha256.New()
	h.Write([]byte(time.Now().String()))

	start := randRange(len(h.Sum(nil)) - 17)
	return fmt.Sprintf("%x", h.Sum(nil))[start : start+16] // формируем 16 символов из хэша
}

func SaveCertificate(phrase, certificate string) error {
	// подключаемся к бд
	// селект к таблице Phrases в которой phrase = phrase
	// если в поле Used = false нет ни одной записи,
	// меняем поле Used = true у текущей записи
	// и сохраняем изменения

	return nil
}
