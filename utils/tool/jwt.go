// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package tool

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	signKey = []byte("rtsys/www.b5net.com/admin")
)

// CreateTokenJwt 生成token
func CreateTokenJwt(data any, exp int64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"data": data, "exp": exp, "iss": "b5net"}).SignedString(signKey)
}

func ParseTokenJwt(str string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token失效")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token失效")
	}
	if e := claims.Valid(); e != nil {
		return nil, errors.New("token失效")
	}
	return claims, nil
}
