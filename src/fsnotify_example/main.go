package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan struct{})
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Printf("%s %s", event.Op, event.Name)
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Print(event.Op, event.Name)
				}
			case err := <-watcher.Errors:
				log.Println(err)
			}
		}
	}()

	err = watcher.Add("./tmp/")
	if err != nil {
		log.Fatalln(err)
	}

	<-done
}
