package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"log"
	"os"

	"github.com/dgryski/go-anubis"
	"github.com/pedroalbanese/cmac"
	"github.com/pedroalbanese/edgetk/eax"
	"github.com/pedroalbanese/whirlpool"
)

var (
	aad    = flag.String("a", "", "Additional Associated data.")
	dec    = flag.Bool("d", false, "Decrypt instead Encrypt.")
	file   = flag.String("f", "", "Target file. ('-' for STDIN)")
	iter   = flag.Int("i", 1, "Iterations. (for PBKDF2)")
	key    = flag.String("k", "", "Symmetric key to Encrypt/Decrypt.")
	length = flag.Int("b", 256, "Key length: 128, 192 or 256.")
	mac    = flag.Bool("m", false, "Cipher-based message authentication code.")
	pbkdf  = flag.String("p", "", "Password-based key derivation function 2.")
	random = flag.Bool("r", false, "Generate random cryptographic key with given bit-length.")
	salt   = flag.String("s", "", "Salt. (for PBKDF2)")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "AnubisCrypt - ALBANESE Research Lab (c) 2020-2022 ")
		fmt.Fprintln(os.Stderr, "Barreto & Rijmen 128-bit block cipher in EAX Mode\n")
		fmt.Fprintln(os.Stderr, "Usage of "+os.Args[0]+":")
		fmt.Fprintln(os.Stderr, os.Args[0]+" [-d] [-b N] -p \"pass\" [-i N] [-s \"salt\"] -f <file.ext>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *random == true {
		var key []byte
		var err error
		key = make([]byte, *length/8)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hex.EncodeToString(key))
		os.Exit(0)
	}

	var keyHex string
	var keyRaw []byte

	if *mac {
		if *pbkdf != "" {
			keyRaw = pbkdf2.Key([]byte(*pbkdf), []byte(*salt), *iter, 16, whirlpool.New)
			keyHex = hex.EncodeToString(keyRaw)
		} else {
			keyHex = *key
		}
		var err error
		var ciph cipher.Block
		ciph = anubis.New([]byte(keyHex))
		if err != nil {
			log.Fatal(err)
		}
		h, _ := cmac.New(ciph)
		var data io.Reader
		if *file == "-" {
			data = os.Stdin
		} else {
			data, _ = os.Open(*file)
		}
		io.Copy(h, data)
		fmt.Println(hex.EncodeToString(h.Sum(nil)))
		os.Exit(0)
	} else {
		if *pbkdf != "" {
			keyRaw = pbkdf2.Key([]byte(*pbkdf), []byte(*salt), *iter, *length/8, whirlpool.New)
			keyHex = hex.EncodeToString(keyRaw)
		} else {
			keyHex = *key
		}
		var key []byte
		var err error
		if keyHex == "" {
			key = make([]byte, *length/8)
			_, err = io.ReadFull(rand.Reader, key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(os.Stderr, "Key=", hex.EncodeToString(key))
		} else {
			key, err = hex.DecodeString(keyHex)
			if err != nil {
				log.Fatal(err)
			}
			if len(key) != 32 && len(key) != 24 && len(key) != 16 {
				log.Fatal(err)
			}
		}

		buf := bytes.NewBuffer(nil)
		var data io.Reader
		if *file == "-" {
			data = os.Stdin
		} else {
			data, _ = os.Open(*file)
		}
		io.Copy(buf, data)
		msg := buf.Bytes()

		var c cipher.Block
		c = anubis.New(key)
		aead, _ := eax.NewEAX(c)

		if *dec == false {
			nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())

			out := aead.Seal(nonce, nonce, msg, []byte(*aad))
			fmt.Printf("%s", out)

			os.Exit(0)
		}

		if *dec == true {
			nonce, msg := msg[:aead.NonceSize()], msg[aead.NonceSize():]

			out, err := aead.Open(nil, nonce, msg, []byte(*aad))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", out)

			os.Exit(0)
		}
	}
}
