package enum

type UserType string

const (
	GUEST      UserType = "GUEST"
	CUSTOMER   UserType = "CUSTOMER"
	SPECIALIST UserType = "SPECIALIST"
)

func (e UserType) IsValidate() bool {
	switch e {
	case GUEST, CUSTOMER, SPECIALIST:
		return true
	}

	return false
}
