package main

import (
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
)

func main() {
	// Vulnerable Hash (SHAttered)
	h := sha1.New()
	h.Write([]byte("password"))
	
	// Broken Hash (Collision attacks)
	m := md5.New()
	m.Write([]byte("sensitive_data"))

	// Deprecated Public Key (Shor's Algorithm)
	key, _ := rsa.GenerateKey(nil, 2048)

	fmt.Println("This app is highly vulnerable to Harvest-Now-Decrypt-Later (HNDL) attacks.", key, h, m)
}
