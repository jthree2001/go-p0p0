package p0p0

import (
  "fmt"

  "github.com/go-chat-bot/bot"
)

func hello(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf(":holdup: Hello civilian %s :holdup:", command.User.RealName)
	return
}

func init() {
	bot.RegisterCommand(
		"hello",
		"Says hi",
    "",
		hello)
}
