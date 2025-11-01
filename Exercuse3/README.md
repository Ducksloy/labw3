# Lab #1 — Password Cracking (exercise3)

This folder contains a Go program to attempt to crack a SHA-512 hash using a supplied wordlist.

Files:

- `main.go` — program that reads a wordlist (default `nord_vpn.txt`), computes SHA-512 for each entry, compares to target hash, and writes verbose output to `verbose.txt`.
- `utils/crack/sha512.go` — helper that computes the SHA-512 hex string (calls `crypto/sha512`).
- `nord_vpn.txt` — (should be present) wordlist used by the lab.

Quick run (from this folder):

```powershell
cd "d:\Learning\Year 3\Term 1\Cryptography\Lab\Lab3\exercise3"
go run .
```

Defaults:

- target hash: `485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38` (from the lab slide)
- wordlist: `nord_vpn.txt`
- verbose output: `verbose.txt` (contains per-line SHA-512s and a FOUND line if cracked)

You can override defaults with flags:

```powershell
go run . -wordlist=path\to\list.txt -hash=<sha512hash> -out=verbose.txt
```
