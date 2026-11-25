package bitbucket

import (
	"encoding/json"
	"testing"
	"time"
)

// Unset type and timestamps must not be sent: Bitbucket rejects them as
// "extra keys not allowed" on write requests (#41). See the postprocess
// target in GNUmakefile, which keeps these tags across regenerations.
func TestProjectOmitsUnsetReadOnlyFields(t *testing.T) {
	b, err := json.Marshal(Project{Name: "test", Key: "TTEST", IsPrivate: true})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(b), `{"is_private":true,"key":"TTEST","name":"test"}`; !jsonEqual(t, got, want) {
		t.Fatalf("got %s, want %s", got, want)
	}
}

// Set values are still sent, and responses still decode.
func TestProjectKeepsSetFields(t *testing.T) {
	created := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	b, err := json.Marshal(Project{Type_: "project", Key: "K", CreatedOn: created})
	if err != nil {
		t.Fatal(err)
	}
	var out Project
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if out.Type_ != "project" || !out.CreatedOn.Equal(created) || out.Key != "K" {
		t.Fatalf("round trip lost fields: %s -> %+v", b, out)
	}
}

// Bools that Bitbucket requires are still sent when false (#42).
func TestRequiredBoolsSentWhenFalse(t *testing.T) {
	b, err := json.Marshal(DeploymentVariable{Key: "K", Value: "v"})
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	if v, ok := m["secured"]; !ok || v != false {
		t.Fatalf("secured not sent as false: %s", b)
	}
	if _, ok := m["type"]; ok {
		t.Fatalf("unset type sent: %s", b)
	}
}

func jsonEqual(t *testing.T, a, b string) bool {
	t.Helper()
	var x, y any
	if err := json.Unmarshal([]byte(a), &x); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(b), &y); err != nil {
		t.Fatal(err)
	}
	xb, _ := json.Marshal(x)
	yb, _ := json.Marshal(y)
	return string(xb) == string(yb)
}
