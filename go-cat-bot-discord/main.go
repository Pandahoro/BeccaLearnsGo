package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

//variables used for CLI params

var (
	Token string
)

const KuteGoAPIURL = "https:///kutego-api-xxxxx-ew.a.run.app"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main(){
	//creat new discord session using bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating discord session,", err)
		return
	}

	//reg the message create func asa call back for its events
	dg.AddHandler(messageCreate)

	//receving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	//open websocket connection and listen
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	//wait until CNTRL-C or other term signal is received
	fmt.Println("Bot is running, Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	//cleanly close down session
	dg.Close()

}

type Cat struct{
	Name string `json: "name"`

}

//this will be called every time a new message is created on any channel its authed to have access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//ignore own messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!cat" {

		// call the kutego api and get cat
		response, err := http.Get(KuteGoAPIURL + "/cat/" + "AngryCat")
		if err != nil {
			fmt.Println(err)

		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "AngryCat.gif", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get AngryCat :<")
		}
	}

	if m.Content == "!random" {

		//call to get random cat
		response, err := http.Get(KuteGoAPIURL + "/cat/random/")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "random-cat.gif", response.Body)
			iff err != nil {
				fmt.Println(err)
			}
		
		} else {
			fmt.Println("Error: no get random cat :<")
		}
	}

	if m.Content == "!Cats" {
		// call and display list of cats
		response, err := http.Get(KuteGoAPIURL + "/cats/")
		if err != nil {
			fmt.Println(err)

		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// transform to [] byte
			body, err := ioutil.ReadALL(response.Body)
			if err != {
				fmt.Println(err)
			}

			//put only needed info of the JSON doc in arrray of cats
			var data []Cat
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println(err)
			}

			// create string with cats name and blank line
			var cats string.Builder
			for _, cat := range data {
				cats.WriteString(cat.Name + "\n")
			}

			//send a text message with list of cats
			_, err = s.ChannelMessageSend(m.ChannelID, cats.String())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: no found list of cats D:")
		}
	}

}