package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	IO              Category = "IO"
	Notification    Category = "Notification"
	Otp             Category = "Otp"
)

const (
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"

	Migration SubCategory = "Migration"
	Select    SubCategory = "Select"
	Insert    SubCategory = "Insert"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Rollback  SubCategory = "Rollback"

	RedisSet SubCategory = "RedisSet"
	RedisGet SubCategory = "RedisGet"

	SendNotification SubCategory = "SendNotification"

	OtpGenerate SubCategory = "OtpGenerate"

	Api             SubCategory = "Api"
	HashPassword    SubCategory = "HashPassword"
	ValidationError SubCategory = "ValidationError"
)

const (
	AppName             ExtraKey = "AppName"
	LoggerName          ExtraKey = "LoggerName"
	ClientIp            ExtraKey = "ClientIp"
	Method              ExtraKey = "Method"
	StatusCode          ExtraKey = "StatusCode"
	BodySize            ExtraKey = "BodySize"
	Path                ExtraKey = "Path"
	RequestBody         ExtraKey = "RequestBody"
	ResponseBody        ExtraKey = "ResponseBody"
	ErrorMessage        ExtraKey = "ErrorMessage"
	NotificationMessage ExtraKey = "NotificationMessage"
	NotificationSubject ExtraKey = "NotificationSubject"
	UserName            ExtraKey = "UserName"
	Email               ExtraKey = "Email"
)
