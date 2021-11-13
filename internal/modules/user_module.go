package modules

import "github.com/ceit-ssc/nc_backend/pkg/repository"


type UserModule struct{
	UserRepo repository.UserRepository
}
//return user ID
func (u *UserModule) RegisterNewUser()(int,error){

	return -1, nil
}

