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
	// iterating through the analytics channel to print out the details of events on the console
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// setting environments
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5114609547669-5141250110704-PrrNivC08xenH5E8KRVDvj3d")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A053FCLF5J7-5141243392256-771b167c047a5b5546fd0e6f9b12c6fd17a081c45803d6fe11e6a5d7d152abf2")

	// assigning bot with token client
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// go routine to print out event info
	go printCommandEvents(bot.CommandEvents())

	// bot action
	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "Age Calculator",
		Examples:    []string{"my year of birth is 2001"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// taking string 'year' from the command argument
			year := request.Param("year")
			// converting the string year into INT for calculation
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			// calculating the age
			age := 2023 - yob

			// printing the age
			r := fmt.Sprintf("age is %d", age)
			// writing the response for the slack reply
			response.Reply(r)
		},
	})

	// using Background context as the base context to listen to requests
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// listening to events arose by bot ingested by ctx
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
