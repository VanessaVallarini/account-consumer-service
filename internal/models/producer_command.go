package models

type ProducerCommand string

const (
	Create     ProducerCommand = "C"
	GetById    ProducerCommand = "GBI"
	GetByEmail ProducerCommand = "GBE"
	GetByPhone ProducerCommand = "GBP"
	List       ProducerCommand = "L"
	Update     ProducerCommand = "U"
	Delete     ProducerCommand = "D"
)

func (fs ProducerCommand) String() string {
	switch fs {
	case Create:
		return "C"
	case GetById:
		return "GBI"
	case GetByEmail:
		return "GBE"
	case GetByPhone:
		return "GBP"
	case List:
		return "L"
	case Update:
		return "U"
	case Delete:
		return "D"
	}

	return "unkown"
}
