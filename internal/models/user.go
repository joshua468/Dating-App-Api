package models

type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name" binding:"required"`
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Age            int    `json:"age" binding:"required"`
	Gender         string `json:"gender" binding:"required"`
	Location       string `json:"location" binding:"required"`
	JobDescription string `json:"job_description"`
}
