// command main starts and configures go-astilectron for us.
package main

import (
	"fmt"
	"os"

	"github.com/sbrow/go-astilectron"
)

func main() {
	var port string

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Initialize astilectron
	var a, err = astilectron.New(astilectron.Options{
		SkipAstilectronSetup: true,
		TCPPort:              port,
	})
	if err != nil {
		panic(err)
	}
	defer a.Close()

	// Start the TCP server for go-astilectron
	fmt.Printf("Starting go-astilectron on port %s...\n", port)
	a.Start()
	// Wait until electron is ready before creating a new window.
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
	// Wai until we get the kill signal from astilectron.
	a.Wait()
	fmt.Println("...Done!")
}
