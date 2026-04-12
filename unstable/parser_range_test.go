package unstable

import (
	"errors"
	"testing"
)

// Regression test for https://github.com/pelletier/go-toml/issues/1047:
// Parser.Range must use the real slice offset, not len(data)-len(slice).
func TestParser_Range_HighlightAfterComment(t *testing.T) {
	input := []byte("# comment\n= \"value\"")

	var p Parser
	p.Reset(input)
	for p.NextExpression() {
	}
	err := p.Error()
	if err == nil {
		t.Fatal("expected an error")
	}

	var perr *ParserError
	if !errors.As(err, &perr) {
		t.Fatalf("expected *ParserError, got %T", err)
	}

	r := p.Range(perr.Highlight)
	shape := p.Shape(r)

	if r.Offset != 10 {
		t.Errorf("Range offset: got %d, want 10", r.Offset)
	}
	if shape.Start.Line != 2 || shape.Start.Column != 1 {
		t.Errorf("position: got %d:%d, want 2:1", shape.Start.Line, shape.Start.Column)
	}
}
