package components

// BiographyCurrentState The user's current state, whether they are studying, working or unemployed.
type BiographyCurrentState string

// Current employment state of a WASAPhoto user.
const (
	Studying   BiographyCurrentState = "studying"
	Unemployed BiographyCurrentState = "unemployed"
	Working    BiographyCurrentState = "working"
)

const InternalServerError string = "{\"code\": 500, \"message\": \"Internal Server Error\"}"
const BadRequestError string = "{\"code\": 400, \"message\": \"Bad Request\"}"
const UnauthorizedError string = "{\"code\": 401, \"message\": \"Unauthorized\"}"

const InternalServerErrorF string = "{\"code\": 500, \"message\": \"Internal Server Error: %s\"}"
const BadRequestErrorF string = "{\"code\": 400, \"message\": \"Bad Request: %s\"}"
const UnauthorizedErrorF string = "{\"code\": 401, \"message\": \"Unauthorized: %s\"}"
