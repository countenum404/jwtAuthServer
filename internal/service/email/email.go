package email

import (
	"go.uber.org/fx"
	"net/smtp"
	"os"
	"strings"
)

var Module = fx.Module("Warning", fx.Provide(fx.Annotate(NewWarningEmailService, fx.As(new(WarningService)))))

type WarningService interface {
	SendWarning(warning, subject, userEmail string) error
}

type WarningEmailService struct{}

func NewWarningEmailService(lc fx.Lifecycle) *WarningEmailService {
	return &WarningEmailService{}
}

func (w WarningEmailService) SendWarning(warning, subject, userEmail string) error {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_SENDER"), os.Getenv("APP_SMTP_KEY"), "smtp.gmail.com")
	builder := strings.Builder{}

	to := []string{userEmail}

	builder.WriteString("To:")
	builder.WriteString(userEmail)
	builder.WriteString("\r\n")

	builder.WriteString("Subject:")
	builder.WriteString(subject)
	builder.WriteString("\r\n")

	builder.WriteString(warning)
	builder.WriteString("\r\n")

	message := []byte(builder.String())

	return smtp.SendMail("smtp.gmail.com:587", auth, "countenum404@gmail.com", to, message)
}
