package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kubosaka/qrchecker/api/drivers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id         int    `json:"id" `
	User_Id    int    `json:"user_id" param:"user_id"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	Major      string `json:"major"`
	Course     string `json:"course"`
	Grade      int    `json:"grade"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Lectures struct {
	Id         int    `json:"id" param:"id"`
	Number     int    `json:"number" param:"lecture_id"`
	Name       string `json:"name"`
	Teacher    string `json:"teacher"`
	Year       string `json:"year"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
}

type Attendances struct {
	Id             int    `json:"id"`
	Lecture_Number int    `json:"lecture_number"  param:"lecture_id"`
	User_Number    int    `json:"user_number" param:"user_id"`
	First          string `json:"first"`
	Second         string `json:"second"`
	Third          string `json:"third"`
	Fourth         string `json:"fourth"`
	Fifth          string `json:"fifth"`
	Sixth          string `json:"sixth"`
	Seventh        string `json:"seventh"`
	Eighth         string `json:"eighth"`
	Ninth          string `json:"ninth"`
	Tenth          string `json:"tenth"`
	Eleventh       string `json:"eleventh"`
	Twelfth        string `json:"twelfth"`
	Thirteenth     string `json:"thirteenth"`
	Fourteenth     string `json:"fourteenth"`
	Fifteenth      string `json:"fifteenth"`
}

func getUsers(c echo.Context) error {
	users := []User{}
	drivers.DB.Find(&users)
	//fmt.Println(users)
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	drivers.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func getLectures(c echo.Context) error {
	Lectures := []Lectures{}
	drivers.DB.Find(&Lectures)
	//fmt.Println(Lectures)
	return c.JSON(http.StatusOK, Lectures)
}

func getLecture(c echo.Context) error {
	lecture := Lectures{}
	if err := c.Bind(&lecture); err != nil {
		return err
	}
	drivers.DB.Take(&lecture)
	return c.JSON(http.StatusOK, lecture)
}

func getAttendances(c echo.Context) error {
	attendances := []Attendances{}
	drivers.DB.Find(&attendances)
	//fmt.Println(attendances)
	return c.JSON(http.StatusOK, attendances)
}

func updateAttendance(c echo.Context) error {
	//fmt.Println("出席")
	attendance := Attendances{}
	attendances := []Attendances{}
	drivers.DB.Find(&attendances)
	if err := c.Bind(&attendance); err != nil {
		return err
	}
	fmt.Println(attendance)
	for _, a := range attendances {
		if a.Lecture_Number == attendance.Lecture_Number && a.User_Number == attendance.User_Number {
			attendance.Id = a.Id
			attendance.First = a.First
			attendance.Second = a.Second
			attendance.Third = a.Third
			attendance.Fourth = a.Fourth
			attendance.Fifth = a.Fifth
			attendance.Sixth = a.Sixth
			attendance.Seventh = a.Seventh
			attendance.Eighth = a.Eighth
			attendance.Ninth = a.Ninth
			attendance.Tenth = a.Tenth
			attendance.Eleventh = a.Eleventh
			attendance.Twelfth = a.Twelfth
			attendance.Thirteenth = a.Thirteenth
			attendance.Fourteenth = a.Fourteenth
			attendance.Fifteenth = a.Fifteenth
		}
	}

	//fmt.Println(c)
	switch c.Param("count") {
	case "1":
		attendance.First = c.Param("atenndanceStatus")
	case "2":
		attendance.Second = c.Param("atenndanceStatus")
	case "3":
		attendance.Third = c.Param("atenndanceStatus")
	case "4":
		attendance.Fourth = c.Param("atenndanceStatus")
	case "5":
		attendance.Fifth = c.Param("atenndanceStatus")
	case "6":
		attendance.Sixth = c.Param("atenndanceStatus")
	case "7":
		attendance.Seventh = c.Param("atenndanceStatus")
	case "8":
		attendance.Eighth = c.Param("atenndanceStatus")
	case "9":
		attendance.Ninth = c.Param("atenndanceStatus")
	case "10":
		attendance.Tenth = c.Param("atenndanceStatus")
	case "11":
		attendance.Eleventh = c.Param("atenndanceStatus")
	case "12":
		attendance.Twelfth = c.Param("atenndanceStatus")
	case "13":
		attendance.Thirteenth = c.Param("atenndanceStatus")
	case "14":
		attendance.Fourteenth = c.Param("atenndanceStatus")
	case "15":
		attendance.Fifteenth = c.Param("atenndanceStatus")
	}

	drivers.DB.Save(&attendance)
	return c.JSON(http.StatusCreated, attendance)
}

func getAttendance(c echo.Context) error {
	//attendance := Attendances{}
	attendances := []Attendances{}
	returnAttendance := []Attendances{}

	drivers.DB.Find(&attendances)
	//fmt.Println(attendances)

	// if err := c.Bind(&attendance); err != nil {
	// 	return err
	// }
	// fmt.Println(attendance)
	check_user_id, _ := strconv.Atoi(c.Param("user_id"))
	fmt.Println(check_user_id)
	for _, a := range attendances {
		if a.User_Number == check_user_id {
			returnAttendance = append(returnAttendance, a)
		}
	}

	//fmt.Println(returnAttendance)
	return c.JSON(http.StatusCreated, returnAttendance)
}

func getHealthy(c echo.Context) error {
	return c.JSON(http.StatusOK, "true")
}

func main() {
	//echoのインスタンス
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	//db接続
	drivers.Connect()
	sqlDB, _ := drivers.DB.DB()
	defer sqlDB.Close()

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

	e.Logger.Fatal(e.Start(":1323"))
}
