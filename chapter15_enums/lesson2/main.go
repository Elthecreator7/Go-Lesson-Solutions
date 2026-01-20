package main

type emailStatus int

const (
	EmailBounced = iota
	EmailInvalid
	EmailDelivered
	EmailOpened
)
