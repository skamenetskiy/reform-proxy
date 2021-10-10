# reform-proxy
A simple (embedded) database proxy for primary/secondary splitting using 
[reform orm](https://github.com/go-reform/reform).

The secondary nodes are selected using round-robin algorithm, inspired by hlts2's 
[implementation](https://github.com/hlts2/round-robin) or round-robin.

[![Go Report Card](https://goreportcard.com/badge/github.com/skamenetskiy/reform-proxy)](https://goreportcard.com/report/github.com/skamenetskiy/reform-proxy)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/skamenetskiy/reform-proxy)](https://pkg.go.dev/github.com/skamenetskiy/reform-proxy)