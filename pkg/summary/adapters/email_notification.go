package adapters

type EmailNotification struct {
}

func NewEmailNotification() EmailNotification {
	return EmailNotification{}
}

func (e EmailNotification) Send(email string, message string) error {
	return nil
}
