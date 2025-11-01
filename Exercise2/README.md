# Lab #1 — Password Cracking (exercise2)

This folder contains a Go program to attempt to crack a SHA1 hash using a supplied wordlist.

Files:

- `main.go` — program that reads a wordlist (default `nord_vpn.txt`), computes SHA1 for each entry, compares to target hash, and writes verbose output to `verbose.txt`.
- `utils/crack/sha1.go` — helper that computes the SHA1 hex string (calls `crypto/sha1`).
- `nord_vpn.txt` — (should be present) wordlist used by the lab.

Quick run (from this folder):

```powershell
cd "d:\Learning\Year 3\Term 1\Cryptography\Lab\Lab3\exercise2"
go run .
```

Defaults:

- target hash: `aa1c7d931cf140bb35a5a16adeb83a551649c3b9` (from the lab slide)
- wordlist: `nord_vpn.txt`
- verbose output: `verbose.txt` (contains per-line SHA1s and a FOUND line if cracked)

You can override defaults with flags:

```powershell
go run . -wordlist=path\to\list.txt -hash=<sha1hash> -out=verbose.txt
```
