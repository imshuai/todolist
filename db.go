package main

import (
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db *bolt.DB
)

func dbInit() error {
	var err error
	db, err = bolt.Open("missions.db", os.ModePerm, &bolt.Options{
		Timeout: time.Second * 10,
	})
	if err != nil {
		return err
	}
	return nil
}
