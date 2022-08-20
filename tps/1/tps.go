package tps

import (
	"os"
)

const data = "I believe you have my stapler?"

func WriteReportFile(path string) {
	err := os.WriteFile(path, []byte(data), 0o666)
	if err != nil {
		panic("Sounds like somebody's got a case of the Mondays!")
	}
}
