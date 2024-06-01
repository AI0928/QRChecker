package model

type Lecture struct {
	Id         int    `json:"id" param:"id"`
	Number     int    `json:"number" param:"lecture_id"`
	Name       string `json:"name"`
	Teacher    string `json:"teacher"`
	Year       string `json:"year"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
}
