// Package agent provides primitives for building structured LLM pipelines.
//
// Small models need more hand-holding than large ones. Rather than a single
// open-ended prompt, agents in this package are composed of explicit Steps
// that transform a shared State, making the reasoning process inspectable
// and testable.
package agent

import (
	"context"
	"fmt"
)

// State holds the mutable data that flows through a Pipeline.
// Steps read from and write to State, passing information forward.
type State struct {
	Input    string
	Output   string
	Memory   []MemoryEntry
	Metadata map[string]any
}

// MemoryEntry is a single item stored in the agent's working memory.
type MemoryEntry struct {
	Key     string
	Value   string
	Pinned  bool // pinned entries are always included in context regardless of budget
}

// Step is a single processing unit in a Pipeline.
type Step interface {
	Name() string
	Run(ctx context.Context, state *State) error
}

// Agent is a named entity that processes an input and produces an output.
type Agent interface {
	Name() string
	Run(ctx context.Context, input string) (string, error)
}

// Pipeline runs a slice of Steps in sequence, sharing a single State.
type Pipeline struct {
	name  string
	steps []Step
}

// NewPipeline creates a Pipeline with the given name and steps.
func NewPipeline(name string, steps ...Step) *Pipeline {
	return &Pipeline{name: name, steps: steps}
}

// Name returns the pipeline's name.
func (p *Pipeline) Name() string { return p.name }

// Run executes each step in order. If any step returns an error the pipeline
// halts and returns that error with the step name for context.
func (p *Pipeline) Run(ctx context.Context, input string) (string, error) {
	state := &State{
		Input:    input,
		Metadata: make(map[string]any),
	}
	for _, step := range p.steps {
		if err := step.Run(ctx, state); err != nil {
			return "", fmt.Errorf("step %q: %w", step.Name(), err)
		}
	}
	return state.Output, nil
}
