package bcrypt

import "golang.org/x/crypto/bcrypt"

func Generate(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func MustHash(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func Compare(newPassHashed, oldPassHashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(oldPassHashed), []byte(newPassHashed))

	return err == nil
}
