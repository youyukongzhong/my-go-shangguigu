package main

import "learnGo/router"

func main() {
	r := router.Router()
	r.Run(":8081")
}
