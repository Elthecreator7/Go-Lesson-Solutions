package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	statErr := em.recipient.updateStatus(em.status)

	if statErr != nil {
		return fmt.Errorf("error updating user status: %w", statErr)
	}

	trackErr := a.track(em.status)
	if trackErr != nil {
		return fmt.Errorf("error tracking user bounce: %w", trackErr)
	}

	return nil
}
