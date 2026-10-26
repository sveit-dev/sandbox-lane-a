package protected

import "testing"

func TestLoad(t *testing.T) {
	c, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if c.Currency != "EUR" || c.DefaultPageSize != 20 {
		t.Fatalf("unexpected config: %+v", c)
	}
}
