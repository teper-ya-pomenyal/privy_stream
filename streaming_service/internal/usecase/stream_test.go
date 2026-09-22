package usecase

import (
	"testing"
	"time"
)

func TestIsAdult(t *testing.T) {
	date := func(s string) time.Time {
		d, err := time.Parse(time.DateOnly, s)
		if err != nil {
			t.Fatal(err)
		}
		return d
	}

	tests := []struct {
		name  string
		birth string
		now   string
		want  bool
	}{
		{"день в день 18 лет", "2008-03-15", "2026-03-15", true},
		{"за день до 18 лет", "2008-03-15", "2026-03-14", false},
		{"родился в декабре, сейчас январь", "2008-12-20", "2026-01-10", false},
		{"давно взрослый", "1990-01-01", "2026-09-22", true},
		{"29 февраля, невисокосный год", "2008-02-29", "2026-02-28", false},
		{"29 февраля, 1 марта", "2008-02-29", "2026-03-01", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdult(date(tt.birth), date(tt.now)); got != tt.want {
				t.Errorf("isAdult(%s, %s) = %v, want %v", tt.birth, tt.now, got, tt.want)
			}
		})
	}
}
