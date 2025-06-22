package apikey

import (
	"errors"
	"time"
)

func ValidateTimestamps(created, updated time.Time) error {
	const maxFutureDrift = 10 * time.Second
	now := time.Now().UTC()

	if created.After(now.Add(maxFutureDrift)) {
		return errors.New("created_at cannot be in the future")
	}
	if updated.After(now.Add(maxFutureDrift)) {
		return errors.New("updated_at cannot be in the future")
	}
	if updated.Before(created) {
		return errors.New("updated_at cannot be before created_at")
	}
	return nil
}
