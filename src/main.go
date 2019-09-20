// command main starts and configures go-astilectron for us.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/asticode/go-astilectron"
)

type ConnectionInfo struct {
	Addr string `json:"addr"`
}

func main() {
	port := new(int)

	if len(os.Args) > 1 {
		p, err := strconv.Atoi(os.Args[1])
		if err != nil {
			println(os.Args)
			panic(err)
		}
		*port = p
	}

	// Initialize astilectron
	var a, err = astilectron.New(astilectron.Options{SkipSetup: true, TCPPort: port})
	if err != nil {
		panic(err)
	}
	defer a.Close()

	fmt.Printf("Starting on port \"%s\"...\n", strconv.Itoa(*port))
	// Wait until electron is ready before creating a new window.
	a.Start()

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
	// Wait until we get the kill signal from astilectron.
	a.Wait()
	fmt.Println("...Done!")
}
