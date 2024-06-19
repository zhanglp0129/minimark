package utils

import (
	"fmt"
	"testing"
)

func TestCreateJwtToken(t *testing.T) {
	data := map[string]any{
		"username": "zhanglp0129",
	}
	token, err := CreateJwtToken(data)
	fmt.Println(token, err)
}

func TestParseJwtToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc0OTg5NzUsImlzcyI6Im1pbmltYXJrIiwidXNlcm5hbWUiOiJ6aGFuZ2xwMDEyOSJ9.EjKH-VOWkZdBAdQh3e3TvX6GrqlIjAdXSqWuHR3gcvo"
	token, err := ParseJwtToken(tokenString)
	fmt.Println(token, err)
}
