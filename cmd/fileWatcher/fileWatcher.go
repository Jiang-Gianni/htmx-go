package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var (
	reload = make(chan interface{})
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Watch entire directory
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		err = watcher.Add(e.Name())
		if err != nil {
			log.Fatal(err)
		}
	}

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
				onHtmlUpdate(name)
			}

			if name[len(name)-3:] == "sql" {
				onSqlUpdate(name)
			}

			if name[len(name)-2:] == "go" {
				onGoUpdate()
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
	cmd := exec.Command("qtc", "-file="+fileName)
	RunCmd(cmd)
}

// Run sqlc generate
func onSqlUpdate(fileName string) {
	cmd := exec.Command("sqlc", "generate")
	RunCmd(cmd)
}

// Kill main if in execution, run main.go and signal to reload chan
func onGoUpdate() {
	cmd := exec.Command("pkill", "main")
	err := RunCmd(cmd)
	if err != nil {
		println(err.Error())
	}

	go func() {
		cmd = exec.Command("go", "run", "main.go")
		RunCmd(cmd)
	}()

	reload <- struct{}{}
}

// Send empty struct on reload signal
func handleWs(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer ws.Close()

	done := make(chan interface{})

	go func() {
		ws.ReadMessage()
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
