package auth

type AuthorizationType = string

const (
	AuthorizationTypeBearer AuthorizationType = "Bearer"
	AuthorizationTypeBasic  AuthorizationType = "Basic"
	AuthorizationTypeNone   AuthorizationType = "None"
	AuthorizationTypeOther  AuthorizationType = "Other"
)
