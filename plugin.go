package bot

// A Plugin does an action on the received message.
type Plugin interface {
	DoAction(EventInfo) bool
	Help() string
}
