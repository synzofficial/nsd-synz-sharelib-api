package enum

type SuccessType string

const (
	SuccessType_CREATED SuccessType = "Create successfully"
	SuccessType_UPDATED SuccessType = "Update successfully"
	SuccessType_DELETED SuccessType = "Delete successfully"
)
