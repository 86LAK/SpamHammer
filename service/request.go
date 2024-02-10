package service

import "time"

type Request struct {
	ID    string `json:"id"`
	Email struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
		Headers struct {
			Customer  string `json:"X-Customer-Id"`
			MessageId string `json:"X-Message-Id"`
			Seed      string `json:"X-SpamHammer-Fingerprint"`
		} `json:"headers"`
		Date time.Time `json:"date"`
	} `json:"email"`
}
