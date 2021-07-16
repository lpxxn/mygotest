package strategy

import "fmt"

type UserType int

const (
	UserTypeNormal UserType = iota
	UserTypeVIP
	UserTypeExtra
)

type User struct {
	Name string
	Age  int32
	Type UserType
}

type VIPCenter struct {
	providers map[UserType]ServiceProvider
}

func (v *VIPCenter) addProviders(t UserType, s ServiceProvider) {
	v.providers[t] = s
}

func (v *VIPCenter) ServiceUser(user *User) {
	v.providers[user.Type].Service(user)
}

type ServiceProvider interface {
	Service(user *User) error
}

type normalUserServiceProvider struct {
}

func (n *normalUserServiceProvider) Service(user *User) error {
	fmt.Println("normal user")
	return nil
}

type vipUserServiceProvider struct {
}

func (v *vipUserServiceProvider) Service(user *User) error {
	fmt.Println("vip user")
	return nil
}

type extraUserServiceProvider struct {
}

func (v *extraUserServiceProvider) Service(user *User) error {
	fmt.Println("extra user")
	return nil
}
