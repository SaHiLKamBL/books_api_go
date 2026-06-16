package util

import (
    "golang.org/x/crypto/bcrypt"
	"fmt"
	
)

func HashPassword(password string)(string,error){
	
	hashed,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)  // 10 is default cost means how much it will iterate to hash it well
     
	if err !=nil{
		return "",fmt.Errorf("error hashing password: %w",err)
	}

	return string(hashed),nil

}

func CheckPassword(password string,hashedpassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword),[]byte(password))  //compare hased and og password
}