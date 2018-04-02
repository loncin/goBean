package httpReq

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

const orgData string = "abcdefg123456"
const orgData2 string = "H4sIAAAAAAAAAFXMwQrCMAyA4XfJuYd24JAdpyAqeNGLx26tGOjWEtLDEN/d1G2H3fL9CflAouhyz48peWiqWgETpllQgYIXxeGAPAnby10Cx5XX87I/Wv6fa7PXO11LJc+baIxE63LgWx46T9AYBf0bg1utxZYIyzzmEAo7HE9kXXnzhO8PKs0nAKwAAAA="

func TestGzipEncode(t *testing.T) {
	encode, err := GzipEncode([]byte(orgData))
	if err != nil {
		t.Errorf("TestGzipEncode error:%s", err.Error())
	}

	t.Errorf("TestGzipEncode result:%s", base64.StdEncoding.EncodeToString(encode))

	decode, err := GzipDecode(encode)
	if err != nil {
		t.Errorf("TestGzipDecode error:%s", err.Error())
	}

	decodeStr := string(decode)
	if string(decode) != orgData {
		t.Errorf("TestGzipDecode failed:%s", decodeStr)
	}

	t.Errorf("TestGzipDecode result:%s", decodeStr)
}

func TestGzipDecode(t *testing.T) {
	base64de, err := base64.StdEncoding.DecodeString(orgData2)
	
	if err != nil {
		t.Errorf("TestGzipDecode error:%s", err.Error())
	}

	t.Errorf("TestGzipDecode base64Decode result:%s", hex.EncodeToString(base64de))

	gzipde, err := GzipDecode([]byte(base64de))
	if err != nil {
		t.Errorf("TestGzipDecode error:%s", err.Error())
	}

	t.Errorf("TestGzipDecode result:%s", hex.EncodeToString(gzipde))
}