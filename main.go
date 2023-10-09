package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

// main function sets the environment variables for the Slack bot token and app token, creates a new Slack bot client, and listens for commands.
func main() {
	// Set the Slack bot token and app token environment variables.
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6012312555300-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0606V9S221-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Create a new Slack bot client.
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Print command events.
	go printCommandEvents(bot.CommandEvents())

	// Define a new command for the bot to listen to.
	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// Get the year parameter from the command request.
			year := request.Param("year")
			// Convert the year parameter to an integer.
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			// Calculate the age based on the year parameter.
			age := 2023 - yob
			// Format the response message with the calculated age.
			r := fmt.Sprintf("Age is %d", age)
			// Send the response message back to the user.
			response.Reply(r)
		},
	})

	// Create a new context and cancel function.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for commands using the bot client.
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
