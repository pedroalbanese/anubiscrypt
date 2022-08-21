# AnubisCrypt
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/anubiscrypt/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/anubiscrypt?status.png)](http://godoc.org/github.com/pedroalbanese/anubiscrypt)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/anubiscrypt)](https://goreportcard.com/report/github.com/pedroalbanese/anubiscrypt)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/anubiscrypt)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/anubiscrypt)](https://github.com/pedroalbanese/anubiscrypt/releases)  

Barreto & Rijmen 128-bit block cipher in EAX Mode provides both authenticated encryption (confidentiality and authentication) and the ability to check the integrity and authentication of additional authenticated data (AAD) that is sent in the clear..
### Command-line Anubis-EAX Encryption Tool
<pre>anubiscrypt [-d] [-b N] -p "pass" [-i N] [-s "salt"] -f <file.ext>
  -b int
        Key length: 128, 192 or 256. (default 256)
  -d    Decrypt instead Encrypt.
  -f string
        Target file. ('-' for STDIN)
  -i int
        Iterations. (for PBKDF2) (default 1024)
  -k string
        Symmetric key to Encrypt/Decrypt.
  -m    Cipher-based message authentication code.
  -p string
        Password-based key derivation function 2.
  -r    Generate random cryptographic key with given bit-length.
  -s string
        Salt. (for PBKDF2)</pre>

## License

This project is licensed under the ISC License.

##### Industrial-Grade Reliability. Copyright (c) 2020-2022 ALBANESE Research Lab.
