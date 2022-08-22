package token

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kkodecaffeine/go-common/rest"
)

/*
 * Extract token from header
 *
 * @author kkodecaffeine@gmail.com
 */
func ExtractBearerToken(header string) (string, *rest.CustomError) {
	if header == "" {
		return "", &rest.CustomError{CodeDesc: &rest.TOKEN_MISSING, Message: ""}
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", &rest.CustomError{CodeDesc: &rest.TOKEN_INVALID_FORMAT, Message: ""}
	}

	return jwtToken[1], nil
}

/*
 * Extract values ("iss") from user claims. Then get secret key according to issuer type
 *
 * @author kkodecaffeine@gmail.com
 */
func ExtractClaims(jwtToken string) (string, *rest.CustomError) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return "", &rest.CustomError{CodeDesc: &rest.TOKEN_INVALID_FORMAT, Message: err.Error()}
	}

	var iss string
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		iss = fmt.Sprint(claims["iss"])
	}

	return iss, nil
}

func verifyToken(c *gin.Context, isBasic bool) (*jwt.Token, *rest.CustomError) {
	var accessSecret string
	if isBasic {
		accessSecret = os.Getenv("ACCESS_SECRET_INTERNAL")
	} else {
		accessSecret = os.Getenv("ACCESS_SECRET_MASTER")
	}

	SecretKey := "-----BEGIN CERTIFICATE-----\n" + accessSecret + "\n-----END CERTIFICATE-----"
	header := c.GetHeader("Authorization")

	jwtToken := strings.Split(header, " ")[1]

	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
	if er != nil {
		return nil, &rest.CustomError{CodeDesc: &rest.TOKEN_PARSING_ERROR, Message: ""}
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		switch {
		case strings.Contains(strings.ToLower(err.Error()), "token is expired"):
			return nil, &rest.CustomError{CodeDesc: &rest.TOKEN_EXPIRED, Message: ""}
		case strings.Contains(strings.ToLower(err.Error()), "crypto/rsa: verification error"):
			return nil, &rest.CustomError{CodeDesc: &rest.BAD_REQUEST, Message: "SECRET_KEY 가 일치하지 않습니다. 요청한 계정이 관리자 계정인지 혹은 일반 계정인지 확인해주세요."}
		default:
			return nil, &rest.CustomError{CodeDesc: &rest.TOKEN_PARSING_ERROR, Message: er.Error()}
		}
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	}

	return token, nil
}

/*
 * Perform token validation of a regular account, or an administrator account
 *
 * @author kkodecaffeine@gmail.com
 */
func TokenAuthMiddleware(c *gin.Context) {
	response := rest.NewApiResponse()

	bearerToken, err := ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		response.Error(err.CodeDesc, err.Message, nil)
		c.JSON(err.CodeDesc.HttpStatusCode, response)
		c.Abort()
		return
	}

	var isBasic bool = true
	iss, nil := ExtractClaims(bearerToken)
	if strings.Contains(strings.ToLower(iss), os.Getenv("REALM_MASTER")) {
		isBasic = false
	}

	token, err := verifyToken(c, isBasic)
	if err != nil {
		response.Error(err.CodeDesc, err.Message, nil)
		c.JSON(err.CodeDesc.HttpStatusCode, response)
		c.Abort()
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		response.Error(&rest.FAILED_INTERNAL_ERROR, "unable to parse claims", nil)
		c.JSON(rest.FAILED_INTERNAL_ERROR.HttpStatusCode, response)
		c.Abort()
		return
	}
	c.Next()
}
