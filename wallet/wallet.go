package wallet

import (
	crypto "crypto-wallet/crypto"
	"crypto/sha256"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type NewWalletBody struct {
	Phrase              string
	Address             string
	EncryptedMnemonic   string
	EncryptedPrivateKey string
}

func DeriveAESKey(password string) string {
	salt := []byte("A_SALT")
	aesKey := pbkdf2.Key([]byte(password), salt, 2048, 32, sha256.New)
	return string(aesKey)
}

func GeneratePhrases() ([]string, error) {
	mnemonic, err := crypto.GeneratePhrase()
	if err != nil {
		return nil, err
	}
	return mnemonic, nil
}

func CreateNewWallet(password string) *NewWalletBody {
	mnemonic, _ := GeneratePhrases()
	mnemonicStr := strings.Join(mnemonic, " ")

	var seed []byte = crypto.GenerateSeed(mnemonicStr, password)
	masterKey, masterChain := crypto.GenerateMasterKey(seed)
	childIndex := uint32(0x80000000)

	childKey, _, _ := crypto.DeriveHardenedChilds(masterKey, masterChain, childIndex)
	publicKey := crypto.PrivateKeyToPublicKey(childKey)
	address := crypto.GenerateP2PKeyAddress(publicKey)

	aesKey := DeriveAESKey(password)
	encryptedMnemonic, _ := crypto.AESEncrypt(mnemonicStr, aesKey)
	encryptedPrivateKey, _ := crypto.AESEncrypt(string(childKey), aesKey)

	return &NewWalletBody{
		Phrase: mnemonicStr,
		Address: address,
		EncryptedMnemonic: encryptedMnemonic,
		EncryptedPrivateKey: encryptedPrivateKey,
	}
}
