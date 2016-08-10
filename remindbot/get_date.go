package remindbot

// GetDueDate finds the due date of the assignment from the datastore
func GetDueDate(assignment string) (string, error) {
	return "tomorrow", nil
}
