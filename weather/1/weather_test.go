package weather_test

import (
	"testing"

	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestFormatURL_IncludesLocation(t *testing.T) {
	t.Parallel()
	location := "Nowhere"
	want := "https://weather.example.com/?q=Nowhere"
	got := weather.FormatURL(location)
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse(t *testing.T) {
	t.Parallel()
	data := []byte(`{"weather":[{"main":"Clouds"}]}`)
	want := weather.Forecast{
		Summary: "Clouds",
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
