package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/kyokomi/emoji/v2"
)

func extractKeys(emojies map[string]string) []string {
	var res []string
	for k, _ := range emojies {
		res = append(res, k)
	}
	return res
}

func selectedKeys(keys []string, num int) []string {
	var sks []string
	for i := 0; i < num; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(keys))))
		if err != nil {
			panic(err)
		}
		sks = append(sks, keys[r.Int64()])
	}
	return sks
}

func formatKey(key string) string {
	return fmt.Sprintf("put %v in the pair of : to get %v", strings.Trim(key, ":"), key)
}

func main() {
	codeMap := emoji.CodeMap()
	k := extractKeys(codeMap)
	sk := selectedKeys(k, 3)
	for _, rk := range sk {
		emoji.Println(formatKey(rk))
	}
}
