# Lab #1 — Password Cracking (exercise1)

This folder contains a small Go program to attempt to crack an MD5 hash using a supplied wordlist.

Files:

- `main.go` — program that reads a wordlist (default `nord_vpn.txt`), computes MD5 for each entry, compares to target hash, and writes verbose output to `verbose.txt`.
- `utils/crack/md5.go` — helper that computes the MD5 hex string (calls `crypto/md5`).
- `nord_vpn.txt` — (already present) wordlist used by the lab.

Quick run (from this folder):

```powershell
cd "d:\Learning\Year 3\Term 1\Cryptography\Lab\Lab3\exercise1"
go run .
```

Defaults:

- target hash: `6a85dfd77d9cb35770c9dc6728d73d3f` (from the lab slide)
- wordlist: `nord_vpn.txt`
- verbose output: `verbose.txt` (contains per-line MD5s and a FOUND line if cracked)

You can override defaults with flags:

```powershell
go run . -wordlist=path\to\list.txt -hash=<md5hash> -out=verbose.txt
```
