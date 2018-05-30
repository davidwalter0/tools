package hash // import "github.com/davidwalter0/tools/hash"

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"

	"github.com/btcsuite/btcutil/base58"
)

const (
	Sha256aHash                   = "ba7816bf 8f01cfea 414140de 5dae2223 b00361a3 96177a9c b410ff61 f20015ad"
	Sha256aHashCompare            = "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	Sha3256aHash                  = "3a985da74fe225b2 045c172d6bd390bd 855f086e3e9d525b 46bfe24511431532"
	Sha3256aHashCompare           = "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532"
	Sha256EmptyStringHash         = "e3b0c442 98fc1c14 9afbf4c8 996fb924 27ae41e4 649b934c a495991b 7852b855"
	Sha3256EmptyStringHash        = "a7ffc6f8bf1ed766 51c14756a061d662 f580ff4de43b49fa 82d80a4b80f8434a"
	Sha256EmptyStringHashCompare  = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	Sha3256EmptyStringHashCompare = "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a"
)

func PrintSpace16(hex string) {
	for i := 0; i < len(hex); i += 16 {
		var slice = hex[i:]
		if len(slice) > 16 {
			fmt.Printf("%s ", string(slice[:16]))
		} else {
			fmt.Printf("%s \n", string(slice))
		}
	}
}

func Base32Encode(in []byte) (out string) {
	out = base32.StdEncoding.EncodeToString(in)
	return
}
func Base64Encode(in []byte) (out string) {
	out = base64.StdEncoding.EncodeToString(in)
	return
}

func Base32Decode(in string) (out []byte) {
	//	var err error
	// out, err = base32.StdEncoding.DecodeString(in)
	out, _ = base32.StdEncoding.DecodeString(in)
	//	CheckError(err)
	return
}

func Base64Decode(in string) (out []byte) {
	//	var err error
	// out, err = base64.StdEncoding.DecodeString(in)
	out, _ = base64.StdEncoding.DecodeString(in)
	//	CheckError(err)
	return
}

// Base58Encode
func Base58Encode(data []byte) string {
	return base58.Encode(data)
}

// Base58Decode
func Base58Decode(data string) []byte {
	return base58.Decode(data)
}

// Sum text crypto hash with specified Hash or default sha3.Sum256
func Sum(text []byte, Hash ...crypto.Hash) (hash []byte) {
	var cryptoHash = crypto.SHA3_256
	if len(Hash) > 0 {
		cryptoHash = Hash[0]
	}
	switch cryptoHash {
	case crypto.MD5:
		sum := md5.Sum(text)
		hash = sum[:]
	case crypto.SHA256:
		sum := sha256.Sum256(text)
		hash = sum[:]
	case crypto.SHA3_256:
		sum := sha3.Sum256(text)
		hash = sum[:]
	default:
		sum := sha3.Sum256(text)
		hash = sum[:]
	}
	return hash
}

// Hex hex formatted string
func Hex(text []byte) string {
	hexified := make([][]byte, len(text))
	for i, data := range text {
		hexified[i] = []byte(fmt.Sprintf("%02x", data))
	}
	text = bytes.Join(hexified, []byte(""))
	return string(text)
}

// Hex hex formatted string
func HexDecode(in string) (out []byte) {
	out = make([]byte, len(in)/2)
	if _, err := hex.Decode(out, []byte(in)); err != nil {
		panic(err)
	}
	return
}
