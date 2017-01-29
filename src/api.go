package main

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

type User struct {
    Firstname string `json:"name,omitempty"`
    Lastname string `json:"lastname,omitempty"`
    Username string `json:"username,omitempty"`
    Password string `json:"password,omitempty"`
}

func index(c echo.Context) error {
  // return c.String(http.StatusOK, "Hello World")
    // return c.JSON(http.StatusOK, "/{"status": "you are logged in"}")
    return c.JSON(http.StatusOK, "Hello World")
}

func getUsers(c echo.Context) error {
  pat := User{
    Firstname:"aosmin",
    Lastname:"leekdee",
    Username:"merxer",
  }
  return c.JSON(http.StatusOK, pat)
}

func getUserByID(c echo.Context) error {
  id := c.Param("id")
  return c.JSON(http.StatusOK, id)
}
func saveUser(c echo.Context) error {
  user := new(User)
  if err := c.Bind(user);err != nil {
    return c.JSON(http.StatusBadRequest, nil)
  }
  return c.JSON(http.StatusOK, user)
}

func main()  {
  e := echo.New()
  e.Use(middleware.Logger())
  e.GET("/", index)
  e.GET("/users", getUsers)
  e.GET("/users/:id", getUserByID)
  e.POST("/users", saveUser)
  e.Logger.Fatal(e.Start(":1324"))
}
