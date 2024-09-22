package main

import (
	"CRYPTO_LABS/Lab2/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
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
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run CRYPTO_LABS <file>")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Enter key(len 32 bytes):")
	var keyStr string
	_, err = fmt.Scanln(&keyStr)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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

	encryptData, err := cipher.CmEncrypt(data, key, syncMessage)
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = file.Write(encryptData)
	if err != nil {
		log.Fatal(err)
		return
	}
}
