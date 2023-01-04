package types

type Contact struct {
	Firstname string
	Lastname  string
	Code      int8
	Phone     int64
	Email     string
	Address   Address
}

func init() {
}
