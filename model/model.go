package model

import "time"

var (
	ErQuize    string = `{"code":"unknown","message":"Not a quiz slide"}`
	ErNotfound string = `{"code":"unknown","message":"Not Found"}`
)

type JsonClient struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Courses struct {
	Courses []*Course `json:"courses"`
}

type Course struct {
	Course Id `json:"course"`
}

type Id struct {
	Id int `json:"id"`
}

type Lessons struct {
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	CourseId          int       `json:"course_id"`
	CreatedAt         time.Time `json:"created_at"`
	Id                int       `json:"id"`
	LastViewedSlideId int       `json:"last_viewed_slide_id"`
	Number            int       `json:"number"`
	Status            string    `json:"status"`
	Title             string    `json:"title"`
	UserId            int       `json:"user_id"`
}

type Announcement struct {
	Title    string
	Deadline *time.Time
}
