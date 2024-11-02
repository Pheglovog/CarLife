package model

const (
	Undefined         = "undefined"
	ComponentSupplier = "componentSupplier"
	Manufacturer      = "manufacturer"
	Store             = "store"
	Insurer           = "insurer"
	Maintenancer      = "maintenancer"
	Consumer          = "consumer"
)

type User struct {
	UserID   string   `json:"userID"`
	UserType string   `json:"userType"`
	Password string   `json:"password"`
	CarList  []string `json:"carList"`
}
