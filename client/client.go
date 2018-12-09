package client

import "fmt"

func Client() {
	fmt.Println("Hello this is client")
}

// err := helpers.CheckDir(os.Getenv("HOME") + "/gArch")
// if err != nil {
// 	log.Println(err)
// } else {
// 	fmt.Println("Creating Directory :" + os.Getenv("HOME") + "/gArch")
// 	os.Mkdir(os.Getenv("HOME")+"/gArch", 0700)

// }
// for true {
// 	go routines.CheckNewFiles(os.Getenv("HOME") + "/gArch")
// }
