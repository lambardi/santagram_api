package userdao

import "santagram_api/server/user"

type DAO interface {
	Store(user user.User) (*user.User, error)
	Authenticate(user user.User, password string) (bool, error)
}
