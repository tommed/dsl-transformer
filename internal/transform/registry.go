package transform

import (
	"fmt"
	"github.com/tommed/dsl-transformer/internal/model"
)

type Registry struct {
	ops map[string]Operator
}

func NewRegistry() *Registry {
	return &Registry{ops: map[string]Operator{}}
}

func (r *Registry) Register(op Operator) {
	r.ops[op.Name()] = op
}

func (r *Registry) Apply(ctx *ExecutionContext, input map[string]interface{}, instr model.Instruction) bool {
	op, ok := r.ops[instr.Op]
	if !ok {
		return ctx.Handle(fmt.Errorf("unknown op: %q", instr.Op))
	}
	err := op.Apply(ctx, input, instr)
	if err != nil {
		return ctx.Handle(err)
	}
	return true
}
