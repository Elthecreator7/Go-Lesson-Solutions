package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int {
	var importanceScore int
	if d.isUrgent {
		importanceScore = 50
	} else {
		importanceScore = d.priorityLevel
	}

	return importanceScore
}

func (g groupMessage) importance() int {
	importanceScore := g.priorityLevel

	return importanceScore
}

func (s systemAlert) importance() int {
	importanceScore := 100
	return importanceScore
}

func processNotification(n notification) (string, int) {
	dm, isDm := n.(directMessage)
	gm, isGm := n.(groupMessage)
	sa, isSa := n.(systemAlert)
	if isDm {
		return dm.senderUsername, dm.importance()
	}
	if isGm {
		return gm.groupName, gm.importance()
	}
	if isSa {
		return sa.alertCode, sa.importance()
	}
	return "", 0
}
