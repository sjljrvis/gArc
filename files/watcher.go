package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/sjljrvis/gArch/helpers"
)

// DirWatcher  *is directory observer to watch file-events in directory
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
	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()
	done := make(chan bool)

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
	switch op := event.Op.String(); op {
	case "CREATE":
		log.Println("File created -", event.Name)
		handleCreate(event.Name)

	default:
		log.Println("->", op)
	}
}

func handleCreate(path string) {

	_file, err := os.Open(path)
	defer _file.Close()

	if err != nil {
		log.Fatalln(err)
	}

	fileInfo, _ := _file.Stat()
	log.Println(fileInfo.Size())
	var fileSize = fileInfo.Size()
	const fileChunk = 1 * (1 << 20)
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		_file.Read(partBuffer)

		fileName := os.Getenv("HOME") + "/gArch/chunks/chunk_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		fmt.Println("Split to : ", fileName)
	}
}
