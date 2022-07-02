package signature

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

func TestPersonalSign(t *testing.T) {
	// Load private key from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyString := os.Getenv("PRIVATE_KEY")
	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}
	// Sign message
	signature, err := PersonalSign("Hello World.", privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signature)
}
