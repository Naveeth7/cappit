package apikey

import (
	"net/http"
	neturl "net/url"
	"strings"
	"time"

	"github.com/cappit/internal/util"
)

func ValidateTenantInput(name, baseURL string, created, updated time.Time) error {
	if err := validateName(name); err != nil {
		return err
	}
	if err := validateBaseURL(baseURL); err != nil {
		return err
	}
	if err := validateTimestamps(created, updated); err != nil {
		return err
	}
	return nil
}

func validateTimestamps(created, updated time.Time) error {
	const maxFutureDrift = 10 * time.Second
	now := time.Now().UTC()

	if created.After(now.Add(maxFutureDrift)) {
		return util.NewErrorResponse("created_at cannot be in the future", http.StatusBadRequest)
	}
	if updated.After(now.Add(maxFutureDrift)) {
		return util.NewErrorResponse("updated_at cannot be in the future", http.StatusBadRequest)
	}
	if updated.Before(created) {
		return util.NewErrorResponse("updated_at cannot be before created_at", http.StatusBadRequest)
	}
	return nil
}

func validateName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return util.NewErrorResponse("name is required", http.StatusNotFound)
	}
	if len(name) < 3 || len(name) > 50 {
		return util.NewErrorResponse("name must be between 3 and 50 characters", http.StatusBadRequest)
	}
	if strings.Contains(name, "\u0000") {
		return util.NewErrorResponse("name cannot contain null characters", http.StatusBadRequest)
	}
	return nil
}

func validateBaseURL(url string) error {
	url = strings.TrimSpace(url)

	if url == "" {
		return util.NewErrorResponse("base_url is required", http.StatusNotFound)
	}
	if strings.ContainsRune(url, '\u0000') {
		return util.NewErrorResponse("base_url cannot contain null characters", http.StatusBadRequest)
	}

	parsed, err := neturl.ParseRequestURI(url)
	if err != nil {
		return util.NewErrorResponse("base_url must be a valid URL", http.StatusBadRequest)
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return util.NewErrorResponse("base_url must start with http or https", http.StatusBadRequest)
	}

	return nil
}
