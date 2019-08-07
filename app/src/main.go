package main

import (
    "net/http"
    "fmt"
    "github.com/labstack/echo"
)

func getIndexHandler(c echo.Context) error {
  return c.HTML(http.StatusOK, fmt.Sprintf("<h1>Hello</h1>",))
}

func main() {
    e := echo.New()
    e.GET("/", getIndexHandler)
    e.Logger.Fatal(e.Start(":8080"))
}
