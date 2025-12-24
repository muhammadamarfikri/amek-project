package internal

import "fmt"

func (uu *UserUsecase) Registration() {
	fmt.Println("test usecase")
	uu.UserRepo.CreateUser()
}
