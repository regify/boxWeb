package main

// This is a simple web server example that listens for POST requests
// on /post and hands over the received json "data" to the regiboxctl
// command.
// It is not an official regify product and is not supported by regify.
// Use at your own risk!
import (
	"fmt"
	"os/exec"

	"github.com/labstack/echo/v4"
)

type JSONBody struct {
	Data string `json:"data"`
}

func main() {
	configDir := "/opt/regiboxd/config" // the regiboxd config folder
	listen := "localhost:8081"          // what interface and port to listen on

	e := echo.New()
	e.HideBanner = true

	e.POST("/post", func(c echo.Context) error {
		var body JSONBody
		err := c.Bind(&body)
		if err != nil {
			return c.String(500, "Bad request")
		}

		cmd := exec.Command("regiboxctl", "-c", configDir, "-t", "10", "-j", body.Data)
		result, err := cmd.Output()
		if err != nil {
			return c.String(500, "Error executing regiboxctl")
		}

		return c.String(200, string(result))
	})

	fmt.Println("                 _ __")
	fmt.Println("   _______ ___ _(_) /  ___ __ _")
	fmt.Println("  / __/ -_) _ `/ / _ \\/ _ \\\\ \\/")
	fmt.Println(" /_/  \\__/\\_, /_/_.__/\\___/_\\_\\ ")
	fmt.Println("         /___/")
	fmt.Println("")
	fmt.Println("This is not an official regify product!")
	fmt.Println("No support, no warranty. Use at your own risk!")

	e.Start(listen)
}
