package main

import "github.com/stingalleman/go-thingy/netw"

func main() {
	l, err := netw.Server("[::]", 8998)
	if err != nil {
		panic(err)
	}

	for l.Accept() {}
}
