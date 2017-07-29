package reader

import "testing"

func TestLogID_Parse(t *testing.T) {
    stringId := "e16272e8-626d-48fb-bad1-ea23144bfeb4"
    id := LogID{}
    err := id.Parse(stringId)
    if err != nil {
        t.Errorf("Unexpected error during parse: %s", err.Error())
    }

    if res := id.String(); res != stringId {
        t.Errorf("Log ID did not round-trip correctely. Expected: %s; Actual: %s", stringId, res)
    }

}
