package constant

type Role struct {
	Customer string
	Admin    string
}

var RoleUser Role = Role{
	Customer: "Customer",
	Admin:    "Admin",
}
