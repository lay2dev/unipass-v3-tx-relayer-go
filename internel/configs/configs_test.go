package configs

import (
	"encoding/json"
	"testing"
)

func TestDecode(t *testing.T) {
	conf, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}

	a, _ := json.Marshal(conf)
	t.Error(string(a))
}
