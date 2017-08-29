package s3

import (
	"testing"
)

func TestParsePathSuccess(t *testing.T) {
	var validCases = []struct{
		in string
		want Path
	}{
		{"s3://bucket/key", Path{Bucket: "bucket", Key: "/key"}},
		{"s3://bucket/key1/key2", Path{Bucket: "bucket", Key: "/key1/key2"}},
	}

	for _, v := range validCases {
		got := ParsePath(v.in)
		if got != v.want {
			t.Errorf("Reverse(%q) == %q, want %q", v.in, got, v.want)
		}
	}
}
