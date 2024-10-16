package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

// Структура для ключей Рабина
type Rabin struct {
	P *big.Int
	Q *big.Int
	N *big.Int
}

// Генерация простых чисел P и Q, таких что P ≡ 3 (mod 4) и Q ≡ 3 (mod 4)
func GenerateKeys(bitSize int) *Rabin {
	p := generatePrimeMod3Mod4(bitSize / 2)
	q := generatePrimeMod3Mod4(bitSize / 2)
	n := new(big.Int).Mul(p, q)
	return &Rabin{P: p, Q: q, N: n}
}

// Генерация простого числа, которое удовлетворяет условию p ≡ 3 (mod 4)
func generatePrimeMod3Mod4(bits int) *big.Int {
	for {
		prime, _ := rand.Prime(rand.Reader, bits)
		// Проверяем условие prime % 4 == 3
		if new(big.Int).Mod(prime, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
			return prime
		}
	}
}

// Сохранение открытого ключа (N) в файл
func SavePublicKey(fileName string, N *big.Int) error {
	return os.WriteFile(fileName, []byte(N.Text(16)), 0644)
}

// Сохранение закрытых ключей (P и Q) в файл
func SavePrivateKey(fileName string, P, Q *big.Int) error {
	data := fmt.Sprintf("%s\n%s", P.Text(16), Q.Text(16))
	return os.WriteFile(fileName, []byte(data), 0644)
}

// Загрузка открытого ключа из файла
func LoadPublicKey(fileName string) (*big.Int, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	n := new(big.Int)
	n.SetString(string(data), 16)
	return n, nil
}

// Загрузка закрытых ключей из файла
func LoadPrivateKey(fileName string) (*big.Int, *big.Int, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}
	var pStr, qStr string
	_, err = fmt.Sscanf(string(data), "%s\n%s", &pStr, &qStr)
	if err != nil {
		return nil, nil, err
	}
	p := new(big.Int)
	q := new(big.Int)
	p.SetString(pStr, 16)
	q.SetString(qStr, 16)
	return p, q, nil
}

// Шифрование: c = m^2 mod N
func (r *Rabin) Encrypt(m *big.Int) *big.Int {
	c := new(big.Int).Exp(m, big.NewInt(2), r.N)
	return c
}

// Дешифрование: находим 4 квадратных корня по модулю N
func (r *Rabin) Decrypt(c *big.Int) []*big.Int {
	// mp = c^(p+1)/4 mod p
	mp := new(big.Int).Exp(c, new(big.Int).Div(new(big.Int).Add(r.P, big.NewInt(1)), big.NewInt(4)), r.P)
	// mq = c^(q+1)/4 mod q
	mq := new(big.Int).Exp(c, new(big.Int).Div(new(big.Int).Add(r.Q, big.NewInt(1)), big.NewInt(4)), r.Q)

	// Китайская теорема об остатках (CRT)
	r1 := r.crt(mp, mq)
	r2 := r.crt(new(big.Int).Sub(r.P, mp), mq)
	r3 := r.crt(mp, new(big.Int).Sub(r.Q, mq))
	r4 := r.crt(new(big.Int).Sub(r.P, mp), new(big.Int).Sub(r.Q, mq))

	return []*big.Int{r1, r2, r3, r4}
}

// Китайская теорема об остатках (CRT)
func (r *Rabin) crt(mp, mq *big.Int) *big.Int {
	qInv := new(big.Int).ModInverse(r.Q, r.P)
	h := new(big.Int).Mul(qInv, new(big.Int).Sub(mp, mq))
	h.Mod(h, r.P)
	m := new(big.Int).Add(mq, new(big.Int).Mul(h, r.Q))
	return m.Mod(m, r.N)
}

// Чтение файла в виде большого числа
func FileToBigInt(filePath string) (*big.Int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

// Запись большого числа в файл
func BigIntToFile(n *big.Int, filePath string) error {
	return os.WriteFile(filePath, n.Bytes(), 0644)
}

func main() {
	// Опции для использования программы
	var option string
	fmt.Println("Выберите опцию: ")
	fmt.Println("1 - Генерация ключей")
	fmt.Println("2 - Шифрование файла")
	fmt.Println("3 - Дешифрование файла")
	fmt.Scan(&option)

	switch option {
	case "1":
		// Генерация ключей
		bitSize := 512
		rabin := GenerateKeys(bitSize)

		// Сохранение ключей
		publicKeyFile := "public_key.txt"
		privateKeyFile := "private_key.txt"

		err := SavePublicKey(publicKeyFile, rabin.N)
		if err != nil {
			fmt.Println("Ошибка при сохранении открытого ключа:", err)
			return
		}
		fmt.Printf("Открытый ключ сохранен в файл %s\n", publicKeyFile)

		err = SavePrivateKey(privateKeyFile, rabin.P, rabin.Q)
		if err != nil {
			fmt.Println("Ошибка при сохранении закрытого ключа:", err)
			return
		}
		fmt.Printf("Закрытый ключ сохранен в файл %s\n", privateKeyFile)

	case "2":
		// Шифрование файла
		publicKeyFile := "public_key.txt"
		inputFile := "input.txt"
		encryptedFile := "encrypted.txt"

		// Загружаем открытый ключ
		N, err := LoadPublicKey(publicKeyFile)
		if err != nil {
			fmt.Println("Ошибка при загрузке открытого ключа:", err)
			return
		}

		// Чтение входного файла
		m, err := FileToBigInt(inputFile)
		if err != nil {
			fmt.Println("Ошибка чтения файла:", err)
			return
		}

		// Шифрование
		rabin := &Rabin{N: N}
		c := rabin.Encrypt(m)

		// Запись зашифрованного сообщения в файл
		err = BigIntToFile(c, encryptedFile)
		if err != nil {
			fmt.Println("Ошибка записи зашифрованного файла:", err)
			return
		}
		fmt.Printf("Файл %s зашифрован и сохранен как %s\n", inputFile, encryptedFile)

	case "3":
		// Дешифрование файла
		privateKeyFile := "private_key.txt"
		encryptedFile := "encrypted.txt"

		// Загружаем закрытые ключи
		p, q, err := LoadPrivateKey(privateKeyFile)
		if err != nil {
			fmt.Println("Ошибка при загрузке закрытых ключей:", err)
			return
		}

		// Чтение зашифрованного файла
		c, err := FileToBigInt(encryptedFile)
		if err != nil {
			fmt.Println("Ошибка чтения зашифрованного файла:", err)
			return
		}

		// Дешифрование
		rabin := &Rabin{P: p, Q: q, N: new(big.Int).Mul(p, q)}
		roots := rabin.Decrypt(c)

		// Запись одного из корней в файл (например, первого)
		for i, root := range roots {
			decryptedFile := ""

			err = BigIntToFile(root, fmt.Sprintf("decrypted/decrypted%v.txt", i))
			if err != nil {
				fmt.Println("Ошибка записи расшифрованного файла:", err)
				return
			}
			fmt.Printf("Файл %s расшифрован и сохранен как %s\n", encryptedFile, decryptedFile)
		}

	default:
		fmt.Println("Неверная опция.")
	}
}
