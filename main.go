//go:generate wire wire.go

package main

func main() {
	s := InitServer()
	s.ListenAndServe()
}
