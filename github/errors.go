package github

import (
	"log"
	"strings"

	"github.com/google/go-github/v47/github"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func shouldRetryError(err error) bool {
	if _, ok := err.(*github.RateLimitError); ok {
		log.Printf("[WARN] Received Rate Limit Error")
		return true
	}
	if strings.Contains(err.Error(), "secondary rate limit") {
		return true
	}
	return false
}

// function which returns an ErrorPredicate for Github API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		if err != nil {
			for _, item := range notFoundErrors {
				if strings.Contains(err.Error(), item) {
					return true
				}
			}
		}
		return false
	}
}
