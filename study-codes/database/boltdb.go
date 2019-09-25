package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type response struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits response `json:"fruits"`
}

// nested json is possible
func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	test := response2{1, response{"apple", "peach", "pear"}}

	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		current, _ := json.Marshal(&test)
		err = b.Put([]byte("tt"), current)
		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("tt"))
		var res response2
		err1 := json.Unmarshal(v, &res)
		fmt.Println(res.Fruits.A)
		return err1
	})
}
