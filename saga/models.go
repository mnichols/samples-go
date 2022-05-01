package saga

type Compensatable struct {
	Result *ActivityResult
	Err    *string
}

func NewCompensatable(result *ActivityResult, err error) *Compensatable {
	s := err.Error()
	return &Compensatable{Result: result, Err: &s}
}

type SagaState struct {
	A *Compensatable
	B *Compensatable
	C *Compensatable
	D *Compensatable
	E *Compensatable
}
type Params struct {
	Opts  *Opts
	State *SagaState
}
type Opts struct {
	FailB bool
	FailC bool
	FailD bool
}
type ActivityResult struct{}
