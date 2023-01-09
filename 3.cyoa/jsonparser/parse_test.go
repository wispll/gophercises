package jsonparser

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(t *testing.T){
    s := Unmarshal("gopher.json")
    r, _ := json.MarshalIndent(s, "", "\t")
    t.Log(string(r))
}
