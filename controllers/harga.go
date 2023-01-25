package controllers

import (
	"FP/database"
	"FP/repository"
	"FP/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllHarga(c *gin.Context) {
	var (
		result gin.H
	)

	harga, err := repository.GetAllHarga(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": harga,
		}
	}

	c.JSON(http.StatusOK, result)
}

func AddHarga(c *gin.Context) {
	var harga structs.Harga

	err := c.ShouldBindJSON(&harga)

	if err != nil {
		panic(err)
	}

	err = repository.AddHarga(database.DbConnection, harga)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil input Harga",
	})
}

func UpdateHarga(c *gin.Context) {
	var harga structs.Harga

	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&harga)
	if err != nil {
		panic(err)
	}

	harga.IdHarga = int16(id)

	err = repository.UpdateHarga(database.DbConnection, harga)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"results": "Sukses Update Harga",
	})
}

func DeleteHarga(c *gin.Context) {
	var harga structs.Harga
	id, err := strconv.Atoi(c.Param("id"))

	harga.IdHarga = int16(id)

	err = repository.DeleteHarga(database.DbConnection, harga)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Sukses Hapus Harga",
	})
}
