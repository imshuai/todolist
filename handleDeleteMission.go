package main

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func handleDeleteMission(c *gin.Context) {
	id := c.Param("id")
	err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("missions")); err != nil {
			log.Println("create", err)
			return err
		}

		bck := tx.Bucket([]byte("missions"))
		return bck.Delete([]byte(id))
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
		StatusMsg:  "ok",
		Data:       nil,
	})
	return
}
