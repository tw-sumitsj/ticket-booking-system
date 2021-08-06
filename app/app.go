package app

import (
	"fmt"
)

func StartServer() {
	r := CreateRouter()
	fmt.Println("Starting the Server")
	r.Run(":8080")
}
