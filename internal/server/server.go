package server

type Server interface {
	Start() error
	ServerType() string
}
