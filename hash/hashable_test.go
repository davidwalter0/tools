package hash_test // import "github.com/davidwalter0/tools/hash"

import (
	"crypto"
	. "github.com/davidwalter0/tools/hash"

	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Crypto_256_Base58(t *testing.T) {
	var sha256EmptyString = NewHashSum(crypto.SHA256, []byte(""))
	assert.Equal(t, sha256EmptyString.Base58(), Base58Encode(HexDecode(sha256EmptyStringHashCompare)), "comparing known hash as base58 encoding failed")
	var sha256_abc_Hash = NewHashSum(crypto.SHA256, []byte("abc"))
	assert.Equal(t, sha256_abc_Hash.Base58(), Base58Encode(HexDecode(sha256_abc_HashCompare)), "comparing known hash as base58 encoding failed")
}

func Test_Crypto3_256_Base58(t *testing.T) {
	var sha3_256EmptyString = NewHashSum(crypto.SHA3_256, []byte(""))
	assert.Equal(t, sha3_256EmptyString.Base58(), Base58Encode(HexDecode(sha3_256EmptyStringHashCompare)), "comparing known hash as base58 encoding failed")
	var sha3_256_abc_Hash = NewHashSum(crypto.SHA3_256, []byte("abc"))
	assert.Equal(t, sha3_256_abc_Hash.Base58(), Base58Encode(HexDecode(sha3_256_abc_HashCompare)), "comparing known hash as base58 encoding failed")
}

func Test_Crypto_256_Hajime(t *testing.T) {
	var sha256EmptyString = NewHajime(crypto.SHA256, []byte(""))
	assert.Equal(t, sha256EmptyString.Sum().Hex(), sha256EmptyStringHashCompare, "comparing known hash values failed")
	var sha256_abc_Hash = NewHajime(crypto.SHA256, []byte("abc"))
	assert.Equal(t, sha256_abc_Hash.Sum().Hex(), sha256_abc_HashCompare, "comparing known hash values failed")
}

func Test_Crypto3_256_Hajime(t *testing.T) {
	var sha3_256EmptyString = NewHajime(crypto.SHA3_256, []byte(""))
	assert.Equal(t, sha3_256EmptyString.Sum().Hex(), sha3_256EmptyStringHashCompare, "comparing known hash values failed")
	var sha3_256_abc_Hash = NewHajime(crypto.SHA3_256, []byte("abc"))
	assert.Equal(t, sha3_256_abc_Hash.Sum().Hex(), sha3_256_abc_HashCompare, "comparing known hash values failed")
}

func Test_Crypto_256_HashFactory(t *testing.T) {
	var sha256EmptyString = NewCryptoHashableStream(crypto.SHA256, []byte(""))
	assert.Equal(t, sha256EmptyString.Sum().Hex(), sha256EmptyStringHashCompare, "comparing known hash values failed")
	var sha256_abc_Hash = NewCryptoHashableStream(crypto.SHA256, []byte("abc"))
	assert.Equal(t, sha256_abc_Hash.Sum().Hex(), sha256_abc_HashCompare, "comparing known hash values failed")
}

func Test_Crypto3_256_HashFactory(t *testing.T) {
	var sha3_256EmptyString = NewCryptoHashableStream(crypto.SHA3_256, []byte(""))
	assert.Equal(t, sha3_256EmptyString.Sum().Hex(), sha3_256EmptyStringHashCompare, "comparing known hash values failed")
	var sha3_256_abc_Hash = NewCryptoHashableStream(crypto.SHA3_256, []byte("abc"))
	assert.Equal(t, sha3_256_abc_Hash.Sum().Hex(), sha3_256_abc_HashCompare, "comparing known hash values failed")
}

func Test_Crypto_256HashSum(t *testing.T) {
	var sha256EmptyString = NewHashSum(crypto.SHA256, []byte(""))
	assert.Equal(t, sha256EmptyString.Hex(), sha256EmptyStringHashCompare, "comparing known hash values failed")
	var sha256_abc_Hash = NewHashSum(crypto.SHA256, []byte("abc"))
	assert.Equal(t, sha256_abc_Hash.Hex(), sha256_abc_HashCompare, "comparing known hash values failed")
}

func Test_Crypto3_256_HashSum(t *testing.T) {
	var sha3_256EmptyString = NewHashSum(crypto.SHA3_256, []byte(""))
	assert.Equal(t, sha3_256EmptyString.Hex(), sha3_256EmptyStringHashCompare, "comparing known hash values failed")
	var sha3_256_abc_Hash = NewHashSum(crypto.SHA3_256, []byte("abc"))
	assert.Equal(t, sha3_256_abc_Hash.Hex(), sha3_256_abc_HashCompare, "comparing known hash values failed")
}

var (
	sha256EmptyStringHash          = "e3b0c442 98fc1c14 9afbf4c8 996fb924 27ae41e4 649b934c a495991b 7852b855"
	sha3_256EmptyStringHash        = "a7ffc6f8bf1ed766 51c14756a061d662 f580ff4de43b49fa 82d80a4b80f8434a"
	sha256EmptyStringHashCompare   = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	sha3_256EmptyStringHashCompare = "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a"

	sha256_abc_Hash           = "ba7816bf 8f01cfea 414140de 5dae2223 b00361a3 96177a9c b410ff61 f20015ad"
	sha256_abc_HashCompare    = "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	sha256_abc_Base58_Compare = Base58Encode([]byte("ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"))
	sha3_256_abc_Hash         = "3a985da74fe225b2 045c172d6bd390bd 855f086e3e9d525b 46bfe24511431532"
	sha3_256_abc_HashCompare  = "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532"
)
