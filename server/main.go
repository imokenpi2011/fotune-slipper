package main

import (
	"log"

	"github.com/imokenpi2011/fotune-slipper/server/app/controllers"
)

func main() {
	log.Println("start main method.")
	controllers.StartMainServer()

	// fotune, err := services.GetFotuneResult()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(fotune)
	log.Println("End main method.")
}
