package main

import (
	"testing"
)

func TestEqual(t *testing.T) {
	result := score_gotoh_str("test", "test")
	if result  != 4 {
		t.Fatalf("expected 4, got %v", result)
	}
}


func TestGap(t *testing.T) {
	result := score_gotoh_str("test", "teast")
	if result  != -1 {
		t.Fatalf("expected 4, got %v", result)
	}
}

func TestGapExtend(t *testing.T) {
	result := score_gotoh_str("test", "teaist")
	if result  != -2 {
		t.Fatalf("expected 4, got %v", result)
	}
}
