package main

func main() {
	app := App{}

	app.Initialize("jones", "1qaz@WSX", "rest_api_example")

	app.Run(":8080")
}
