package types

// PollOperation represents an operation for a poll command.
type PollOperation string

// Constants representing available poll operations.
const (
	PollOperationAcknowledge PollOperation = "ack"
	PollOperationRequest     PollOperation = "req"
)

// Poll represents a poll command.
type Poll struct {
	Poll PollCommand `xml:"command>poll"`
}

// PollCommand represents the (attribute) data from a poll command tag.
type PollCommand struct {
	Operation PollOperation `xml:"op,attr"`
	MessageID string        `xml:"msgID,attr"`
}
