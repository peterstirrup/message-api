package usecases

// Messages enables access to a host of messaging repositories.
type Messages struct {
	WhatsApp WhatsApp
}

func New(w WhatsApp) *Messages {
	return &Messages{
		WhatsApp: w,
	}
}
