package api

// BiographyCurrentState The user's current state, whether they are studying, working or unemployed.
type BiographyCurrentState string

// Current employment state of a WASAPhoto user.
const (
	Studying   BiographyCurrentState = "studying"
	Unemployed BiographyCurrentState = "unemployed"
	Working    BiographyCurrentState = "working"
)
