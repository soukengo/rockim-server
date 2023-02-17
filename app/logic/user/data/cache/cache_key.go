package cache

import (
	"fmt"
	"strings"
)

const keySplit = ":"

const (
	keyUser        = "user"
	keyUserAccount = "user.account"
)

func genKey(keys ...interface{}) string {
	var strArr []string
	for _, v := range keys {
		switch v1 := v.(type) {
		case string:
			strArr = append(strArr, v1)
		default:
			strArr = append(strArr, fmt.Sprintf("%v", v))
		}
	}
	return strings.Join(strArr, keySplit)
}

func GenUserCacheKey(productId string, uid string) string {
	return genKey(keyUser, productId, uid)
}
func GenUserAccountCacheKey(productId string, account string) string {
	return genKey(keyUserAccount, productId, account)
}
