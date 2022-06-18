package main

import "fmt"

// Factory Notification
type IntNotificationFactory interface {
	SendNotification()
	GetSender() IntSenderFactory
}

// Factory Sender
type IntSenderFactory interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS Notifications
type SMSNotification struct {
}

// Methods of SMSNotification
func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification SMS")
}
func (SMSNotification) GetSender() IntSenderFactory {
	return SMSNotificationSender{}
}

// SMS Notification Sender
type SMSNotificationSender struct {
}

// Methods of SMSNotificationSender
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Email Notification
type EmailNotification struct {
}

// Methods of Email Notification
func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification Email")
}
func (EmailNotification) GetSender() IntSenderFactory {
	return EmailNotificationSender{}
}

// Email Notification Sender
type EmailNotificationSender struct {
}

// Methods of EmailNotificationSender
func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "Gmail"
}
