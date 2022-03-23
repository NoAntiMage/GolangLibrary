package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func main() {
	db, _ := bolt.Open("./wuji.db", 0600, nil)
	defer db.Close()

	// demo1: database -> transaction -> bucket -> kv -> b+tree
	tx, _ := db.Begin(true)
	defer tx.Rollback()

	var bucketName []byte = []byte("wujibucket")

	bucket, _ := tx.CreateBucketIfNotExists(bucketName)
	bucket.Put([]byte("foo"), []byte("bar"))
	tx.Commit()

	ret := bucket.Get([]byte("foo"))
	fmt.Println(string(ret))

	// demo2: bolt.db API
	db.Update(func(tx *bolt.Tx) error {
		bucket, _ := tx.CreateBucketIfNotExists([]byte("mybucket"))
		bucket.Put([]byte("hello"), []byte("wujimaster"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("mybucket"))
		val := bucket.Get([]byte("hello"))
		fmt.Println(string(val))
		return nil
	})

}
