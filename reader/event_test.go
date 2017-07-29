package reader

import (
    "testing"
)

func TestEventID_RoundTrip(t *testing.T) {
    idString := "cca637fa6cab3e8f5399f8c7bf259e0b7a5b003a881f00619e6f420e135ed74b"
    id := EventID{}
    err := id.Parse(idString)
    if err != nil {
        t.Errorf("Unexpected error during parse: %s", err.Error())
    }

    if res := id.String(); res != idString {
        t.Errorf("Did not parse ID correctly. Expected: %s; Actual: %s", idString, res)
    }
}
