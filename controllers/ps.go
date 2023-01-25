package controllers

import (
	"FP/database"
	"FP/repository"
	"FP/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllPS(c *gin.Context) {
	var (
		result gin.H
	)

	ps, err := repository.GetAllPS(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": ps,
		}
	}

	c.JSON(http.StatusOK, result)
}

func AddPS(c *gin.Context) {
	var ps structs.MasterPS

	err := c.ShouldBindJSON(&ps)

	if err != nil {
		panic(err)
	}

	err = repository.AddPS(database.DbConnection, ps)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil input PS",
	})
}

func UpdatePS(c *gin.Context) {
	var ps structs.MasterPS

	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&ps)
	if err != nil {
		panic(err)
	}

	ps.IdPS = int16(id)

	err = repository.UpdatePS(database.DbConnection, ps)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"results": "Sukses Update PS",
	})
}

func DeletePS(c *gin.Context) {
	var ps structs.MasterPS
	id, err := strconv.Atoi(c.Param("id"))

	ps.IdPS = int16(id)

	err = repository.DeletePS(database.DbConnection, ps)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Sukses Hapus PS",
	})
}
