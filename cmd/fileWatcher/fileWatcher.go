package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var (
	reload  = make(chan interface{})
	clients = map[*websocket.Conn]struct{}{}
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Watch entire directory and subdirectories
	addEntries := func(string) {}
	addEntries = func(name string) {
		entries, err := os.ReadDir(name)
		if err != nil {
			log.Fatal(err)
		}
		for _, e := range entries {
			if e.IsDir() && e.Name() != ".git" && e.Name() != "cmd" {
				var entryName string
				if name == "./" {
					entryName = "./" + e.Name()
				} else {
					entryName = name + "/" + e.Name()
				}
				err = watcher.Add(entryName)
				if err != nil {

					log.Fatal(err)
				}
				fmt.Println(entryName)
				addEntries(entryName)
			}
		}
	}
	addEntries("./")

	http.HandleFunc("/ws", handleWs)
	go func() {
		if err := http.ListenAndServe(":1234", nil); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("watching for file changes")
	for {
		select {

		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			name := event.Name

			if name[len(name)-4:] == "html" {
				onScssUpdate()
				onHtmlUpdate(name)
			}

			if name[len(name)-3:] == "sql" {
				onSqlUpdate()
			}

			if name[len(name)-2:] == "go" {
				onGoUpdate()
			}

			if name[len(name)-4:] == "scss" {
				onScssUpdate()
				reloadSignal()
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}

}

// Compile modified html
func onHtmlUpdate(fileName string) {
	trackTime(time.Now(), "qtc")
	cmd := exec.Command("qtc", "-file="+fileName)
	RunCmd(cmd)
}

// Run sqlc generate
func onSqlUpdate() {
	trackTime(time.Now(), "sqlc")
	cmd := exec.Command("sqlc", "generate")
	RunCmd(cmd)
}

// Run sass compiler
func onScssUpdate() {
	trackTime(time.Now(), "scss + purgecss")
	cmd := exec.Command("sass", "--style=compressed", "style/main.scss", "assets/style.css")
	RunCmd(cmd)

	cmd = exec.Command("purgecss", "-css", "assets/style.css", "--content", "views/*.html", "--output", "assets/")
	RunCmd(cmd)
}

// Kill main if in execution, run main.go and signal to reload chan
func onGoUpdate() {
	trackTime(time.Now(), "go")
	cmd := exec.Command("pkill", "main")
	err := RunCmd(cmd)
	if err != nil {
		println(err.Error())
	}

	go func() {
		cmd = exec.Command("go", "run", "main.go")
		RunCmd(cmd)
	}()

	reloadSignal()
}

// Send empty struct on reload signal
func handleWs(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	clients[ws] = struct{}{}
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer ws.Close()

	done := make(chan interface{})

	go func() {
		ws.ReadMessage()
		delete(clients, ws)
		done <- struct{}{}
	}()
	for {
		select {
		case <-done:
			return
		case <-reload:
			println("reload")
			ws.WriteJSON(struct{}{})
		}
	}
}

func RunCmd(cmd *exec.Cmd) error {
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func reloadSignal() {
	if len(clients) > 0 {
		reload <- struct{}{}
	}
}

func trackTime(start time.Time, update string) {
	fmt.Println(update, " took ", time.Since(start))
}
