package main

import "os"

func main() {
	a := App{}
	a.Initializer(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	a.Run(":8080")
}
