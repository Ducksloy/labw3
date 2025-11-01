package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"exercise2/utils/crack"
)

func main() {
    wordlist := flag.String("wordlist", "nord_vpn.txt", "path to wordlist file")
    target := flag.String("hash", "aa1c7d931cf140bb35a5a16adeb83a551649c3b9", "target SHA1 hash to crack")
    out := flag.String("out", "verbose.txt", "output verbose log file")
    flag.Parse()

    f, err := os.Open(*wordlist)
    if err != nil {
        log.Fatalf("failed to open wordlist %q: %v", *wordlist, err)
    }
    defer f.Close()

    outFile, err := os.Create(*out)
    if err != nil {
        log.Fatalf("failed to create output file %q: %v", *out, err)
    }
    defer outFile.Close()

    writer := bufio.NewWriter(outFile)
    defer writer.Flush()

    scanner := bufio.NewScanner(f)
    lineNo := 0
    found := false
    t := strings.ToLower(*target)
    for scanner.Scan() {
        lineNo++
        word := strings.TrimSpace(scanner.Text())
        if word == "" {
            continue
        }
        sum := crack.SHA1Hex(word)
        fmt.Fprintf(writer, "line %d: trying %q -> %s\n", lineNo, word, sum)
        if lineNo%1000 == 0 {
            fmt.Printf("processed %d lines, last tried %q -> %s\n", lineNo, word, sum)
        }
        if sum == t {
            fmt.Printf("FOUND at line %d: %s\n", lineNo, word)
            fmt.Fprintf(writer, "FOUND at line %d: %s -> %s\n", lineNo, word, sum)
            found = true
            break
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatalf("error reading wordlist: %v", err)
    }
    if !found {
        fmt.Println("Hash not found in wordlist.")
        fmt.Fprintln(writer, "Hash not found in wordlist.")
    }
}
