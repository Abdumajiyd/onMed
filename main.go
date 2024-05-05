package main

import (
    "github.com/gin-gonic/gin"
    "hospital-system/config"
    "hospital-system/handlers"
)

func main() {
    r := gin.Default()
    db := config.SetupDatabaseConnection()
    hospitalHandler := handlers.HospitalHandler{DB: db}

    r.GET("/hospitals", hospitalHandler.GetHospitals)
    r.POST("/hospitals", hospitalHandler.CreateHospital)

    r.Run() // localhost:8080 da serverni ishga tushiradi
}
