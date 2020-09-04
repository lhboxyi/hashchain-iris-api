package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

/**
创建JWT
*/
func CreateJwt(pwd string, claimsMap map[string]interface{}) (tokenStr string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	if len(claimsMap) > 0 {
		for k, v := range claimsMap {
			claims[k] = v
		}
	}
	//claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	//claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenStr, err = token.SignedString([]byte(pwd))
	if err != nil {
		fmt.Println(err)
	}
	return
}

/**
token验证
*/
func ValidateToken(tokenStr string, pwd string) (checkTooken bool, err error) {
	checkTag := false
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(pwd), nil
	})
	if err != nil {
		// 第一种
		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// fmt.Println("+++")
		// return
		//}
		//fmt.Println([]byte("vector.sign"))
		// 第二种
		if err, ok := err.(*jwt.ValidationError); ok {
			if err.Errors&jwt.ValidationErrorMalformed != 0 {
				return checkTag, err
			}
			if err.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println(err)
				return checkTag, err
			}
		}
	}
	checkTag = token.Valid
	//finToken := token.Claims.(jwt.MapClaims)
	//fmt.Println(finToken["iss"])
	return checkTag, err
}

/**
取得元素
*/
func Getclaims(tokenStr string, pwd string, claimsArrayKey []string) (claimsArray []interface{}, err error) {
	claimsArray = make([]interface{}, 1)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(pwd), nil
	})
	finToken := token.Claims.(jwt.MapClaims)
	if len(finToken) > 0 {
		for _, v := range claimsArrayKey {
			claim := finToken[v]
			claimsArray = append(claimsArray, claim)
		}
	}
	// 内存优化
	// 切片保留对底层数组的引用。只要切片存在于内存中，数组就不能被垃圾回收。这在内存管理方便可能是值得关注的。假设我们有一个非常大的数组，而我们只需要处理它的一小部分，为此我们创建这个数组的一个切片，并处理这个切片。这里要注意的事情是，数组仍然存在于内存中，因为切片正在引用它。
	// 解决该问题的一个方法是使用 copy 函数 func copy(dst, src []T) int 来创建该切片的一个拷贝。这样我们就可以使用这个新的切片，原来的数组可以被垃圾回收。
	newSinceArray := make([]interface{}, len(claimsArray))
	copy(newSinceArray, claimsArray)
	return newSinceArray, err
}

/**
根据解密前的token字符串解析token，获取jwt.MapClaims对象
根据key值获取对应的用户信息
用户名： "user_name"
用户真实姓名："real_name"
项目接入appCode：" app_code"
用户类型："user_type"   （1超级管理员，2管理员，3普通用户）
部门编码： " department"
用户邮箱：" user_email"
用户手机：" user_tel"
用户职务：" job_title "
*/
func ParserJwt(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("my_secret_key"), nil
	})

	var claims jwt.MapClaims
	var ok bool

	if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	return claims
}
