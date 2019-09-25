# go-astilectron-from-js-example

This project provides a basic example of how to start [go-astilectron](https://github.com/asticode/astilectron) 
from the electron renderer process.

# Why?

go-astilectron is a great way to build a go app with a UI from scratch. It also a great way to integrate a Golang backend into an existing electron project! This repo demonstrates how to set that up.

# Setting up

If you don't already have it installed, install [npm](https://nodejs.org/en/).  
Then, simply `cd` into the project directory and run:

```bash
$ npm install
```

## Setting up Go

This package comes with the go-astilectron entry point already compiled (for Windows), so there's no need to have go installed to run it.
However, In order to run it on darwin / linux, or to modify the go code: you'll need to install [Golang](https://golang.org/dl/) and run:

```bash
$ go get -u github.com/asticode/astilectron
$ npm build
```

# Starting

When you're ready to try it out, just run:

```bash
$ npm start
```
