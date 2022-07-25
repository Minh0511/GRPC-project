package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	var r map[string]interface{}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	res, err := es.Info()
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	var db *sqlx.DB
	db, err = sqlx.Open("mysql", "root:1@tcp(0.0.0.0:3306)/Transaction")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	SyncCustomer(es, db)
}

type DataSync struct {
	TransactionID int    `db:"TransactionID"`
	CustomerName  string `db:"CustomerName"`
	Phone         string `db:"Phone"`
	Email         string `db:"Email"`
	Product       string `db:"Product"`
	ReadyToPush   int    `db:"ReadyToPush"`
	Version       int    `db:"Version"`
}

type DataPush struct {
	TransactionID int    `json:"TransactionID"`
	CustomerName  string `json:"CustomerName"`
	Phone         string `json:"Phone"`
	Email         string `json:"Email"`
	Product       string `json:"Product"`
	ReadyToPush   int    `json:"ReadyToPush"`
	Version       int    `json:"Version"`
}

func NewDataPush(ds DataSync) DataPush {
	return DataPush{
		TransactionID: ds.TransactionID,
		CustomerName:  ds.CustomerName,
		Phone:         ds.Phone,
		Email:         ds.Email,
		Product:       ds.Product,
		ReadyToPush:   ds.ReadyToPush,
		Version:       ds.Version,
	}
}

func SyncCustomer(es *elasticsearch.Client, db *sqlx.DB) {
	start := time.Now()
	for i := 1; i < 2; i++ {
		listCustomer := `SELECT TransactionID, CustomerName, Phone, Email, Product, ReadyToPush, Version 
							FROM Customer WHERE TransactionID > ? AND TransactionID < ?`
		var dbList []*DataSync
		err := db.Select(&dbList, listCustomer, i*1000, i*1000+1001)
		if err != nil {
			panic(err)
		}

		bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:  "customer",
			Client: es,
		})
		generate := make([]interface{}, 0)
		placeholders := make([]string, 0)
		for j := 0; j < 1000; j++ {
			data := NewDataPush(*dbList[j])
			els, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			err = bi.Add(context.Background(),
				esutil.BulkIndexerItem{
					Action:     "index",
					DocumentID: strconv.Itoa(data.TransactionID),
					Body:       bytes.NewReader(els),
				})
			if err != nil {
				panic(err)
			}
			generate = append(generate, data.TransactionID, "", "", "", "", 1, data.Version)
			placeholders = append(placeholders, "(?,?,?,?,?,?,?)")
		}
		tx := db.MustBegin()
		query := fmt.Sprintf("INSERT INTO Customer(TransactionID, CustomerName, Phone, Email, Product, ReadyToPush, Version) VALUES %s"+
			"ON DUPLICATE KEY UPDATE ReadyToPush = 0",
			strings.Join(placeholders, ","))
		tx.MustExec(query, generate...)
		tx.Commit()
		if err = bi.Close(context.Background()); err != nil {
			panic(err)
		}
	}
	log.Println("Time taken:", time.Since(start))
}
