package channel

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Channel struct {
	Major string
	Minor string
}

func (c *Channel) MajorVersion() int {
	num, _ := strconv.Atoi(c.Major)
	return num
}

func (c *Channel) MinorVersion() int {
	if c.Minor == "x" {
		return -1
	}

	num, _ := strconv.Atoi(c.Minor)
	return num
}

func emptyWhenX(s string) string {
	if s == "x" {
		return ""
	}
	return s
}

func validateChannelPart(part string) error {
	if part == "x" {
		return nil
	}

	_, err := strconv.Atoi(part)
	if err != nil {
		return errors.New("invalid channel part: " + part)
	}
	return nil
}

func Parse(channel string) (*Channel, error) {
	parts := strings.Split(channel, ".")
	if len(parts) < 2 || len(parts) > 3 {
		return nil, fmt.Errorf("invalid channel format: %s", channel)
	}

	if len(parts) == 3 && parts[2] != "x" {
		return nil, fmt.Errorf("invalid channel format: %s", channel)
	}

	if err := validateChannelPart(parts[0]); err != nil || parts[0] == "x" {
		return nil, fmt.Errorf("invalid major version for channel: %s", err)
	}

	if err := validateChannelPart(parts[1]); err != nil {
		return nil, fmt.Errorf("invalid minor version for channel: %s", err)
	}

	return &Channel{
		Major: emptyWhenX(parts[0]),
		Minor: emptyWhenX(parts[1]),
	}, nil
}
