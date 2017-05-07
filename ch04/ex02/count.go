package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

const (
	SHA256  = "SHA256"
	SHA384  = "SHA384"
	SHA512  = "SHA512"
	message = "choose hash algorithm. SHA256, SHA384, or SHA512"
)

var hash string

func init() {
	flag.StringVar(&hash, "hash", SHA256, message)
}

func main() {
	flag.Parse()
	for _, v := range flag.Args() {
		switch hash {
		case SHA256:
			fmt.Printf("SHA256: %s -> %x\n", v, sha256.Sum256([]byte(v)))
		case SHA384:
			fmt.Printf("SHA384: %s -> %x\n", v, sha512.Sum384([]byte(v)))
		case SHA512:
			fmt.Printf("SHA512: %s -> %x\n", v, sha512.Sum512([]byte(v)))
		default:
			fmt.Println(message)
		}
	}
}
