package model

type Usersc struct {
	Users    []User `yaml:"users"`
	UserDirsMap map[string]bool
}

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`
}
