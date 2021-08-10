package main

import (
  "fmt"
  "github.com/tyler-smith/go-bip39"
	"encoding/hex"
  "os"
	// "reflect"
  // "github.com/tyler-smith/go-bip32"
)

func main(){
  // Generate a mnemonic for memorization or user-friendly seeds
  // entropy, _ := bip39.NewEntropy(256)
  // hexCode0 := os.Getenv("HEX_ENT0")
  // hexCode1 := os.Getenv("HEX_ENT1")
  hexCodeLong := os.Getenv("HEX_ENT_LONG")

	entropy, _ := hex.DecodeString(hexCodeLong)
  mnemonic, _ := bip39.NewMnemonic(entropy)
	// hexCode := hex.EncodeToString(entropy)
	fmt.Println(mnemonic)


	// fmt.Println(entropy)
	// fmt.Println(dataStr)

	// fmt.Printf("%T", entropy)
	// fmt.Println(entropy)
	// fmt.Println(hex.EncodeToString(entropy))

  // // Generate a Bip32 HD wallet for the mnemonic and a user supplied password
  // seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

  // masterKey, _ := bip32.NewMasterKey(seed)
  // publicKey := masterKey.PublicKey()

  // // Display mnemonic and keys
  // fmt.Println("Mnemonic: ", mnemonic)
  // fmt.Println("Master private key: ", masterKey)
  // fmt.Println("Master public key: ", publicKey)
}