package app

const (
	argId            = "id"
	argVersion       = "version"
	argDocument      = "document"
	StatusReady      = "READY"
	StatusProcessing = "PROCESSING"
	StatusError      = "ERROR"
	LatestVersion    = "latest"
	ReservedVersion  = "v0"
)

// Enum values for AttributeStatus
const (
	AttributeStatusReady      string = "READY"
	AttributeStatusProcessing string = "PROCESSING"
	AttributeStatusError      string = "ERROR"
)

// Enum values for AttributeVersion
const (
	AttributeVersionLatestVersion   string = "latest"
	AttributeVersionReservedVersion string = "v0"
)

// Enum values for AttributeTableName
const (
	AttributeTableNameApp string = "dynamodb-table-app"
)
