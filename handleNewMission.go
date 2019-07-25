package main

import (
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func handleNewMission(c *gin.Context) {
	m := &mission{}
	ms := make(map[string]*mission, 0)
	err := c.BindJSON(m)
	if err != nil {
		c.JSON(http.StatusOK, msg{
			StatusCode: 400,
			StatusMsg:  err.Error(),
			Data:       nil,
		})
		return
	}
	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("missions")); err != nil {
			return err
		}
		bck := tx.Bucket([]byte("missions"))
		id, _ := bck.NextSequence()
		v, err := m.Serialize()
		if err != nil {
			return err
		}
		k := strconv.FormatUint(id, 10)
		err = bck.Put([]byte(k), v)
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
