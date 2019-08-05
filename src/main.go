package main

import (
	"encoding/json"
	"fmt"

	"github.com/sbrow/go-astilectron"
)

type ConnectionInfo struct {
	Addr string `json:"addr"`
}

func main() {
	// Initialize astilectron
	var a, err = astilectron.New(astilectron.Options{})
	if err != nil {
		panic(err)
	}
	defer a.Close()

	// Start astilectron
	sock, err := a.Listen()
	if err != nil {
		panic(err)
	} else {
		info := ConnectionInfo{Addr: sock.Addr().String()}
		if data, err := json.Marshal(info); err != nil {
			panic(err)
		} else {
			fmt.Println(string(data))
		}
	}
	a.WaitOn("app.event.ready")

	// Create a new window
	w, err := a.NewWindow("app/index.html", &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(600),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Creating window.")
	if err = w.Create(); err != nil {
		panic(err)
	}

	fmt.Println("Waiting...")
	a.Wait()
	fmt.Println("...Done!")
}
