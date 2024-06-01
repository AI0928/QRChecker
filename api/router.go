package main

import (
	"net/http"

	"github.com/AI0928/QRChecker/api/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	//echoのインスタンス
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	//db接続
	model.Connect()

	// CORS対策
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000", "127.0.0.1:8000", "http://localhost:1323"}, // ドメイン
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.GET("/healthy", getHealthy)
	e.GET("/users", getUsers)
	e.GET("/users/:user_id", getUser)
	e.GET("/lectures", getLectures)
	e.GET("/lecture/:lecture_id", getLecture)
	e.GET("/attendances", getAttendances)
	e.GET("/attendances/:user_id", getAttendance)
	e.PUT("/attendances/:lecture_id/:user_id/:count/:atenndanceStatus", updateAttendance)

	return e
}
