package fs

import (
	"reflect"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestMinFunction_Evaluate(t *testing.T) {
	type args struct {
		context *opt.FormulaContext
		args    []*opt.LogicalExpression
	}
	tests := []struct {
		name    string
		m       *MinFunction
		args    args
		want    *opt.Argument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MinFunction{}
			got, err := m.Evaluate(tt.args.context, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinFunction.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinFunction.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
