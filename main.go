package main

import (
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const (
	numbers   = "0123456789"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols   = "!@#$%^&*():;',."
)

func main() {
	// Parse command-line arguments
	privateKey := flag.String("private-key", "", "User's private key")
	masterKey := flag.String("master-key", "", "User's master key")
	siteName := flag.String("site-name", "", "Site or app name")
	salt := flag.String("salt", "", "Additional tag used as \"salt\"")
	length := flag.Int("length", 16, "Hash length")
	charset := flag.String("charset", "alnum", "Character set to use for hash generation(num, lower, upper, alph, alnum, and all).")
	flag.Parse()

	// Validate command-line arguments
	if *privateKey == "" || *masterKey == "" || *siteName == "" || *salt == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Choose hash algorithm based on the selected hash length
	var h = sha512.New()

	// Concatenate the private key, master key, site name, and salt into a single string
	str := fmt.Sprintf("%s%s%s%s", *privateKey, *masterKey, *siteName, *salt)

	// Hash the concatenated string
	h.Write([]byte(str))
	hashBytes := h.Sum(nil)

	// Generate a hash of the desired length using the selected character set
	var chars string
	switch *charset {
	case "num":
		chars = numbers
	case "lower":
		chars = lowercase
	case "upper":
		chars = uppercase
	case "alph":
		chars = lowercase + uppercase
	case "alnum":
		chars = numbers + lowercase
	case "all":
		chars = numbers + lowercase + uppercase + symbols
	default:
		fmt.Println("Invalid character set. Allowed sets are num, lower, upper, alph, alnum, and all.")
		os.Exit(1)
	}

	// Convert the hash bytes to a hex string and use it as a seed for the random number generator
	seed, err := hex.DecodeString(fmt.Sprintf("%x", hashBytes))
	if err != nil {
		fmt.Println("Error decoding hash bytes:", err)
		os.Exit(1)
	}
	rng := randWithSeed(seed)

	// Generate a hash of the desired length using the selected character set and the deterministic RNG
	var sb strings.Builder
	for i := 0; i < *length; i++ {
		sb.WriteByte(chars[rng.Intn(len(chars))])
	}
	fmt.Println(sb.String())
}

// randWithSeed returns a deterministic random number generator seeded with the given bytes.
func randWithSeed(seed []byte) *rand.Rand {
	// Use the first 8 bytes of the seed to initialize the RNG
	var seedInt int64
	for _, b := range seed[:8] {
		seedInt = (seedInt << 8) + int64(b)
	}
	return rand.New(rand.NewSource(seedInt))
}
