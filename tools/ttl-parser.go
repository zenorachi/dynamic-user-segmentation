package tools

import "time"

func ParseTTL(ttl string) (time.Duration, error) {
	if ttl == "" {
		return 0, nil
	}

	ttlDuration, err := time.ParseDuration(ttl)
	if err != nil {
		return 0, err
	}

	return ttlDuration, nil
}
