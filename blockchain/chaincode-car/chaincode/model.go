package chaincode

import "time"

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

type Car struct {
	CarID    string      `json:"carID"`
	Tires    CarTires    `json:"tires"`
	Body     CarBody     `json:"body"`
	Interior CarInterior `json:"interior"`
	Manu     CarManu     `json:"manu"`
	Store    CarStore    `json:"store"`
	Insure   *CarInsure  `json:"insure"`
	Maint    *CarMaint   `json:"maint"`
	Owner    string      `json:"owner"`
	Record   *CarRecord  `json:"record"`
}

type CarTires struct {
	Time     time.Time `json:"time"`
	Width    float32   `json:"width"`
	Radius   float32   `json:"radius"`
	Workshop string    `json:"workshop"`
	TxID     string    `json:"txID"`
}

type CarBody struct {
	Time     time.Time `json:"time"`
	Material string    `json:"material"`
	Weitght  float32   `json:"weitght"`
	Color    string    `json:"color"`
	Workshop string    `json:"workshop"`
	TxID     string    `json:"txID"`
}

type CarInterior struct {
	Time     time.Time `json:"time"`
	Material string    `json:"material"`
	Weitght  float32   `json:"weitght"`
	Color    string    `json:"color"`
	Workshop string    `json:"workshop"`
	TxID     string    `json:"txID"`
}

type CarManu struct {
	Time     time.Time `json:"time"`
	Workshop string    `json:"workshop"`
	TxID     string    `json:"txID"`
}

type CarStore struct {
	Time  time.Time `json:"time"`
	Store string    `json:"store"`
	Cost  float32   `json:"cost"`
	Owner string    `json:"owner"`
	TxID  string    `json:"txID"`
}

type CarInsure struct {
	Insures []Insure `json:"insures"`
}

type Insure struct {
	Name      string    `json:"name"`
	BeginTime time.Time `json:"beginTime"`
	EndTime   time.Time `json:"endTime"`
	Cost      float32   `json:"cost"`
	TxID      string    `json:"txID"`
}

type CarMaint struct {
	Maints []Maint `json:"maints"`
}

type Maint struct {
	Time   time.Time `json:"time"`
	Part   string    `json:"part"`
	Extent string    `json:"extent"`
	Cost   float32   `json:"cost"`
	TxID   string    `json:"txID"`
}

type CarRecord struct {
	Records []Record `json:"record"`
}

type Record struct {
	OldUser string    `json:"oldUser"`
	NewUser string    `json:"newUser"`
	Cost    float32   `json:"cost"`
	Time    time.Time `json:"time"`
	TxID    string    `json:"txID"`
}
