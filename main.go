package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var DataStore map[string]interface{} = make(map[string]interface{})

type Map struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {

	r := gin.Default()
	r.GET("/akansha", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.POST("/insert", insert)
	r.GET("/get", get)

	r.Run()

}

func insert(c *gin.Context) {
	var map1 Map
	data1, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data1, &map1)
	DataStore[map1.Key] = map1.Value

}
func get(c *gin.Context) {

	key := c.Request.URL.Query().Get("key")
	c.JSON(200, gin.H{
		"value": DataStore[key],
	})

}
