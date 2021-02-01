package sdk

// WorkflowCoreActivityReturn Standard return from activity
type WorkflowCoreActivityReturn struct {
	Value      interface{}
	Error      error
	SignalName string
}
