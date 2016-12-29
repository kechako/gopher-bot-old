package bot

import "log"

// SetLogger sets the logger l to the bot.
func (b *Bot) SetLogger(l *log.Logger) {
	b.logger = l
}

func (b *Bot) logf(format string, v ...interface{}) {
	if b.logger != nil {
		b.logger.Printf(format, v...)
	}
}

func (b *Bot) log(v ...interface{}) {
	if b.logger != nil {
		b.logger.Print(v...)
	}
}
