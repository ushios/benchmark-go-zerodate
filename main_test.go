package main

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkZeroHourUsingFmt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		now := time.Now()
		st := fmt.Sprintf("%s 00:00:00 %s", now.Format("2006-01-02"), now.Format("-0700"))
		day, _ := time.Parse("2006-01-02 15:04:05 -0700", st)

		if day.Day() != now.Day() {
			b.Errorf("day not matched")
		}
	}
}

func BenchmarkZeroHourUsingFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		now := time.Now()
		st := now.Format("2006-01-02 00:00:00 -0700 MST")
		day, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", st)
		if day.Day() != now.Day() {
			b.Errorf("day not matched")
		}
	}
}

func BenchmarkZeroHourUsingAdd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		now := time.Now()
		nanosecond := time.Duration(now.Nanosecond())
		second := time.Duration(now.Second())
		minute := time.Duration(now.Minute())
		hour := time.Duration(now.Hour())
		dur := -1 * (nanosecond + second*time.Second + minute*time.Minute + hour*time.Hour)

		day := now.Add(dur)

		if day.Day() != now.Day() {
			b.Errorf("day not matched")
		}
	}
}

func BenchmarkZerHourUnixTime(b *testing.B) {

	for i := 0; i < b.N; i++ {
		now := time.Now()
		ut := now.Unix()
		_, offset := now.Zone()
		day := time.Unix((ut/86400)*86400-int64(offset), 0)

		if day.Day() != now.Day() {
			b.Errorf("day not matched")
		}
	}
}
