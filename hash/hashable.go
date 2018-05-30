package hash // import "github.com/davidwalter0/tools/hash"

import (
	"crypto"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	//	"github.com/btcsuite/btcutil/base58"
)

const (
	TypeLen = 4
)

// CryptoHashableStream prefix a summable with the crypto.Hash
type CryptoHashableStream []byte

// UnhashedStream is a byte array of Unhashed text
type UnhashedStream []byte

func HashTypeToByte(Hash crypto.Hash) (buffer CryptoHashableStream) {
	buffer = make([]byte, TypeLen)
	binary.BigEndian.PutUint32(buffer, uint32(Hash))
	return
}

// CryptHashStrm2HashAndByte
func CryptHashStrm2HashAndByte(buffer CryptoHashableStream) (crypto.Hash, UnhashedStream) {
	if len(buffer) < TypeLen {
		panic(fmt.Sprintf("Buffer too short for CryptoHashableStream min length %d", TypeLen))
	}
	return crypto.Hash(binary.BigEndian.Uint32(buffer[:4])), UnhashedStream(buffer[4:])
}

func NewCryptoHashableStream(cryptoHash crypto.Hash, unhashed UnhashedStream) CryptoHashableStream {
	return append(HashTypeToByte(cryptoHash), unhashed...)
}

// Sum of CryptoHashableStream []byte
func (s CryptoHashableStream) Sum() HashSum {
	return HashSum(SumUnhashedStream(CryptHashStrm2HashAndByte(s)))
}

// HashSum has multiple base string output types, hex, base58
type HashSum []byte

func NewHashSum(cryptoHash crypto.Hash, buffer UnhashedStream) HashSum {
	return HashSum(SumUnhashedStream(cryptoHash, buffer))
}

func SumUnhashedStream(cryptoHash crypto.Hash, buffer UnhashedStream) HashSum {
	return Sum(buffer, cryptoHash)
}

// Hex from HashSum to formatted string
func (hs HashSum) Hex() string { return hs.Base16() }

// Base16 from HashSum to formatted string
func (hs HashSum) Base16() string { return Hex([]byte(hs)) }

// Base58 from HashSum to formatted string
func (hs HashSum) Base58() string {
	// fmt.Println("HEXHEXHEX", Hex(hs))
	return Base58Encode([]byte(hs))
}

// func (hs HashSum) Base58() string { return Base58Encode([]byte(hs)) }

type Hashable interface {
	Marshal() CryptoHashableStream
	Sum() HashSum
	Hex() string
	Base16() string
	Base58() string
	String() string
}

type Hajime struct {
	Hashable interface{}
	crypto.Hash
}

// func NewHajime(T interface{}, Hash ...crypto.Hash) *Hajime {
func NewHajime(Hash crypto.Hash, T interface{}) *Hajime {
	var cryptoHash = crypto.SHA3_256
	cryptoHash = Hash

	return &Hajime{Hashable: T, Hash: cryptoHash}
}

func (hj *Hajime) Marshal() CryptoHashableStream {
	switch hj.Hashable.(type) {
	case []byte:
		return NewCryptoHashableStream(hj.Hash, hj.Hashable.([]byte))
	case UnhashedStream:
		return NewCryptoHashableStream(hj.Hash, hj.Hashable.([]byte))
	case string:
		return NewCryptoHashableStream(hj.Hash, []byte(hj.Hashable.(string)))
	default:
		var err error
		var text = []byte{}
		if text, err = json.Marshal(hj.Hashable); err != nil {
			log.Fatal(err)
		}
		return NewCryptoHashableStream(hj.Hash, text)
	}
}
func (hj *Hajime) Sum() HashSum {
	return hj.Marshal().Sum()
}

func (hj *Hajime) Base16() string {
	return hj.Marshal().Sum().Base16()
}

func (hj *Hajime) Hex() string {
	return hj.Base16()
}

func (hj *Hajime) Base58() string {
	return hj.Marshal().Sum().Base58()
}

func (hj *Hajime) String() string {
	return fmt.Sprintf("Hash        : %s\nBase58      : %s\nMessage     : %s\n",
		hj.Hex(), hj.Sum(), string(hj.Marshal()),
	)
}
