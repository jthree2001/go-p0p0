package example

import (
  "fmt"

  "github.com/go-chat-bot/bot"
)

func hello(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf(":holdup: Hello ccivilian %s :holdup:", command.User.RealName)
	return
}

func init() {
	bot.RegisterCommand(
		"hello",
		"Says hi",
    "",
		hello)
}
