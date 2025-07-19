package wallet

import (
	crypto "crypto-wallet/crypto"
	"fmt"
	"strings"
)

type NewWalletBody struct {
	Phrase              string
	Address             string
	EncryptedMnemonic   string
	EncryptedPrivateKey string
}

type WalletKeys struct {
	Address    string
	PublicKey  []byte
	PrivateKey []byte
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

	keys := GenerateKeysFromPhrase(mnemonicStr, password)

	aesKey := crypto.DeriveAESKey(password)
	encryptedMnemonic, _ := crypto.AESEncrypt(mnemonicStr, aesKey)
	encryptedPrivateKey, _ := crypto.AESEncrypt(string(keys.PrivateKey), aesKey)

	return &NewWalletBody{
		Phrase: mnemonicStr,
		Address: keys.Address,
		EncryptedMnemonic: encryptedMnemonic,
		EncryptedPrivateKey: encryptedPrivateKey,
	}
}


func GenerateKeysFromPhrase(phrase, password string) *WalletKeys {
	var seed []byte = crypto.GenerateSeed(phrase, password)
	masterKey, masterChain := crypto.GenerateMasterKey(seed)
	childIndex := uint32(0x80000000)

	childKey, _, _ := crypto.DeriveHardenedChilds(masterKey, masterChain, childIndex)
	publicKey := crypto.PrivateKeyToPublicKey(childKey)
	address := crypto.GenerateP2PKeyAddress(publicKey)

	return &WalletKeys{
		Address:    address,
		PublicKey:  publicKey,
		PrivateKey: childKey,
	}
}

func RetriveExistingWallet(password, encryptedPhrase string) {
	aesKey := crypto.DeriveAESKey(password)
	mnemonic, _ := crypto.AESDecrypt(encryptedPhrase, aesKey)

	fmt.Println(mnemonic)
}
