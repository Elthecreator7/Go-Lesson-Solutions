package main

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	currentMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return currentMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	var bill int
	bill = costPerSend * messagesSent
	return bill
}
