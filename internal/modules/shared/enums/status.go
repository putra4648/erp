package enums

type Status string

const (
	StatusDraft     Status = "DRAFT"
	StatusPending   Status = "PENDING"
	StatusApproved  Status = "APPROVED"
	StatusRejected  Status = "REJECTED"
	StatusCompleted Status = "COMPLETED"
	StatusCancelled Status = "CANCELLED"
	StatusVoid      Status = "VOID"
)
