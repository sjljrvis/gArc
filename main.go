package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	helpers "github.com/sjljrvis/gArch/helpers"
)

func main() {
	mac := helpers.GetMacAddress()
	fmt.Println("MAC Address ->", mac)
	dirWatcher()
}

func dirWatcher() {
	watcher, err := fsnotify.NewWatcher()
	done := make(chan bool)

	defer watcher.Close()

	if err != nil {
		log.Println("Error Occured in Watcher")
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				handleAction(event)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(os.Getenv("HOME") + "/gArch")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func handleAction(event fsnotify.Event) {
	fmt.Println(event.Name, "-------", event.Op)
	fmt.Println("Take Action")
}
