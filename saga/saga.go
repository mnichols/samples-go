package saga

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/workflow"
)

var acts *Activities

func SagaWorkflow(ctx workflow.Context, opts *Opts) (SagaState, error) {

	var state SagaState
	var ac *ActivityResult
	var bc *ActivityResult
	var cc *ActivityResult
	var dc *ActivityResult
	var ec *ActivityResult
	if err := workflow.ExecuteActivity(ctx, acts.A, &Params{
		State: &state,
		Opts:  opts,
	}).Get(ctx, nil); err != nil {
		state.A = NewCompensatable(ac, err)
		return state, fmt.Errorf("no compensation hence we fail %w", err)
	}

	if err := workflow.ExecuteActivity(ctx, acts.B, &Params{
		State: &state,
		Opts:  opts,
	}).Get(ctx, &bc); err != nil {
		state.B = NewCompensatable(bc, err)
	}
	if err := workflow.ExecuteActivity(ctx, acts.C, &Params{
		State: &state,
		Opts:  opts,
	}).Get(ctx, &cc); err != nil {
		state.C = NewCompensatable(cc, err)
	}
	if err := workflow.ExecuteActivity(ctx, acts.D, &Params{
		State: &state,
		Opts:  opts,
	}).Get(ctx, &dc); err != nil {
		state.D = NewCompensatable(dc, err)
	}
	if err := workflow.ExecuteActivity(ctx, acts.E, &Params{
		State: &state,
		Opts:  opts,
	}).Get(ctx, nil); err != nil {
		state.E = NewCompensatable(ec, err)
	}
	return state, nil
}

type Activities struct {
}

func (a *Activities) A(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}
func (a *Activities) B(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}
func (a *Activities) C(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}
func (a *Activities) D(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}
func (a *Activities) E(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}

func (a *Activities) B1(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}
func (a *Activities) B2(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}
func (a *Activities) C1(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}
func (a *Activities) D1(ctx context.Context, params *Params) (*ActivityResult, error) {
	return nil, nil
}
