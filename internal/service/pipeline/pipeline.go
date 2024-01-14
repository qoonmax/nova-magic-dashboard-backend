package pipeline

import (
	"fmt"
	"magic-dashboard-backend/internal/controller/http/v1/request"
)

type Context struct {
	Req          *request.GenerateRequest
	Instructions []Instruction
}

type Instruction struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func NewContext(req *request.GenerateRequest) *Context {
	return &Context{
		Req: req,
	}
}

func (pctx *Context) AddInstruction(instruction Instruction) {
	pctx.Instructions = append(pctx.Instructions, instruction)
}

type Pipeline struct {
	pipes []func(pctx *Context) error
}

func NewPipeline(pipes ...func(pctx *Context) error) *Pipeline {
	return &Pipeline{
		pipes: pipes,
	}
}

func (p *Pipeline) Execute(pctx *Context) (*Context, error) {
	for _, pipe := range p.pipes {
		if err := pipe(pctx); err != nil {
			return nil, fmt.Errorf("request: %w", err)
		}
	}
	return pctx, nil
}
