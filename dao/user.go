package dao

import (
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/scrypt"
	"log"
	"multiband/model"
)

func CheckUserName(username string) (*model.User, error) {
	var user model.User
	model.Db.Where("username = ?", username).First(&user)
	return &user, nil
}

func SaveUser(data *model.User) error {

	data.Password = ScryotPW(data.Password)
	err := model.Db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func ScryotPW(password string) string {
	const Kenlen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 12}
	HashPW, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, Kenlen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPW)
	return fpw

}

func CheckUsernameAndPassword(u string, p string) (*model.User, error) {
	var user model.User
	password := ScryotPW(p)
	model.Db.Where("username = ?", u).First(&user)
	if user.Id == 0 {
		return nil, errors.New("user not exist")
	}
	if password != user.Password {
		return nil, errors.New("password wrong")
	}
	return &user, nil
}
