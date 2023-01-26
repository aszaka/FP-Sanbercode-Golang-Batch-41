package controllers

import (
	"FP/database"
	"FP/repository"
	"FP/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllPeminjaman(c *gin.Context) {
	var (
		result gin.H
	)

	peminjaman, err := repository.GetAllPeminjaman(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": peminjaman,
		}
	}

	c.JSON(http.StatusOK, result)
}

func AddPeminjaman(c *gin.Context) {
	var peminjaman structs.Peminjaman

	err := c.ShouldBindJSON(&peminjaman)

	if err != nil {
		panic(err)
	}

	err = repository.AddPeminjaman(database.DbConnection, peminjaman)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil input Peminjaman",
	})
}

func UpdatePeminjaman(c *gin.Context) {
	var peminjaman structs.Peminjaman

	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&peminjaman)
	if err != nil {
		panic(err)
	}

	peminjaman.IdPeminjaman = int16(id)

	err = repository.UpdatePeminjaman(database.DbConnection, peminjaman)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"results": "Sukses Update Peminjaman",
	})
}

func DeletePeminjaman(c *gin.Context) {
	var peminjaman structs.Peminjaman
	id, err := strconv.Atoi(c.Param("id"))

	peminjaman.IdPeminjaman = int16(id)

	err = repository.DeletePeminjaman(database.DbConnection, peminjaman)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Sukses Hapus Peminjaman",
	})
}
