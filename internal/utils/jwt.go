package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

const (
	UserRoleVolunteer = "volunteer"
	userIDHeaderName  = "user_id"
	bearerPrefix      = "bearer"
)

const JWTsecret = "eyzzzz*iOiTWak12121"

type PartnerClaims struct {
	UserRole string `json:"user_role"`
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}
type JWTConfig struct {
	Secret     string
	ExtendTime time.Duration
	Method     jwt.SigningMethod
}

func GenerateToken(uid, role string) (string, error) {

	claims := PartnerClaims{
		role,
		uid,
		jwt.StandardClaims{
			Issuer:    "miomi-core",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(JWTsecret)
	tok, err := token.SignedString(secret)

	tok = fmt.Sprintf("bearer %s", tok)

	return tok, err

}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetUserID(c *gin.Context) (int64, error) {
	return fetchInt64JWTHeader(c, userIDHeaderName)
}
func parseJWTToken(jwtKey, secret string, method jwt.SigningMethod) (jwt.MapClaims, error) {
	if jwtKey == "" {
		return nil, errors.New("jwt not valid")
	}
	parsedToken, err := jwt.Parse(jwtKey, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if token.Method != method {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}

func ValidateStandardJWTClaims(claims jwt.MapClaims) error {
	if _, ok := claims["exp"].(float64); !ok {
		return errors.New("exp")
	}
	if _, ok := claims["iat"].(float64); !ok {
		return errors.New("iat")
	}
	if _, ok := claims["nbf"].(float64); !ok {
		return errors.New("nbf")
	}
	return nil
}

func fetchInt64JWTHeader(c *gin.Context, headerName string) (int64, error) {
	headerValue := c.GetHeader("Authorization")
	if headerValue == "" {
		return 0, errors.New(fmt.Sprintf("no %s found in request headers", headerName))
	}

	jwtConf := JWTConfig{Method: jwt.SigningMethodHS256, Secret: JWTsecret, ExtendTime: time.Hour * 48}
	clames, err := ExtractJWTClaims(headerValue, jwtConf)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to parse %s header", headerName))
	}
	var vv string
	if _, ok := clames[headerName]; !ok {
		return 0, errors.New(fmt.Sprintf("no %s found in request headers", headerName))
	}
	vv = clames[headerName].(string)
	int64Value, err := strconv.ParseInt(vv, 10, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to parse %s header", headerName))
	}
	return int64Value, nil
}

func ExtractJWTClaims(jwtToken string, conf JWTConfig) (jwt.MapClaims, error) {
	jwtToken = ExtractTokenFromBearer(jwtToken)
	claims, err := parseJWTToken(jwtToken, conf.Secret, conf.Method)
	if err != nil {
		return claims, err
	}
	if err := ValidateStandardJWTClaims(claims); err != nil {
		return claims, err
	}
	return claims, nil
}

func ExtractTokenFromBearer(bearer string) string {
	authHeaderParts := strings.Split(bearer, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != bearerPrefix {
		return ""
	}
	return authHeaderParts[1]
}
