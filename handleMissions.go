package main

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func handleMissions(c *gin.Context) {
	ms := make(map[string]*mission, 0)
	err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("missions")); err != nil {
			log.Println("create", err)
			return err
		}
		return nil
	})
	if err != nil {
		c.JSON(http.StatusOK, msg{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			Data:       nil,
		})
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		bck := tx.Bucket([]byte("missions"))
		err := bck.ForEach(func(k, v []byte) error {
			m := &mission{}
			err := m.Deserialize(v)
			if err != nil {
				return err
			}
			ms[string(k)] = m
			return nil
		})
		if err != nil {
			log.Println("ergodic", err)
			return err
		}
		return nil
	})
	if err != nil {
		c.JSON(http.StatusOK, msg{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, msg{
		StatusCode: 200,
		StatusMsg:  "OK",
		Data:       ms,
	})
	return
}
