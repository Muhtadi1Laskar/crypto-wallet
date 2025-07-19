package wallet

import (
	crypto "crypto-wallet/crypto"
	"encoding/hex"
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

type RetriveWalletInfo struct {
	Address    string
	PublicKey  string
	PrivateKey string
}

func GeneratePhrases() ([]string, error) {
	mnemonic, err := crypto.GeneratePhrase()
	if err != nil {
		return nil, err
	}
	return mnemonic, nil
}

func CreateNewWallet(password string) (*NewWalletBody, error) {
	mnemonic, err := GeneratePhrases()
	if err != nil {
		return nil, err
	}
	mnemonicStr := strings.Join(mnemonic, " ")

	keys, err := GenerateKeysFromPhrase(mnemonicStr, password)
	if err != nil {
		return nil, err
	}

	aesKey := crypto.DeriveAESKey(password)
	encryptedMnemonic, err := crypto.AESEncrypt(mnemonicStr, aesKey)
	if err != nil {
		return nil, err
	}

	encryptedPrivateKey, err := crypto.AESEncrypt(string(keys.PrivateKey), aesKey)
	if err != nil {
		return nil, err
	}

	return &NewWalletBody{
		Phrase:              mnemonicStr,
		Address:             keys.Address,
		EncryptedMnemonic:   encryptedMnemonic,
		EncryptedPrivateKey: encryptedPrivateKey,
	}, nil
}

func GenerateKeysFromPhrase(phrase, password string) (*WalletKeys, error) {
	var seed []byte = crypto.GenerateSeed(phrase, password)
	masterKey, masterChain := crypto.GenerateMasterKey(seed)
	childIndex := uint32(0x80000000)

	childKey, _, err := crypto.DeriveHardenedChilds(masterKey, masterChain, childIndex)
	if err != nil {
		return nil, err
	}

	publicKey := crypto.PrivateKeyToPublicKey(childKey)
	address := crypto.GenerateP2PKeyAddress(publicKey)

	return &WalletKeys{
		Address:    address,
		PublicKey:  publicKey,
		PrivateKey: childKey,
	}, nil
}

func RetriveExistingWallet(password, encryptedPhrase string) (*RetriveWalletInfo, error) {
	aesKey := crypto.DeriveAESKey(password)
	mnemonic, err := crypto.AESDecrypt(encryptedPhrase, aesKey)
	if err != nil {
		return nil, err
	}

	keys, err := GenerateKeysFromPhrase(mnemonic, password)
	if err != nil {
		return nil, err
	}

	return &RetriveWalletInfo{
		Address:    keys.Address,
		PublicKey:  hex.EncodeToString(keys.PublicKey),
		PrivateKey: hex.EncodeToString(keys.PrivateKey),
	}, nil
}
