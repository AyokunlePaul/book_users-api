package utils

const (
	StatusActive            = "active"
	StatusSuspended         = "suspended"
	StatusEmailNotConfirmed = "pending"
)

func IsNotValidStatus(status string) bool {
	return status != StatusActive && status != StatusSuspended && status != StatusEmailNotConfirmed
}
