package routines

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/sjljrvis/gArch/helpers"
)

func DirWatcher() {
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
	log.Println(helpers.ListFiles(os.Getenv("HOME") + "/gArch"))
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
	log.Println(event.Name, event.Op)
}