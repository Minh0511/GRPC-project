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
	SyncCustomerInfo(es, db)
}

type SyncDB struct {
	CustomerID   int     `db:"CustomerID"`
	CustomerName string  `db:"CustomerName"`
	Store        string  `db:"Store"`
	Lat          float32 `db:"lat"`
	Lon          float32 `db:"lon"`
	ReadyToPush  int     `db:"ReadyToPush"`
	Version      int     `db:"Version"`
}

type PushData struct {
	CustomerID   int    `json:"CustomerID"`
	CustomerName string `json:"CustomerName"`
	Store        string `json:"Store"`
	Location     string `json:"location"`
	ReadyToPush  int    `json:"ReadyToPush"`
	Version      int    `json:"Version"`
}

func NewPushData(ds SyncDB) PushData {
	return PushData{
		CustomerID:   ds.CustomerID,
		CustomerName: ds.CustomerName,
		Store:        ds.Store,
		Location:     fmt.Sprintf("%f,%f", ds.Lat, ds.Lon),
		ReadyToPush:  ds.ReadyToPush,
		Version:      ds.Version,
	}
}

func SyncCustomerInfo(es *elasticsearch.Client, db *sqlx.DB) {
	start := time.Now()
	for i := 0; i < 1; i++ {
		listCustomer := `SELECT CustomerID, CustomerName, Store, lat, lon, ReadyToPush, Version 
							FROM CustomerInfo WHERE CustomerID > ? AND CustomerID < ?`
		var dbList []*SyncDB
		err := db.Select(&dbList, listCustomer, i, 1001)
		if err != nil {
			panic(err)
		}

		bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:  "customerinfo",
			Client: es,
		})
		generate := make([]interface{}, 0)
		placeholders := make([]string, 0)
		for j := 0; j < 1000; j++ {
			data := NewPushData(*dbList[j])
			els, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			err = bi.Add(context.Background(),
				esutil.BulkIndexerItem{
					Action:     "index",
					DocumentID: strconv.Itoa(data.CustomerID),
					Body:       bytes.NewReader(els),
				})
			if err != nil {
				panic(err)
			}
			generate = append(generate, data.CustomerID, "", "", "", "", 1, data.Version)
			placeholders = append(placeholders, "(?,?,?,?,?,?,?)")
		}
		tx := db.MustBegin()
		query := fmt.Sprintf("INSERT INTO CustomerInfo(CustomerID, CustomerName, Store, lat, lon, ReadyToPush, Version) VALUES %s"+
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
