package mkpass

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// GenerateDicewarePassword based on EFF Dice-Generated Passphrases
// https://www.eff.org/dice
// A variation of the method is used in here where the source of
// random generation entropy is replaced by crypto rand library
// from go standard library
// The entropy on this method is:
// PasswordList = 7776
// possibleCombinations = PasswordList ^ length
// entropy = log_2(possibleCombinations)
// For example:
// length = 5 => entropy = log_2(2.843028803×10¹⁹) = 64.624062518
// length = 5 => entropy = log_2(2.2107×10²³) = 77.548875022
// to reach the recommended 64 bits of entropy, for a strong password
// length of 5 is required
func GenerateDicewarePassword(length int) string {
	buf, err := os.ReadFile("./diceware.txt")
	if err != nil {
		// if the password file is not available, then we use base64
		// words of length 3 which produced slightly higher entropy
		return GenerateBase64(length * 3)
	}
	list := strings.Split(string(buf), "\n")

	passwordList := make([]string, length)

	for i := 0; i < length; i++ {
		passwordList[i] = list[GetRandomNumber(len(list))]
	}

	return strings.Join(passwordList, " ")
}

func GetRandomNumber(to int) int {
	base := new(big.Int)
	base.SetString(fmt.Sprint(to), 10)

	randInt, _ := rand.Int(rand.Reader, base)
	return int(randInt.Int64())
}

// GenerateBase64 generates a base64 string of length length
func GenerateBase64(length int) string {
	base := new(big.Int)
	base.SetString("64", 10)

	base64 := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
	tempKey := ""
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, base)
		tempKey += string(base64[int(index.Int64())])
	}
	return tempKey
}
