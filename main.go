package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	helpers "github.com/sjljrvis/gArch/helpers"
)

func main() {
	mac := helpers.GetMacAddress()
	fmt.Println("MAC Address ->", mac)

	dirWatcher()
}

func dirWatcher() {

	err := helpers.CheckDir(os.Getenv("HOME") + "/gArch")
	if err != nil {
		log.Println(err)
		goto initwatcher
	} else {
		fmt.Println("Creating Directory :" + os.Getenv("HOME") + "/gArch")
		os.Mkdir(os.Getenv("HOME")+"/gArch", 0700)
		goto initwatcher
	}

initwatcher:

	log.Println("Adding watcher to directory")
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
	fmt.Println("=====================================")
	fmt.Println(event.Name, "---", event.Op)
	files, err := filepath.Glob(os.Getenv("HOME") + "/gArch/*")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println("=====================================")

	fmt.Println("Take Action")
}
