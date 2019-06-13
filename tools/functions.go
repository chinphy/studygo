package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/scrypt"
)

// GetRandomString 生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// StringToMD5 生成32位MD5
func StringToMD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// BytesToMD5 生成32位MD5
func BytesToMD5(bytes []byte) string {
	ctx := md5.New()
	ctx.Write(bytes)

	return hex.EncodeToString(ctx.Sum(nil))
}

// PasswordHash 密码hash
func PasswordHash(password *string, salt *string) string {
	hash, _ := scrypt.Key([]byte(*password), []byte(*salt), 1<<15, 8, 1, 32)

	return BytesToMD5(hash)
}

// PasswordVerify 验证密码是否和散列值匹配
func PasswordVerify(password *string, salt *string, hashStr *string) bool {
	newHash := PasswordHash(password, salt)
	fmt.Println("salt:" + *salt)
	fmt.Println("password:" + *password)
	fmt.Println("newHash:" + newHash)
	fmt.Println("oldHash:" + *hashStr)
	return 0 == strings.Compare(newHash, *hashStr)
}
