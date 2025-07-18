package crypto

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"crypto/sha256"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

var curve = btcec.S256()
var curveN = curve.Params().N

func uint32ToBytes(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func DeriveHardenedChilds(parentPrivKey, parentChainCode []byte, index uint32) ([]byte, []byte, error) {
	if index < 0x80000000 {
		return nil, nil, fmt.Errorf("index must be >= 0x80000000 for hardened derivation")
	}

	data := append([]byte{0x00}, parentPrivKey...)
	data = append(data, uint32ToBytes(index)...)

	mac := hmac.New(sha512.New, parentChainCode)
	mac.Write(data)
	I := mac.Sum(nil)

	IL := I[:32]
	IR := I[32:]

	ilInt := new(big.Int).SetBytes(IL)
	kParInt := new(big.Int).SetBytes(parentPrivKey)

	childKeyInt := new(big.Int).Add(ilInt, kParInt)
	childKeyInt.Mod(childKeyInt, curveN)

	if childKeyInt.Sign() == 0 {
		return nil, nil, fmt.Errorf("derived key is zero")
	}

	childKey := childKeyInt.Bytes()

	if len(childKey) < 32 {
		padded := make([]byte, 32)
		copy(padded[32-len(childKey):], childKey)
		childKey = padded
	}

	return childKey, IR, nil
}


func PrivateKeyToPublicKey(privateKey []byte) []byte {
	priv, _ := btcec.PrivKeyFromBytes(privateKey)
	pubKey := priv.PubKey()
	return pubKey.SerializeCompressed()
}

func publicKeyHash(pubKey []byte) []byte {
	sha := sha256.Sum256(pubKey)

	ripmd := ripemd160.New()
	ripmd.Write(sha[:])
	return ripmd.Sum(nil)
}

func addVersion(pubKeyHash []byte) []byte {
	return append([]byte{0x00}, pubKeyHash...)
}

func checkSum(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return second[:4]
}

func base58CheckEncode(data []byte) string {
	ck := checkSum(data)
	full := append(data, ck...)
	return base58.Encode(full)
}

func GenerateP2PKeyAddress(privateKey []byte) string {
	pubKey := PrivateKeyToPublicKey(privateKey)
	pubKeyHashed := publicKeyHash(pubKey)
	versioned := addVersion(pubKeyHashed)
	return base58CheckEncode(versioned)
}


func hmacSha512(seed []byte) []byte {
	secrectKey := []byte("Bitcoin seed")

	h := hmac.New(sha512.New, secrectKey)
	h.Write(seed)

	return h.Sum(nil)
}

func GenerateMasterKey(seed []byte) ([]byte, []byte) {
	keyedHash := hmacSha512(seed)
	IL := keyedHash[:32]
	IR := keyedHash[32:]

	return IL, IR
}