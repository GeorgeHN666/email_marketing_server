package token

import "time"

type Payload struct {
	User      string    `json:"user"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
