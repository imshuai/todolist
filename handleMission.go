package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func handleMission(c *gin.Context) {
	k := c.Param("id")
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
		v := bck.Get([]byte(k))
		if v == nil {
			return errors.New("invalid key")
		}
		m := &mission{}
		err := m.Deserialize(v)
		if err != nil {
			return err
		}
		ms[k] = m
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
