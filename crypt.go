package xdripgo

import (
	"crypto/aes"
	log "github.com/Sirupsen/logrus"
	"github.com/andreburgaud/crypt2go/ecb"
)

func cryptKey(id string) string {
	key := "00" + id + "00" + id
	return key
}

func encrypt(buffer []byte, id string) []byte {
	key := []byte(cryptKey(id))
	log.Debugf("key=%x", key)
	return encryptBytes(buffer, key)
}

func encryptBytes(pt, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	log.Debugf("mode=%v", mode)
	log.Debugf("pt=%x", pt)
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return ct
}

func calculateHash(data []byte, id string) []byte {
	if len(data) != 8 {
		log.Fatalf("calculateHash failed data(%x) not length of 8", data)
	}
	doubleData := make([]byte, 16)
	copy(doubleData[0:8], data)
	copy(doubleData[8:16], data)
	log.Debugf("doubleData=%x", doubleData)

	encrypted := encrypt(doubleData, id)
	encrypted_return := make([]byte, 8)
	copy(encrypted_return, encrypted[0:8])
	log.Debugf("encrypted=%x", encrypted)
	log.Debugf("encrypted_return=%x", encrypted_return)
	return encrypted_return
}
