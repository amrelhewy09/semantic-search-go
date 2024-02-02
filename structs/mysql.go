package structs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func ConnectToDatabse(username string, password string, host string, port int, database string) *MySQL {
	connection := username + ":" + password + "@tcp(" + host + ":" + fmt.Sprint(port) + ")/" + database + "?parseTime=true"
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	return &MySQL{
		db: db,
	}
}

func (mysql *MySQL) CreateTable() (sql.Result, error) {
	result, err := mysql.db.Exec("CREATE TABLE IF NOT EXISTS embeddings (text TEXT, embedding BLOB)")

	return result, err
}

func (mysql *MySQL) InsertEmbedding(text string, embedding []byte) (sql.Result, error) {
	res, err := mysql.db.Exec("INSERT INTO embeddings (text, embedding) VALUES (?, ?)", text, embedding)

	return res, err
}

func (mysql *MySQL) GetRelatedEmbeddings(embedding []byte) []string {
	res, err := mysql.db.Query("SELECT text, dot_product(embedding, ?) as similarity FROM embeddings ORDER BY similarity DESC LIMIT 3", embedding)
	if err != nil {
		log.Fatal("Error querying database:", err)
	}
	var relatedEmbeddings []string
	for res.Next() {
		var text string
		var similarity float32
		err = res.Scan(&text, &similarity)
		if err != nil {
			log.Fatal()
		}
		relatedEmbeddings = append(relatedEmbeddings, text)
	}

	return relatedEmbeddings
}

func (mysql *MySQL) Close() {
	mysql.db.Close()
}
