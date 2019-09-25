# go-astilectron-from-js-example

This project provides a basic example of how [go-astilectron](https://github.com/asticode/astilectron) could be started
from an electron renderer process.

# Why?

go-astilectron is a great way to build a go app with a UI from scratch.
However, its practically impossible to integrate it into an existsing electron app.
This package shows that with just a few tweaks, go-astilectron can easily become a drop-in package suitable for any type of electron project.

# Setting up

If you don't already have it installed, install [npm](https://nodejs.org/en/).  
Then, simply `cd` into the project directory and run:

```bash
$ npm install
```

## Setting up Go

This package comes with the go-astilectron entry point already compiled (for Windows), so there's no need to have go installed to run it.
However, If you want to modify the go code, or run it on darwin / linux: you'll need to install [Golang](https://golang.org/dl/) and run:

```bash
$ go get -u github.com/asticode/astilectron
$ npm build
```

# Starting

When you're ready to try it out, just run:

```bash
$ npm start
```
