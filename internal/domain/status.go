package domain

type TaskStatus string

const (
	StatusAssigned       TaskStatus = "ASSIGNED"
	StatusStarted        TaskStatus = "STARTED"
	StatusWaiting        TaskStatus = "WAITING"
	StatusFinishedSuccess TaskStatus = "FINISHED_SUCCESS"
	StatusFinishedError  TaskStatus = "FINISHED_ERROR"
)

var validTransitions = map[TaskStatus][]TaskStatus{
	StatusAssigned:       {StatusStarted},
	StatusStarted:        {StatusWaiting, StatusFinishedSuccess, StatusFinishedError},
	StatusWaiting:        {StatusWaiting, StatusFinishedSuccess, StatusFinishedError},
	StatusFinishedSuccess: {},
	StatusFinishedError:  {},
}

func (s TaskStatus) IsValid() bool {
	switch s {
	case StatusAssigned, StatusStarted, StatusWaiting, StatusFinishedSuccess, StatusFinishedError:
		return true
	}
	return false
}

func (s TaskStatus) CanTransitionTo(target TaskStatus) bool {
	allowed, ok := validTransitions[s]
	if !ok {
		return false
	}
	for _, t := range allowed {
		if t == target {
			return true
		}
	}
	return false
}

func (s TaskStatus) IsTerminal() bool {
	return s == StatusFinishedSuccess || s == StatusFinishedError
}

func (s TaskStatus) String() string {
	return string(s)
}
