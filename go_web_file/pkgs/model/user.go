package model
type Usersc struct {
	Users []User `yaml:"users"`
}
type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}