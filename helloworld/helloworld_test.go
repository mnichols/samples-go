package helloworld

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	//var GenerateCheckRunSummaryReport = func(ctx context.Context) error { return nil }
	env.RegisterActivityWithOptions(func() error { return fmt.Errorf("not implemented") }, activity.RegisterOptions{Name: "GenerateCheckRunSummaryReport"})
	// Mock activity implementation
	env.OnActivity(Activity, mock.Anything, "Temporal").Return("Hello Temporal!", nil)
	env.OnActivity("GenerateCheckRunSummaryReport", mock.Anything).Return(nil)
	env.ExecuteWorkflow(Workflow, "Temporal")

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result string
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, "Hello Temporal!", result)
}

func Test_Activity(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(Activity)

	val, err := env.ExecuteActivity(Activity, "World")
	require.NoError(t, err)

	var res string
	require.NoError(t, val.Get(&res))
	require.Equal(t, "Hello World!", res)
}
