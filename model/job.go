// Package main provides ...
package model

import "time"

type Job struct {
	Title      string    `json:"title"`
	PartnerID  int       `json:"partner_id"`
	Status     string    `json:"status"`
	CategoryID int       `json:"category_id"`
	ExpiresAt  time.Time `json:"expires_at"`
}
