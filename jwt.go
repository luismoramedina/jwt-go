package main

import "fmt"
import "github.com/dgrijalva/jwt-go"
import "github.com/dgrijalva/jwt-go/test"

func main() {

   // External token
/*
   rsaPublicKey := test.LoadRSAPublicKeyFromDisk("resources/sample_key_mine.pub")
   var tokenString = "eyJraWQiOiJraWQiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ0b20iLCJuYmYiOjE0OTUxOTQzMzAsImlzcyI6Imlzc3VlciIsImV4cCI6MTUyNjczMDMzMCwiaWF0IjoxNDk1MTk0MzMwLCJqdGkiOiJqd3RJZCJ9.Rz4gUJmD2GUqCaDjyvqbl5eOTPAvg0vK7AYjnghR9tRQmL95Csao3mC_hoaC8S3ITOBLWqjk8N6IzbexCA-rBvDWzlfrPw3c_GM-3FK-IPbzK8LfzndaTaO-6lzUsWW9PKOcIrV9r0DyzjkxQtEM1TXWrfGDEsR1kqAGH6hHvP4"
*/

   // Generating the token
   rsaPublicKey := test.LoadRSAPublicKeyFromDisk("resources/sample_key.pub")

   claims := &jwt.StandardClaims {
      ExpiresAt: 2221039552,//year 2040
//       ExpiresAt: 1500, //Expired
      Issuer:    "test",
      Subject:    "bob",
   }

   rsaPrivateKey := test.LoadRSAPrivateKeyFromDisk("resources/sample_key")
   var tokenString = MakeSampleToken(claims, rsaPrivateKey)
   fmt.Printf("The token ------> %+v\n", tokenString)

   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
       return rsaPublicKey, nil
   })

   if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      fmt.Printf("----> GOOD TOKEN! for %s", claims["sub"])
   } else {
      fmt.Println("Errror", err)
   }
}

func MakeSampleToken(c jwt.Claims, key interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	token.Header["kid"] = "istio-kid"
	s, e := token.SignedString(key)

	if e != nil {
		panic(e.Error())
	}

	return s
}