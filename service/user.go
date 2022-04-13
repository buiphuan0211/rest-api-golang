package service

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

type user struct {
	ID   string `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age  string `json:"age" form:"age"`
}

var users = make([]user, 0)

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		log.Fatal("unable to load configuration")
	}
}

// Middleware
func ServerMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside middleware")
		return next(c)
	}

}

func Start() {
	e := echo.New()

	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "1231"
	}

	e.GET("/users", GetUsers, ServerMessage)

	e.GET("/users/:id", GetUser)

	e.POST("/users", PostUser)

	e.PUT("/users/:id", PutUser)

	e.DELETE("/users/:id", DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"users": users,
	})
}

func GetUser(c echo.Context) (err error) {

	id := c.Param("id")

	if id == "" || len(id) < 32 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Khong tim thay thong tin user",
		})
	}

	for _, user := range users {
		if user.ID != id {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Khong tim thay thong tin user",
			})
		}

		return c.JSON(http.StatusOK, user)
	}

	return nil
}

func PostUser(c echo.Context) (err error) {
	payload := new(user)

	if err = c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Khong doc duoc du lieu",
		})
	}

	// validate
	if payload.Name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "khong duoc de trong",
		})
	}

	u := user{
		ID:   uuid.New().String(),
		Name: payload.Name,
		Age:  payload.Age,
	}

	users = append(users, u)

	return nil
}

func PutUser(c echo.Context) (err error) {
	payload := new(user)
	id := c.Param("id")

	if err = c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Khong doc duoc du lieu",
		})
	}

	if id == "" || len(id) < 32 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Khong tim thay thong tin user",
		})
	}

	log.Println(payload)

	for i, _ := range users {

		if users[i].ID != id {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Khong tim thay thong tin user",
			})
		}
		users[i].Name = payload.Name
		users[i].Age = payload.Age
		return c.JSON(http.StatusOK, echo.Map{})
	}
	return nil
}

func DeleteUser(c echo.Context) (err error) {
	id := c.Param("id")
	for i, _ := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, echo.Map{})
		}
	}
	return nil
}
