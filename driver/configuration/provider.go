package configuration

type Provider interface {
	LockfilePath() string
}
