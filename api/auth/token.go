package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/responses"
	"github.com/nitinda/microservice-change-log/api/utils/console"
	"github.com/nitinda/microservice-change-log/config"
	"github.com/nitinda/microservice-change-log/logger"
)

// CreateToken method
func CreateToken(teamName string, clientSecret string) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["tokenOwner"] = teamName
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString(config.SECRETKEY)

	// fmt.Printf("%v %v", ss, err)
	// fmt.Sprintf("%s%s", config.SECRETKEY, "dfasdfa")

	return ss, err
}

// ValidateToken method
func ValidateToken(rw http.ResponseWriter, r *http.Request) error {

	tokenString := ExtractToken(r)

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return config.SECRETKEY, nil
	})

	if err != nil {
		logger.Error.Println(err)
		switch e := err.(type) {
		case *jwtgo.ValidationError:
			if e.Inner == nil {
				return err
			}
			return e.Inner
		}
		return err
	}

	var tokenOwner string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["foo"], claims["nbf"])
		console.ToJSON(claims)

		tokenOwner = claims["tokenOwner"].(string)
	}

	err = ValidateClaimOwner(rw, r, tokenOwner)
	if err != nil {
		logger.Error.Println("Invalid Token owner")
		return err
	}

	return nil
}

// ExtractToken method
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// ValidateClaimOwner method
func ValidateClaimOwner(rw http.ResponseWriter, r *http.Request, tokenOwner string) error {

	var bodyBytes []byte
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return err
	}

	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	changeLog := models.ChangeLog{}
	err = json.Unmarshal(bodyBytes, &changeLog)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return err
	}

	if changeLog.ServiceTeamName != tokenOwner {

		fmt.Println("-=-===-=-==-")

		logger.Error.Println("Invalid Token owner")
		// responses.ToJSON(rw, http.StatusUnauthorized, map[string]string{"unauthorized": "Invalid Token owner"})
		return errors.New("Invalid Token owner")
	}
	return nil
}
