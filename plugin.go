package bot

// A Plugin does an action on the received message.
type Plugin interface {
	Hello(BotInfo)
	DoAction(EventInfo) bool
	Help() string
}
