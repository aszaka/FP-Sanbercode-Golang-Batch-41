package main

import (
	"FP/controllers"
	"FP/database"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	//ENV Config
	err = godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("Gagal load env")
	} else {
		fmt.Println("Sukses load env")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()

	if err != nil {
		fmt.Println("Koneksi gagal")
		panic(err)
	} else {
		fmt.Println("Koneksi berhasil")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	//Router GIN
	router := gin.Default()
	router.GET("/user", controllers.GetAllUser)
	router.POST("/user/user", controllers.AddUser)
	router.POST("/user/admin", controllers.AddAdmin)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	router.GET("/ps", controllers.GetAllPS)
	router.POST("/ps", controllers.AddPS)
	router.PUT("/ps/:id", controllers.UpdatePS)
	router.DELETE("/ps/:id", controllers.DeletePS)

	router.GET("/harga", controllers.GetAllHarga)
	router.POST("/harga", controllers.AddHarga)
	router.PUT("/harga/:id", controllers.UpdateHarga)
	router.DELETE("/harga/:id", controllers.DeleteHarga)

	router.GET("/peminjaman", controllers.GetAllPeminjaman)
	router.POST("/peminjaman", controllers.AddPeminjaman)
	router.PUT("/peminjaman/:id", controllers.UpdatePeminjaman)
	router.DELETE("/peminjaman/:id", controllers.DeletePeminjaman)

	router.Run(":" + os.Getenv("PORT"))
	//router.Run("localhost:8080")
}
