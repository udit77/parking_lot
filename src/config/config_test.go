package config

import (
	"testing"
	"os"
	"encoding/json"
)

func TestReadConfig(t *testing.T){
	environ := os.Getenv("PARKINGLOT")
	if environ == "" {
		environ = "development"
	}

	filename := "../../files/etc/parking_lot/parking_lot."+environ+".json"
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("err: [TestReadConfig] error occured in reading config file %v", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config{})
	if err != nil {
		t.Errorf("err: [TestReadConfig] error occured in decoding config file %v", err)
	}
}
