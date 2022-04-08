package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Zap Logger created")

	serviceClient, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: client.DefaultNamespace,
	})

	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}

	worker := worker.New(serviceClient, "try_temporal", worker.Options{})

	worker.RegisterWorkflow(MyWorkflow)
	worker.RegisterActivity(MyActivity)

	err = worker.Start()
	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}

	select {}
}

func MyWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow myworkflow started")
	return nil
}

func MyActivity(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("My activity called")
	return nil
}
