package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"go.temporal.io/sdk/workflow"
)

type IoTJobsDataPlaneClient interface {
       DescribeJobExecution(ctx workflow.Context, input *iotjobsdataplane.DescribeJobExecutionInput) (*iotjobsdataplane.DescribeJobExecutionOutput, error)
       DescribeJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.DescribeJobExecutionInput) *IotjobsdataplaneDescribeJobExecutionResult

       GetPendingJobExecutions(ctx workflow.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error)
       GetPendingJobExecutionsAsync(ctx workflow.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) *IotjobsdataplaneGetPendingJobExecutionsResult

       StartNextPendingJobExecution(ctx workflow.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error)
       StartNextPendingJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) *IotjobsdataplaneStartNextPendingJobExecutionResult

       UpdateJobExecution(ctx workflow.Context, input *iotjobsdataplane.UpdateJobExecutionInput) (*iotjobsdataplane.UpdateJobExecutionOutput, error)
       UpdateJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.UpdateJobExecutionInput) *IotjobsdataplaneUpdateJobExecutionResult
}

type IoTJobsDataPlaneStub struct {
}

func NewIoTJobsDataPlaneStub() IoTJobsDataPlaneClient {
    return &IoTJobsDataPlaneStub{}
}

type IotjobsdataplaneDescribeJobExecutionResult struct {
	Result workflow.Future
}

func (r *IotjobsdataplaneDescribeJobExecutionResult) Get(ctx workflow.Context) (*iotjobsdataplane.DescribeJobExecutionOutput, error) {
    var output iotjobsdataplane.DescribeJobExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IotjobsdataplaneGetPendingJobExecutionsResult struct {
	Result workflow.Future
}

func (r *IotjobsdataplaneGetPendingJobExecutionsResult) Get(ctx workflow.Context) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error) {
    var output iotjobsdataplane.GetPendingJobExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IotjobsdataplaneStartNextPendingJobExecutionResult struct {
	Result workflow.Future
}

func (r *IotjobsdataplaneStartNextPendingJobExecutionResult) Get(ctx workflow.Context) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error) {
    var output iotjobsdataplane.StartNextPendingJobExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IotjobsdataplaneUpdateJobExecutionResult struct {
	Result workflow.Future
}

func (r *IotjobsdataplaneUpdateJobExecutionResult) Get(ctx workflow.Context) (*iotjobsdataplane.UpdateJobExecutionOutput, error) {
    var output iotjobsdataplane.UpdateJobExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *IoTJobsDataPlaneStub) DescribeJobExecution(ctx workflow.Context, input *iotjobsdataplane.DescribeJobExecutionInput) (*iotjobsdataplane.DescribeJobExecutionOutput, error) {
    var output iotjobsdataplane.DescribeJobExecutionOutput
    err := workflow.ExecuteActivity(ctx, "IoTJobsDataPlane.DescribeJobExecution", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTJobsDataPlaneStub) DescribeJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.DescribeJobExecutionInput) *IotjobsdataplaneDescribeJobExecutionResult {
    future := workflow.ExecuteActivity(ctx, "IoTJobsDataPlane.DescribeJobExecution", input)
    return &IotjobsdataplaneDescribeJobExecutionResult{Result: future}
}

func (a *IoTJobsDataPlaneStub) GetPendingJobExecutions(ctx workflow.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error) {
    var output iotjobsdataplane.GetPendingJobExecutionsOutput
    err := workflow.ExecuteActivity(ctx, "IoTJobsDataPlane.GetPendingJobExecutions", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTJobsDataPlaneStub) GetPendingJobExecutionsAsync(ctx workflow.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) *IotjobsdataplaneGetPendingJobExecutionsResult {
    future := workflow.ExecuteActivity(ctx, "IoTJobsDataPlane.GetPendingJobExecutions", input)
    return &IotjobsdataplaneGetPendingJobExecutionsResult{Result: future}
}

func (a *IoTJobsDataPlaneStub) StartNextPendingJobExecution(ctx workflow.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error) {
    var output iotjobsdataplane.StartNextPendingJobExecutionOutput
    err := workflow.ExecuteActivity(ctx, "IoTJobsDataPlane.StartNextPendingJobExecution", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTJobsDataPlaneStub) StartNextPendingJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) *IotjobsdataplaneStartNextPendingJobExecutionResult {
    future := workflow.ExecuteActivity(ctx, "IoTJobsDataPlane.StartNextPendingJobExecution", input)
    return &IotjobsdataplaneStartNextPendingJobExecutionResult{Result: future}
}

func (a *IoTJobsDataPlaneStub) UpdateJobExecution(ctx workflow.Context, input *iotjobsdataplane.UpdateJobExecutionInput) (*iotjobsdataplane.UpdateJobExecutionOutput, error) {
    var output iotjobsdataplane.UpdateJobExecutionOutput
    err := workflow.ExecuteActivity(ctx, "IoTJobsDataPlane.UpdateJobExecution", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTJobsDataPlaneStub) UpdateJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.UpdateJobExecutionInput) *IotjobsdataplaneUpdateJobExecutionResult {
    future := workflow.ExecuteActivity(ctx, "IoTJobsDataPlane.UpdateJobExecution", input)
    return &IotjobsdataplaneUpdateJobExecutionResult{Result: future}
}
