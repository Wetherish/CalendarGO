package main

import (
	server "Calendar/Server"
	dataaccess "Calendar/data-access"
)

func main() {
	dataaccess.StartDB()
	server.StartServer()
}
