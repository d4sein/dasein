package executioner

const (
	// ErrCommandNotFound defines the error when a command is not found
	ErrCommandNotFound = "Command not found"
	// ErrNoPermission defines the error when a user does not have sufficient permission
	ErrNoPermission = "You can't run this command"
	// ErrTooManyValues defines the error when a user gives more values to an arg then required
	ErrTooManyValues = "Too many values for argument '%s'"
)
