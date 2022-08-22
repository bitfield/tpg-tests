package weather

import (
	"encoding/json"
	"fmt"
)

type Forecast struct {
	Summary string
}

func ParseResponse(data []byte) (Forecast, error) {
	var resp struct {
		Weather []struct {
			Main string
		}
	}
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Forecast{}, fmt.Errorf("invalid API response %q: %w", data, err)
	}
	forecast := Forecast{
		Summary: resp.Weather[0].Main,
	}
	return forecast, nil
}

func FormatURL(location string) string {
	return fmt.Sprintf("https://weather.example.com/?q=%s", location)
}
