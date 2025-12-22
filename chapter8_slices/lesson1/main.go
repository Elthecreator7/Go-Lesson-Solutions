package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgArr := [3]string{primary, secondary, tertiary}
	costOfFirstMsg := len(primary)
	costOfSec := len(primary) + len(secondary)
	costOfThird := len(primary) + len(secondary) + len(tertiary)
	msgCostArr := [3]int{costOfFirstMsg, costOfSec, costOfThird}

	return msgArr, msgCostArr
}
