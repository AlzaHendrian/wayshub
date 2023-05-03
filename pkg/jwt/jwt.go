package jwtToken

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = "SECRET_KEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Parameter pertama menunjukkan metode yang digunakan untuk menandatangani token JWT dengan kunci rahasia mengunakan metode sha
	webtoken, err := token.SignedString([]byte(SecretKey))     //ditantdatangani oleh tokensignedtring agar token dapat di verifikasi,
	if err != nil {                                            //jadi signedtring itu mengembalikan token yang sudah menjadi string
		return "", err //payload (klaim) yang ingin dimasukkan ke dalam token JWT. Payload ini biasanya berisi informasi pengguna atau aplikasi yang meminta token, seperti username, roles, atau hak akses lainnya
	}

	return webtoken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) { //Fungsi ini akan mengembalikan token JWT jika berhasil diverifikasi, dan akan menangani kesalahan jika terjadi.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) { // digunakan untuk mendekode token JWT dan mengambil nilai claims yang terdapat di dalamnya.
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
