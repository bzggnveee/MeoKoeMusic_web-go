package util

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"strings"
)

// CryptoMD5 计算 MD5 哈希（十六进制字符串）
func CryptoMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// CryptoSHA1 计算 SHA1 哈希（十六进制字符串）
func CryptoSHA1(data string) string {
	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// AesEncryptResult AES 加密结果
type AesEncryptResult struct {
	Str string `json:"str"` // hex ciphertext
	Key string `json:"key"` // random key string (16 lowercase chars)
}

// CryptoAesEncrypt AES-CBC 加密，返回 hex 密文
// optKey/optIV 按 UTF-8 字符串逐字节作为密钥（与 CryptoJS 的 Utf8.parse 一致）
func CryptoAesEncrypt(data string, optKey, optIV string) (result AesEncryptResult) {
	if s, ok := tryJSONStr(data); ok {
		data = s
	}
	plaintext := []byte(data)

	var keyBytes, ivBytes []byte
	var tempKey string

	if optKey != "" && optIV != "" {
		// 提供的 key/iv 按 UTF-8 字节处理（与 CryptoJS.enc.Utf8.parse 一致）
		keyBytes = []byte(optKey)
		ivBytes = []byte(optIV)
	} else {
		if optKey != "" {
			tempKey = optKey
		} else {
			tempKey = strings.ToLower(RandomString(16))
		}
		md5Hash := CryptoMD5(tempKey)        // 32 hex chars
keyBytes = []byte(md5Hash[:32])       // 32 bytes = AES-256
		keyBytes = safeKeyBytes(keyBytes); ivBytes = safeIVBytes(ivBytes)
ivBytes = safeIVBytes(ivBytes)
		result.Key = tempKey
	}

	// AES key length: 32 bytes → AES-256, 24 bytes → AES-192, 16 bytes → AES-128
	if len(keyBytes) != 16 && len(keyBytes) != 24 && len(keyBytes) != 32 {
		if len(keyBytes) < 16 {
			newKey := make([]byte, 16)
			copy(newKey, keyBytes)
			keyBytes = newKey
		} else if len(keyBytes) < 24 {
			keyBytes = safeKeyBytes(keyBytes)
		} else {
			keyBytes = keyBytes[:24]
		}
	}
	if len(ivBytes) < 16 {
		newIV := make([]byte, 16)
		copy(newIV, ivBytes)
		ivBytes = newIV
	}
	ivBytes = ivBytes[:16]

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		result.Str = ""
		result.Key = tempKey
		return
	}

	plaintext = pkcs7Pad(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, ivBytes)
	mode.CryptBlocks(ciphertext, plaintext)

	result.Str = hex.EncodeToString(ciphertext)
	if result.Key == "" {
		result.Key = tempKey
	}
	return
}

// CryptoAesDecrypt AES-CBC 解密（输入 hex 密文）
func CryptoAesDecrypt(encryptedHex string, keyStr string, optIV string) interface{} {
	// 与 Node.js 一致：不传 iv 时从 key MD5 派生
	if optIV == "" {
		keyStr = CryptoMD5(keyStr)[:32]
		optIV = keyStr[len(keyStr)-16:]
	}

	keyBytes := []byte(keyStr)
	if len(keyBytes) < 16 {
		newKey := make([]byte, 16)
		copy(newKey, keyBytes)
		keyBytes = newKey
	}
	keyBytes = safeKeyBytes(keyBytes)

	ivBytes := []byte(optIV)
	if len(ivBytes) < 16 {
		newIV := make([]byte, 16)
		copy(newIV, ivBytes)
		ivBytes = newIV
	}
	ivBytes = ivBytes[:16]

	encrypted, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return encryptedHex
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return encryptedHex
	}

	if len(encrypted) < aes.BlockSize || len(encrypted)%aes.BlockSize != 0 {
		// Try ECB fallback
		return aesECBDecryptStr(encrypted, keyBytes)
	}

	mode := cipher.NewCBCDecrypter(block, ivBytes)
	plaintext := make([]byte, len(encrypted))
	mode.CryptBlocks(plaintext, encrypted)
	plaintext = pkcs7Unpad(plaintext)

	var obj interface{}
	if err := json.Unmarshal(plaintext, &obj); err == nil {
		return obj
	}
	return string(plaintext)
}

func aesECBDecryptStr(ciphertext, key []byte) string {
	block, _ := aes.NewCipher(key)
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(plaintext[i:i+aes.BlockSize], ciphertext[i:i+aes.BlockSize])
	}
	plaintext = pkcs7Unpad(plaintext)
	return string(plaintext)
}

// ============================================================
// RSA
// ============================================================

const publicRsaKeyPEM = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDIAG7QOELSYoIJvTFJhMpe1s/gbjDJX51HBNnEl5HXqTW6lQ7LC8jr9fWZTwusknp+sVGzwd40MwP6U5yDE27M/X1+UR4tvOGOqp94TJtQ1EPnWGWXngpeIW5GxoQGao1rmYWAu6oi1z9XkChrsUdC6DJE5E221wf/4WLFxwAtRQIDAQAB
-----END PUBLIC KEY-----`

const publicLiteRsaKeyPEM = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDECi0Np2UR87scwrvTr72L6oO01rBbbBPriSDFPxr3Z5syug0O24QyQO8bg27+0+4kBzTBTBOZ/WWU0WryL1JSXRTXLgFVxtzIY41Pe7lPOgsfTCn5kZcvKhYKJesKnnJDNr5/abvTGf+rHG3YRwsCHcQ08/q6ifSioBszvb3QiwIDAQAB
-----END PUBLIC KEY-----`

func parseRSAPublicKey(pemStr string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	return pub.(*rsa.PublicKey)
}

// CryptoRSAEncrypt 原始 RSA 加密（零填充到 keyLength，与 Node.js node-forge raw 一致）
func CryptoRSAEncrypt(data map[string]interface{}) string {
	keyPem := publicRsaKeyPEM
	if IsLite() {
		keyPem = publicLiteRsaKeyPEM
	}
	pubKey := parseRSAPublicKey(keyPem)
	if pubKey == nil {
		return ""
	}

	jsonBytes, _ := json.Marshal(data)
	buffer := jsonBytes
	keyLen := pubKey.Size() // keyLength in bytes

	// zero-pad at end: Node.js `new Uint8Array(keyLen); padded.set(buffer)` → data first, zeros last
	padded := make([]byte, keyLen)
	copy(padded, buffer)

	// Raw RSA: m^e mod n
	m := new(big.Int).SetBytes(padded)
	encrypted := new(big.Int).Exp(m, big.NewInt(int64(pubKey.E)), pubKey.N)
	hexStr := fmt.Sprintf("%x", encrypted)
	if len(hexStr) < keyLen*2 {
		hexStr = strings.Repeat("0", keyLen*2-len(hexStr)) + hexStr
	}
	return strings.ToUpper(hexStr)
}

// CryptoRSAEncryptString 使用指定 JSON 字符串的 RSA 加密
// Node.js JSON.stringify 保留 key 插入顺序，Go map 按字母序。
// 此函数接受预格式化的 JSON 字符串以匹配 Node.js 的 key 顺序。
func CryptoRSAEncryptString(jsonStr string) string {
	keyPem := publicRsaKeyPEM
	if IsLite() {
		keyPem = publicLiteRsaKeyPEM
	}
	pubKey := parseRSAPublicKey(keyPem)
	if pubKey == nil {
		return ""
	}
	buffer := []byte(jsonStr)
	keyLen := pubKey.Size()
	padded := make([]byte, keyLen)
	copy(padded, buffer)
	m := new(big.Int).SetBytes(padded)
	encrypted := new(big.Int).Exp(m, big.NewInt(int64(pubKey.E)), pubKey.N)
	hexStr := fmt.Sprintf("%x", encrypted)
	if len(hexStr) < keyLen*2 {
		hexStr = strings.Repeat("0", keyLen*2-len(hexStr)) + hexStr
	}
	return strings.ToUpper(hexStr)
}

// RsaEncrypt2 PKCS#1 v1.5 RSA 加密（标准 RSA）
func RsaEncrypt2(data map[string]interface{}) string {
	keyPem := publicRsaKeyPEM
	if IsLite() {
		keyPem = publicLiteRsaKeyPEM
	}
	pubKey := parseRSAPublicKey(keyPem)
	if pubKey == nil {
		return ""
	}

	jsonBytes, _ := json.Marshal(data)
	encrypted, err := rsa.EncryptPKCS1v15(crand.Reader, pubKey, jsonBytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(encrypted)
}

// ============================================================
// KRC 歌词解码
// ============================================================

var krcEnKey = []byte{64, 71, 97, 119, 94, 50, 116, 71, 81, 54, 49, 45, 206, 210, 110, 105}

func DecodeLyrics(val interface{}) string {
	var rawBytes []byte
	switch v := val.(type) {
	case string:
		var err error
		rawBytes, err = base64.StdEncoding.DecodeString(v)
		if err != nil {
			return ""
		}
	case []byte:
		rawBytes = v
	default:
		return ""
	}

	if len(rawBytes) < 4 {
		return ""
	}

	krcBytes := rawBytes[4:]
	for i := range krcBytes {
		krcBytes[i] ^= krcEnKey[i%len(krcEnKey)]
	}

	b := bytes.NewReader(krcBytes)
	r, err := zlib.NewReader(b)
	if err != nil {
		return ""
	}
	defer r.Close()

	var out bytes.Buffer
	io.Copy(&out, r)
	return out.String()
}

// ============================================================
// PKCS7
// ============================================================

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pkcs7Unpad(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	padding := int(data[len(data)-1])
	if padding > len(data) || padding > aes.BlockSize {
		return data
	}
	return data[:len(data)-padding]
}

// ============================================================
// 辅助
// ============================================================

func tryJSONStr(s string) (string, bool) {
	s = strings.TrimSpace(s)
	if len(s) > 0 && (s[0] == '{' || s[0] == '[') {
		return s, true
	}
	return s, false
}

// PlaylistAesEncrypt 歌单操作 AES 加密（base64 密文，6 字节随机 key）
func PlaylistAesEncrypt(data map[string]interface{}) AesEncryptResult {
	jsonBytes, _ := json.Marshal(data)
	key := strings.ToLower(RandomString(6))
	md5Hash := CryptoMD5(key)
	encKey := []byte(md5Hash[:16])  // 16 bytes
	iv := []byte(md5Hash[16:32])    // 16 bytes
	block, _ := aes.NewCipher(encKey)
	plaintext := jsonBytes
	plaintext = pkcs7Pad(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return AesEncryptResult{Str: base64.StdEncoding.EncodeToString(ciphertext), Key: key}
}

// PlaylistAesDecrypt 歌单操作 AES 解密（base64 密文）
func PlaylistAesDecrypt(encryptedBase64 string, key string) interface{} {
	md5Hash := CryptoMD5(key)
	encKey := []byte(md5Hash[:16])
	iv := []byte(md5Hash[16:32])
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil { return encryptedBase64 }
	block, _ := aes.NewCipher(encKey)
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)
	plaintext = pkcs7Unpad(plaintext)
	var obj interface{}
	if err := json.Unmarshal(plaintext, &obj); err == nil { return obj }
	return string(plaintext)
}

func BigIntHex(hexStr string) *big.Int {
	n := new(big.Int)
	n.SetString(hexStr, 16)
	return n
}

func EncodeUtf8(str string) []byte    { return []byte(str) }
func DecodeUtf8(b []byte) string       { return string(b) }
func GenerateMD5(data string) string   { return CryptoMD5(data) }
func BufferFromBase64(s string) ([]byte, error) { return base64.StdEncoding.DecodeString(s) }
func BufferToString(b []byte) string   { return string(b) }

var _ = fmt.Sprintf

func safeKeyBytes(b []byte) []byte {
	if len(b) == 16 || len(b) == 24 || len(b) == 32 { return b }
	if len(b) < 16 { p := make([]byte, 16); copy(p, b); return p }
	return b[:16]
}
func safeIVBytes(b []byte) []byte {
	if len(b) >= 16 { return b[:16] }
	p := make([]byte, 16); copy(p, b); return p
}
