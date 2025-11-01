package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/sha3"
	"os"
	"strings"
)

func main() {
	fmt.Println("======== Name + Hashing Program ========")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')

	// Remove newline characters
	input1 = strings.TrimSpace(input1)
	input2 = strings.TrimSpace(input2)

	// Run all hash algorithms
	proofMe("MD5", hashMD5(input1), hashMD5(input2))
	proofMe("SHA1", hashSHA1(input1), hashSHA1(input2))
	proofMe("SHA256", hashSHA256(input1), hashSHA256(input2))
	proofMe("SHA512", hashSHA512(input1), hashSHA512(input2))
	proofMe("SHA3-256", hashSHA3(input1), hashSHA3(input2))
}

func proofMe(name, out1, out2 string) {
	fmt.Printf("\nHash (%s):\nOutput A: %s\nOutput B: %s\n", name, out1, out2)
	if out1 == out2 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}
}

// Hash functions
func hashMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", hash)
}

func hashSHA1(text string) string {
	hash := sha1.Sum([]byte(text))
	return fmt.Sprintf("%x", hash)
}

func hashSHA256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", hash)
}

func hashSHA512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return fmt.Sprintf("%x", hash)
}

func hashSHA3(text string) string {
	hash := sha3.Sum256([]byte(text))
	return fmt.Sprintf("%x", hash)
}
