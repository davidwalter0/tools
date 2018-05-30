package main

import (
	"crypto"
	"fmt"
	. "github.com/davidwalter0/tools/hash"

	"github.com/davidwalter0/go-signalstor"
)

func main() {
	{
		var Sha256EmptyString = NewCryptoHashableStream(crypto.SHA256, []byte(""))
		// fmt.Println("", Sha256EmptyString, Sha256EmptyString.Sum().Hex())
		var equal = Sha256EmptyString.Sum().Hex() == Sha256EmptyStringHashCompare
		fmt.Println(fmt.Sprintf("%v", equal),
			Sha256EmptyString.Sum().Hex(), "==",
			Sha256EmptyStringHashCompare)
		PrintSpace16(Sha256EmptyStringHashCompare)
		PrintSpace16(Sha256EmptyString.Sum().Hex())
	}
	{
		var Sha3256EmptyString = NewCryptoHashableStream(crypto.SHA3_256, []byte(""))
		// fmt.Println("", Sha3256EmptyString, Sha3256EmptyString.Sum().Hex())
		var equal = Sha3256EmptyString.Sum().Hex() == Sha3256EmptyStringHashCompare
		fmt.Println(fmt.Sprintf("%v", equal),
			Sha3256EmptyString.Sum().Hex(), "==",
			Sha3256EmptyStringHashCompare)
		PrintSpace16(Sha3256EmptyStringHashCompare)
		PrintSpace16(Sha3256EmptyString.Sum().Hex())
	}
	fmt.Println()
	{
		var hajime = NewHajime(crypto.SHA256, "")
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
		var hajime = NewHajime(crypto.SHA3_256, "")
		// fmt.Println(hajime.Marshal())
		// fmt.Println(hajime.Marshal().Sum())
		fmt.Println(hajime.Marshal().Sum().Hex())
		fmt.Println(hajime.Marshal().Sum().Base16())
		fmt.Println(hajime.Marshal().Sum().Base58())
		fmt.Println()
		// fmt.Println(hajime)
	}
	fmt.Println()
	msg := signalstor.SmsMessage{
		Timestamp: "date",
		Date:      "readable_date",
		Address:   "address",
		Body:      "body",
		Type:      "1",
	}
	{
		var hajime = NewHajime(crypto.SHA256, msg)
		// fmt.Println(hajime.Marshal())
		// fmt.Println(hajime.Marshal().Sum())
		fmt.Println(hajime.Marshal().Sum().Hex())
		fmt.Println(hajime.Marshal().Sum().Base16())
		fmt.Println(hajime.Marshal().Sum().Base58())
		fmt.Println()
		// fmt.Println(hajime)
	}
	{
		var hajime = NewHajime(crypto.SHA3_256, msg)
		// fmt.Println(hajime.Marshal())
		// fmt.Println(hajime.Marshal().Sum())
		fmt.Println(hajime.Marshal().Sum().Hex())
		fmt.Println(hajime.Marshal().Sum().Base16())
		fmt.Println(hajime.Marshal().Sum().Base58())
		fmt.Println()
		// fmt.Println(hajime)
	}
}
