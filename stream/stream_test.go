package stream_test

import (
	"testing"
	"time"

	"github.com/chulabs/seer/stream"
)

func TestNew(t *testing.T) {
	s, err := stream.New("sales", 60, 10, 0, 1)

	if err != nil {
		t.Error("unexpected error in New:", err)
	}
	if s.Config.Min != 10 {
		t.Errorf("expected config min %v, but got %v", 10, s.Config.Min)
	}
}

func TestNewErrs(t *testing.T) {
	s, err := stream.New("y", 60, 0, 0, 0)

	if err == nil {
		t.Error("expected error but got nil")
	}
	if s != nil {
		t.Error("expected nil stream but got:", s)
	}
}

func TestStreamUpdate(t *testing.T) {
	vals := []float64{1, 2}
	times := []time.Time{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC)}

	s, err := stream.New("streamy", 86400, 0, 0, 0)
	if err != nil {
		t.Fatal("unexpected error in New:", err)
	}

	err = s.Update(vals, times)
	if err != nil {
		t.Fatal("unexpected error in Update:", err)
	}

	if s.Model.RCE.History[0] == 0 {
		t.Error("model was not updated")
	}
}

func TestStreamUpdateErrs(t *testing.T) {
	tt := []struct {
		name   string
		times  []time.Time
		values []float64
	}{
		{"mismatched lengths", []time.Time{time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC)}, []float64{1, 2}},
		{"wrong next time", []time.Time{time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC)}, []float64{1}},
		{
			"wrong intermediate time",
			[]time.Time{
				time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2016, 1, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2016, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			[]float64{2, 1, 3},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := stream.New("streamy", 86400, 0, 0, 0)
			if err != nil {
				t.Fatal("unexpected error in New:", err)
			}
			err = s.Update([]float64{1}, []time.Time{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)})
			if err != nil {
				t.Fatal("unexpected error in Update:", err)
			}

			err = s.Update(tc.values, tc.times)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
		})
	}
}
