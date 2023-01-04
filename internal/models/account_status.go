package models

type AccountStatus int32

const (
	Pending AccountStatus = iota
	Active
	Disabled
)

func (fs AccountStatus) String() string {
	switch fs {
	case Active:
		return "ACTIVE"
	case Disabled:
		return "DESABLE"
	}

	return "unkown"
}

func AccountStatusString(status string) string {
	switch status {
	case "ACTIVE":
		return "ACTIVE"
	case "DESABLE":
		return "DESABLE"
	}

	return "ACTIVE"
}
