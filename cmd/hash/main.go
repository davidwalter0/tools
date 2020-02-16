package main

import (
	"crypto"
	"fmt"
	"github.com/davidwalter0/tools/hash"
)

func main() {
	{
		var Sha256EmptyString = hash.NewCryptoHashableStream(crypto.SHA256, []byte(""))
		// fmt.Println("", Sha256EmptyString, Sha256EmptyString.Sum().Hex())
		var equal = Sha256EmptyString.Sum().Hex() == hash.Sha256EmptyStringHashCompare
		fmt.Println(fmt.Sprintf("%v", equal),
			Sha256EmptyString.Sum().Hex(), "==",
			hash.Sha256EmptyStringHashCompare)
		hash.PrintSpace16(hash.Sha256EmptyStringHashCompare)
		hash.PrintSpace16(Sha256EmptyString.Sum().Hex())
	}
	{
		var Sha3256EmptyString = hash.NewCryptoHashableStream(crypto.SHA3_256, []byte(""))
		// fmt.Println("", Sha3256EmptyString, Sha3256EmptyString.Sum().Hex())
		var equal = Sha3256EmptyString.Sum().Hex() == hash.Sha3256EmptyStringHashCompare
		fmt.Println(fmt.Sprintf("%v", equal),
			Sha3256EmptyString.Sum().Hex(), "==",
			hash.Sha3256EmptyStringHashCompare)
		hash.PrintSpace16(hash.Sha3256EmptyStringHashCompare)
		hash.PrintSpace16(Sha3256EmptyString.Sum().Hex())
	}
	fmt.Println()
	{
		var hajime = hash.NewHajime(crypto.SHA256, "")
		// fmt.Println(hajime.Marshal())
		// fmt.Println(hajime.Marshal().Sum())
		fmt.Println(hajime.Marshal().Sum().Hex())
		fmt.Println(hajime.Marshal().Sum().Base16())
		fmt.Println(hajime.Marshal().Sum().Base58())
		fmt.Println()
		// fmt.Println(hajime)
	}
	fmt.Println()
	{
		var hajime = hash.NewHajime(crypto.SHA3_256, "")
		// fmt.Println(hajime.Marshal())
		// fmt.Println(hajime.Marshal().Sum())
		fmt.Println(hajime.Marshal().Sum().Hex())
		fmt.Println(hajime.Marshal().Sum().Base16())
		fmt.Println(hajime.Marshal().Sum().Base58())
		fmt.Println()
		// fmt.Println(hajime)
	}
	fmt.Println()

	msg := struct {
		Timestamp string
		Date      string
		Address   string
		Body      string
		Type      string
	}{
		Timestamp: "date",
		Date:      "readable_date",
		Address:   "address",
		Body:      "body",
		Type:      "1",
	}
	{
		var hajime = hash.NewHajime(crypto.SHA256, msg)
		// fmt.Println(hajime.Marshal())
		// fmt.Println(hajime.Marshal().Sum())
		fmt.Println(hajime.Marshal().Sum().Hex())
		fmt.Println(hajime.Marshal().Sum().Base16())
		fmt.Println(hajime.Marshal().Sum().Base58())
		fmt.Println()
		// fmt.Println(hajime)
	}
	{
		var hajime = hash.NewHajime(crypto.SHA3_256, msg)
		// fmt.Println(hajime.Marshal())
		// fmt.Println(hajime.Marshal().Sum())
		fmt.Println(hajime.Marshal().Sum().Hex())
		fmt.Println(hajime.Marshal().Sum().Base16())
		fmt.Println(hajime.Marshal().Sum().Base58())
		fmt.Println()
		// fmt.Println(hajime)
	}
}
