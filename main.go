package main

import (
	"test/go_goroutine"
	"test/study_client"
	"test/study_server"
)

func main() {

	a := 3

	switch a {
	case 1:
		study_client.ClientHandler()
	case 2:
		study_server.ServerHanlder()
	case 3:
		//go_goroutine.TestGo()
		//go_goroutine.TestGo2()
		//go_goroutine.TestGo3()
		go_goroutine.Testgo4()
	}

}
