# AnubisCrypt
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/anubiscrypt/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/anubiscrypt?status.png)](http://godoc.org/github.com/pedroalbanese/anubiscrypt)
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/anubiscrypt/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/anubiscrypt/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/anubiscrypt)](https://goreportcard.com/report/github.com/pedroalbanese/anubiscrypt)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/anubiscrypt)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/anubiscrypt)](https://github.com/pedroalbanese/anubiscrypt/releases)  

Barreto & Rijmen 128-bit block cipher with GCM Mode (RFC 5288) provides both authenticated encryption (confidentiality and authentication) and the ability to check the integrity and authentication of additional authenticated data (AAD) that is sent in the clear. Whirlpool-based PBKDF2. Anubis is a block cipher with SP-network structure designed by Vincent Rijmen and Paulo S. L. M. Barreto in 2000.

### Command-line Anubis-GCM Encryption Tool
<pre>Usage of anubiscrypt:
anubiscrypt [-d] -p "pass" [-i N] [-s "salt"] -f &lt;file.ext&gt;
  -a string
        Additional Associated data.
  -d    Decrypt instead of Encrypt.
  -f string
        Target file. ('-' for STDIN)
  -i int
        Iterations. (for PBKDF2) (default 1)
  -k string
        Symmetric key to Encrypt/Decrypt.
  -p string
        Password-based key derivation function 2.
  -r    Generate random cryptographic key with 128-bit.
  -s string
        Salt. (for PBKDF2)</pre>

#### Example:
```sh
./gostcrypt -k "" -f plaintext.ext > ciphertext.ext
./gostcrypt -d -k $256bitkey -f ciphertext.ext > plaintext.ext
```

## License

This project is licensed under the ISC License.

##### Industrial-Grade Reliability. Copyright (c) 2020-2022 ALBANESE Research Lab.
