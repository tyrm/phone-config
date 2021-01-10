package models

type PhoneMeta struct {
	Lines int
}

func GetPhoneMeta(vendor, model string) *PhoneMeta {
	meta := PhoneMeta{}

	switch vendor {
	case "fanvil":
		switch model {
		case "x5u":
			meta.Lines = 16
		default:
			return nil
		}
	default:
		return nil
	}

	return &meta
}