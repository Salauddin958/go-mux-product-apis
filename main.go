package main

const (
	APP_USER_NAME = "root"
	APP_DB_PWD    = "root"
	APP_DB_NAME   = "productdb"
)

func main() {
	a := App{}
	a.Initialize(
		APP_USER_NAME,
		APP_DB_PWD,
		APP_DB_NAME)
	a.Run(":9010")
}
