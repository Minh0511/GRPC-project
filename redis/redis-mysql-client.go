package main

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	redisClient := newRedisClient()
	result, err := redisPing(redisClient)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//connect redis with mysql
	nameInRedis, err := redisClient.Get("MovieName").Result()
	if err != nil {
		fmt.Println(err)
	} else if err == redis.Nil {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Movies")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		stmtOut, err := db.Prepare("SELECT * FROM Movies WHERE MovieName = ?")
		if err != nil {
			panic(err.Error())
		}
		defer stmtOut.Close()

		rows, err := stmtOut.Query(nameInRedis)
		if err != nil {
			panic(err.Error())
		}

		numRows := 0
		for rows.Next() {
			var nameInSQL string
			err = rows.Scan(&nameInSQL)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(nameInSQL)
			numRows++
		}
		if numRows == 0 {
			fmt.Println("No movie found")
		}
	} else {
		fmt.Println(nameInRedis)
	}

}

func newRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return redisClient
}

func redisPing(client *redis.Client) (string, error) {
	result, err := client.Ping().Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
