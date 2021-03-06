// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://quorumsco.mit-license.org/license.txt for details.

package elastic

import (
	"encoding/json"
	"testing"
)

func TestTermFilter(t *testing.T) {
	f := NewTermFilter("user", "ki")
	f = f.Cache(true)
	f = f.CacheKey("MyTermFilter")
	f = f.FilterName("MyFilterName")
	data, err := json.Marshal(f.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"term":{"_cache":true,"_cache_key":"MyTermFilter","_name":"MyFilterName","user":"ki"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
