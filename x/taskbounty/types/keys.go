package types

const (
	// ModuleName defines the module name
	ModuleName = "taskbounty"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName
)

var (
	// TaskKeyPrefix is the prefix for storing tasks
	TaskKeyPrefix = []byte{0x01}
	
	// TaskCountKey defines key to store the count of tasks
	TaskCountKey = []byte{0x02}
)

// TaskKey returns the store key to retrieve a Task from the index fields
func TaskKey(id string) []byte {
	return append(TaskKeyPrefix, []byte(id)...)
}
