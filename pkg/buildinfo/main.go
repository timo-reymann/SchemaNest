package buildinfo

import "time"

// GitSha of the build
var GitSha = "unknown"

// Version contains the latest version tag or question mark for source builds
var Version = "0.0.0"

// BuildTime contains the build time or question mark for source builds
var BuildTime = "2025-04-25_10:00:00"

var BuildTimeParsed time.Time

var BuildTimeRFC1123 string

func init() {
	parsed, err := time.Parse("2006-01-02_15:04:05", BuildTime)
	if err != nil {
		panic(err)
	}
	BuildTimeParsed = parsed
	BuildTimeRFC1123 = parsed.Format(time.RFC1123)
}
