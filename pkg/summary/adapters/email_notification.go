package adapters

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"time"

	"github.com/eduaravila/stori-challenge/pkg/summary/domain"
	"golang.org/x/exp/slog"
)

type EmailNotification struct {
}

type MonthTransactions struct {
	Month                     string
	TotalNumberOfTransactions string
}

func NewMonthTransactions(month int, transactions int) MonthTransactions {
	return MonthTransactions{
		Month:                     time.Month(month).String(),
		TotalNumberOfTransactions: fmt.Sprintf("%d", transactions),
	}
}

func sliceToMonthTransactions(slice []int) []MonthTransactions {
	var monthTransactions []MonthTransactions

	for index, value := range slice {
		if value < 1 {
			continue
		}

		monthTransactions = append(monthTransactions, NewMonthTransactions(index, value))
	}

	return monthTransactions
}

type EmailSummaryData struct {
	UserName                     string
	AverageCredit                float64
	AverageDebit                 float64
	Total                        float64
	NumberOfTransactionsPerMonth []MonthTransactions
	SenderEmail                  string
	SenderPhone                  string
}

func NewEmailNotification() EmailNotification {
	return EmailNotification{}
}

func (e EmailNotification) Send(user *domain.User, transactionSummary *domain.TransactionsSummary) error {
	// SMTP serverdetails
	smtpHost := os.Getenv("SUMMARY_SMTP_HOST")
	smtpPort := os.Getenv("SUMMARY_SMTP_PORT")
	senderEmail := os.Getenv("SUMMARY_SUPPORT_EMAIL")
	senderUser := os.Getenv("SUMMARY_SUPPORT_USER")
	senderPhone := os.Getenv("SUMMARY_SUPPORT_PHONE")
	senderPassword := os.Getenv("SUMMARY_SUPPORT_EMAIL_PASSWORD")
	recipientEmail := user.Email()

	// Email content
	subject := user.Name() + ", here's your year wrapped!"

	buff := bytes.NewBuffer([]byte{})

	tmplt := template.Must(template.ParseFiles("assets/email_summary_template.html"))

	emailSummaryData := EmailSummaryData{
		UserName:                     user.Name(),
		AverageCredit:                transactionSummary.AvarageCredit(),
		AverageDebit:                 transactionSummary.AvarageDebit(),
		Total:                        transactionSummary.Total(),
		NumberOfTransactionsPerMonth: sliceToMonthTransactions(transactionSummary.NumberOfTransactionsPerMonth()),
		SenderEmail:                  senderEmail,
		SenderPhone:                  senderPhone,
	}

	err := tmplt.Execute(buff, emailSummaryData)

	if err != nil {
		return err
	}

	tmplt.Execute(buff, emailSummaryData)

	htmlContent := buff.String()

	// Compose the email message
	message := "From: " + senderEmail + "\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		htmlContent

	// Set up authentication
	auth := smtp.PlainAuth("", senderUser, senderPassword, smtpHost)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, []byte(message))

	fmt.Println(err, recipientEmail, senderEmail, senderPassword, smtpHost, smtpPort, auth)

	if err != nil {
		slog.Error(err.Error())

		return err
	}

	return nil
}
