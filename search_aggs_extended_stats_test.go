// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://quorumsco.mit-license.org/license.txt for details.

package elastic

import (
	"encoding/json"
	"testing"
)

func TestExtendedStatsAggregation(t *testing.T) {
	agg := NewExtendedStatsAggregation().Field("grade")
	data, err := json.Marshal(agg.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"extended_stats":{"field":"grade"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestExtendedStatsAggregationWithFormat(t *testing.T) {
	agg := NewExtendedStatsAggregation().Field("grade").Format("000.0")
	data, err := json.Marshal(agg.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"extended_stats":{"field":"grade","format":"000.0"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
