package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la función que permite encriptar el password*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
