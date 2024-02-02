package weather_test

import (
	"testing"

	"github.com/bitfield/weather"
)

func TestFormatURL_IncludesLocation(t *testing.T) {
	t.Parallel()
	location := "Nowhere"
	want := "https://weather.example.com/?q=Nowhere"
	got := weather.FormatURL(location)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestParseResponse_GivesExpectedForecast(t *testing.T) {
	t.Parallel()
	data := []byte(`{"weather":[{"main":"Clouds"}]}`)
	want := weather.Forecast{
		Summary: "Clouds",
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}
