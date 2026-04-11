package utils

import (
	"fmt"
	"net/smtp"

	"go.uber.org/zap"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
	SendMessageFromCustomer(subject string, body string) error
}

type smtpEmailService struct {
	config SMTPConfig
	log    *zap.Logger
}

func NewEmailService(config SMTPConfig, log *zap.Logger) EmailService {
	return &smtpEmailService{
		config: config,
		log:    log,
	}
}

func (s *smtpEmailService) SendEmail(to string, subject string, body string) error {
	if s.config.Host == "" {
		s.log.Info("SMTP Host not configured, skipping email send",
			zap.String("to", to),
			zap.String("subject", subject),
			zap.String("body_snippet", body[:min(len(body), 50)]),
		)
		return nil
	}

	auth := smtp.PlainAuth("", s.config.Email, s.config.Password, s.config.Host)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	err := smtp.SendMail(addr, auth, s.config.Email, []string{to}, msg)
	if err != nil {
		s.log.Error("failed to send email", zap.Error(err))
		return err
	}

	s.log.Info("email sent successfully", zap.String("to", to))
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *smtpEmailService) SendMessageFromCustomer(subject string, body string) error {
	to := s.config.Host
	if s.config.Host == "" {
		s.log.Info("SMTP Host not configured, skipping email send",
			zap.String("to", to),
			zap.String("subject", subject),
			zap.String("body_snippet", body[:min(len(body), 50)]),
		)
		return nil
	}

	auth := smtp.PlainAuth("", s.config.Email, s.config.Password, s.config.Host)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	addr := fmt.Sprintf("%s:%d", s.config.Email, s.config.Port)
	err := smtp.SendMail(addr, auth, s.config.Email, []string{to}, msg)
	if err != nil {
		s.log.Error("failed to send message", zap.Error(err))
		return err
	}

	s.log.Info("message sent successfully", zap.String("to", to))
	return nil
}