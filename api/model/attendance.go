package model

type Attendance struct {
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
