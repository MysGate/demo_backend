package errno

var (
	/*
		code = 0 means correct return; code > 0 represents a return with an error
		example: 20001
		2 means error occurred in business logic
		00 reserved digits
		01 the specific error code
	*/
	OK = &Errno{Code: 0, Message: "success"}

	InternalServerErr = &Errno{Code: 10001, Message: "Internal server error."}
	BindErr           = &Errno{Code: 10002, Message: "Failed to bind request data to the struct."}

	InvalidRequestParameter = &Errno{Code: 20001, Message: "Invalid URL request parameter."}
	NotSupportedCase        = &Errno{Code: 20002, Message: "Not supported case with the given request parameter"}
)
