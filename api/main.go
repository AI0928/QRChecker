package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AI0928/QRChecker/api/model"
	"github.com/labstack/echo/v4"
)

func getUsers(c echo.Context) error {
	users := []model.User{}
	model.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func getLectures(c echo.Context) error {
	Lectures := []model.Lecture{}
	model.DB.Find(&Lectures)
	//fmt.Println(Lectures)
	return c.JSON(http.StatusOK, Lectures)
}

func getLecture(c echo.Context) error {
	lecture := model.Lecture{}
	if err := c.Bind(&lecture); err != nil {
		return err
	}
	model.DB.Take(&lecture)
	return c.JSON(http.StatusOK, lecture)
}

func getAttendances(c echo.Context) error {
	attendances := []model.Attendance{}
	model.DB.Find(&attendances)
	//fmt.Println(attendances)
	return c.JSON(http.StatusOK, attendances)
}

func updateAttendance(c echo.Context) error {
	//fmt.Println("出席")
	attendance := model.Attendance{}
	attendances := []model.Attendance{}
	model.DB.Find(&attendances)
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

	model.DB.Save(&attendance)
	return c.JSON(http.StatusCreated, attendance)
}

func getAttendance(c echo.Context) error {
	//attendance := Attendances{}
	attendances := []model.Attendance{}
	returnAttendance := []model.Attendance{}

	model.DB.Find(&attendances)
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
	router := newRouter()
	router.Logger.Fatal(router.Start(":1323"))

	sqlDB, _ := model.DB.DB()
	defer sqlDB.Close()
}
