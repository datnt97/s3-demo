package util

import "time"

func Int(i int) *int {
	return &i
}

func Bool(i bool) *bool {
	return &i
}

func String(i string) *string {
	return &i
}

func Duration(i time.Duration) *time.Duration {
	return &i
}
