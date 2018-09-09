package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const token = "NDg3ODYxNzAwNTQ1MjE2NTEz.DnT0NA.-12IDIZL8zWzmzmV4Md5LsDr4g0"
const AuthToken = "d6d57dbfc3b74b0f8700d954257837c5"
const TgToken = "651393093:AAHRCCVVLWDRaj4zl7fv2jbfdVuxLNhFR-c"
var serversShut = make(map[string]bool)
var bot *telebot.Bot

func main(){
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	discord.Open()
	botT, err := telebot.NewBot(telebot.Settings{
		Token: TgToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil{
		log.Fatal(err)
	}
	bot = botT
	bot.Handle(telebot.OnText, onTgMessage)
	discord.AddHandler(messageCreate)
	discord.UpdateStatus(0, ">>shut for mute")
	bot.Start()
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
	bot.Stop()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){

	if m.Content == ">>shut" {
		serversShut[m.ChannelID] = !serversShut[m.ChannelID]
		if serversShut[m.ChannelID]{
			s.ChannelMessageSend(m.ChannelID, "Теперь я могу говорить!")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Я затихаю...")
		}
		return
	}

	if (m.Author.ID == s.State.User.ID) || (!serversShut[m.ChannelID]) {
		return
	}

	speech, err := getApiAiResponse(m.Content, m.Author.ID + m.ChannelID)
	if (err != nil) || (speech == "") {
		fmt.Println(err)
		s.ChannelMessageSend(m.ChannelID, "Мне нечего ответить на это...")
	}
	speechArray := strings.Split(speech, ":")
	if speechArray[0] == "command" {
		http.PostForm("http://127.0.0.1:8000/", url.Values{"address": {speechArray[1]}, "pizza": {speechArray[2]}, "rest": {speechArray[3]}})
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention() + " Заказ добавлен в базу!")
		return
	}
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention() + " " + speech)
	return

}


type ApiAiInput struct {
	Status struct {
		Code      int
		ErrorType string
	}
	Result struct {
		Action           *string
		ActionIncomplete bool
		Speech           string
	} `json:"result"`
}

func getApiAiResponse(m string, id string) (resp string, err error) {
	params := url.Values{}
	params.Add("query", m)
	params.Set("sessionId", id)

	link := fmt.Sprintf("https://api.api.ai/v1/query?V=20160518&lang=En&%s", params.Encode())
	ai, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return "", err
	}

	ai.Header.Set("Authorization", "Bearer " + AuthToken)

	if resp, err := http.DefaultClient.Do(ai); err != nil {
		return "", err
	} else {
		defer resp.Body.Close()

		var input ApiAiInput
		datastring, _ := ioutil.ReadAll(resp.Body)
		err := json.NewDecoder(strings.NewReader(string(datastring))).Decode(&input)
		if err != nil {
			return "", err
		}

		return input.Result.Speech, nil
	}
}


func onTgMessage(m *telebot.Message){
	speech, err := getApiAiResponse(m.Text, string(m.Sender.ID) + string(m.Chat.ID))
	if (err != nil) || (speech == "") {
		fmt.Println(err)
		bot.Send(m.Sender, "Мне нечего ответить на это...")
		return
	}
	speechArray := strings.Split(speech, ":")
	if speechArray[0] == "command" {
		http.PostForm("http://127.0.0.1:8000/", url.Values{"address": {speechArray[1]}, "pizza": {speechArray[2]}, "rest": {speechArray[3]}})
		bot.Send(m.Sender, "Заказ добавлен в базу!")
		return
	}
	bot.Send(m.Sender, speech)
	return
}
