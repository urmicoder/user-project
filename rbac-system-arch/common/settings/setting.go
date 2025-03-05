package settings

import "time"

var (
	UserServerPort = ":5002"
)

const (
	ApiRateLimitKey         = "urmi_rate_limiting"
	ApiRateLimitMaxRequests = 10
	ApiRateLimitDuration    = time.Minute * 5
)

// Roles
const (
	ADMINROLE      = "ADMIN"
	SUPERADMINROLE = "SUPER"
	USERROLE       = "USER"
	GUESTROLE      = "GUEST"
)
