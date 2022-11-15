package codes

import (
	"fmt"
)

type errCode int

const (
	ErrUnauthorized errCode = iota + 1000
	ErrTokenExpired
	ErrAssertClaims
	ErrMatchingToken
	ErrAssertUserid
	ErrUnauthorizedAgent
	ErrUnauthorizedUrl
	ErrKeyNotFound
	ErrEncodingFile
	ErrConvertingToPdf
	ErrDecodingData
	ErrIdNeeded
	ErrCreatingAccount
	ErrEmailExists
	ErrCreatingSalt
	ErrHashPassword
	ErrFetchingUser
	AccNotFound
	Success
	AccActivationInProcess
	IncorrectPassword
	ErrGenerateJwt
	ErrLogging
	ErrReadingReqBody
	ErrUnmarshall
	ErrParseRegDate
	ErrValidate
)

var errCodes = map[errCode]string{
	ErrUnauthorized:        "UnAuthorized",
	ErrTokenExpired:        "Token is expired",
	ErrMatchingToken:       "Compared literals are not same",
	ErrAssertClaims:        "unable to assert claims",
	ErrAssertUserid:        "unable to assert userid",
	ErrUnauthorizedAgent:   "UnAuthorized user agent",
	ErrUnauthorizedUrl:     "UnAuthorized url",
	ErrKeyNotFound:         "unable to find this Uuid",
	ErrEncodingFile:        "unable to json encode the data",
	ErrConvertingToPdf:     "unable to convert to pdf format",
	ErrIdNeeded:            "id needed",
	ErrDecodingData:        "unable to decode the data",
	ErrCreatingAccount:     "Problem creating account",
	ErrEmailExists:         "Email is already in use",
	ErrCreatingSalt:        "Unable to generate salt",
	ErrHashPassword:        "Unable to generate hashed password",
	Success:                "SUCCESS",
	AccActivationInProcess: "Account activation in progress",
	ErrFetchingUser:        "Problem fetching your account",
	AccNotFound:            "User account was not found",
	IncorrectPassword:      "Incorrect Password",
	ErrGenerateJwt:         "Unable to generate jwt token",
	ErrLogging:             "Problem logging into your account",
	ErrReadingReqBody:      "Unable to read request body",
	ErrUnmarshall:          "Unable to unmarshal request body",
	ErrParseRegDate:        "Unable to parse registration date",
	ErrValidate:            "Validation of fields failed",
}

func GetErr(code errCode) string {
	x, ok := errCodes[code]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%d: %s", code, x)
}
