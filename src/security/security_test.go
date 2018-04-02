package security

import (
	"encoding/hex"
	"testing"
)

const orgData string = "abcdefghijklmnopqrstuvwsyz0123456789"
const key string = "1234567812345678"

func TestGetMD5Hash(t *testing.T) {
	md5Value := GetMD5Hash(orgData)
	t.Errorf("md5 result:%s", md5Value)
}

func TestAesEncrypt(t *testing.T) {
	crypted, err := AesEncrypt([]byte(orgData), []byte(key))
	if err != nil {
		t.Errorf("aes encryt error:%s", err.Error())
	}

	t.Errorf("aes encryt result:%s", hex.EncodeToString(crypted))

	decrypted, err := AesDecrypt(crypted, []byte(key))
	if err != nil {
		t.Errorf("aes decrypt error:%s", err.Error())
	}

	t.Errorf("aes decrypt result:%s", string(decrypted))
}

func TestUniqueId(t *testing.T) {
	uuid :=	UniqueId()
	t.Errorf("uuid:%s", uuid)
}