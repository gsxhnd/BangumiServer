package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrTypeQuery        = &Errno{Code: 10003, Message: "Error type of params."}
	ErrDataQuery        = &Errno{Code: 10004, Message: "Error data of query."}
	ErrValidation       = &Errno{Code: 10201, Message: "Validation failed."}
	ErrDatabase         = &Errno{Code: 10202, Message: "Database error."}
	ErrRecordNotFound   = &Errno{Code: 10203, Message: "Database record not found."}

	// token errors
	ErrToken           = &Errno{Code: 10101, Message: "Error occurred while signing the JSON web token."}
	ErrTokenExpire     = &Errno{Code: 10102, Message: "The token was expire."}
	ErrTokenWillExpire = &Errno{Code: 10103, Message: "The token will be expire."}
	ErrAuthFail        = &Errno{Code: 10104, Message: "Auth connect: connection refused"}
	ErrCodeNoExist     = &Errno{Code: 10105, Message: "The verification code is not exist."}
	ErrCodeInvalid     = &Errno{Code: 10106, Message: "The verification code was invalid."}

	// user errors
	ErrEncrypt              = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound         = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid         = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect    = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrKeyInvalid           = &Errno{Code: 20105, Message: "The key was invalid."}
	ErrTokenRoleInvalid     = &Errno{Code: 20106, Message: "The token role was invalid."}
	ErrRestPassword         = &Errno{Code: 20107, Message: "The new password was same as old password."}
	ErrOldPasswordIncorrect = &Errno{Code: 20108, Message: "The old password was incorrect."}
	ErrUserExist            = &Errno{Code: 20109, Message: "The user was exist."}
	ErrUserWalletNotExist   = &Errno{Code: 20110, Message: "The app user's wallet is not exist"}
)
