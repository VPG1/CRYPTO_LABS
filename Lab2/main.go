package main

import (
	"CRYPTO_LABS/Lab2/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// key: 1f2e3f4f5f6f7f8f1f2e3f4f5f6f7f8f1f2e3f4f5f6f7f8f1f2e3f4f5f6f7f89
// sync: 1f2e3f4f1f2e3f4f1f2e3f4f1f2e3f4f

func generateRandomBytes(n int) []byte {
	// Создаем массив для случайных байтов длиной n
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err) // Обрабатываем ошибку (например, недостаток случайных данных)
	}
	return bytes
}

func main() {
	fmt.Println("Enter data:")
	var data string
	_, err := fmt.Scanln(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter key(len 32 bytes):")
	var keyStr string
	_, err = fmt.Scanln(&keyStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	key, err := hex.DecodeString(keyStr)
	if err != nil {
		fmt.Println("The key is invalid")
		return
	}

	if len(key) != 32 {
		fmt.Println("The key must be 32 bytes long")
		return
	}

	fmt.Println("Enter sync message(len 16 bytes)")
	var syncMessageHex string
	_, err = fmt.Scanln(&syncMessageHex)
	if err != nil {
		fmt.Println(err)
		return
	}

	syncMessage, err := hex.DecodeString(syncMessageHex)
	if err != nil {
		fmt.Println("The sync message is invalid")
		return
	}

	if len(syncMessage) != 16 {
		fmt.Println("The sync message must be 16 bytes long")
		return
	}

	encryptData, err := cipher.CmEncrypt([]byte(data), key, syncMessage)

	fmt.Println("Encrypted data:")
	fmt.Println(string(encryptData))

	decryptData, err := cipher.CmEncrypt(encryptData, key, syncMessage)

	fmt.Println("Decrypted data:")
	fmt.Println(string(decryptData))
}
