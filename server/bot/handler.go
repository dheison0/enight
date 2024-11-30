package bot

import (
	"fmt"
	"log"
	"server/database"
	"server/models"
	"strconv"

	"go.mau.fi/whatsmeow/types/events"
)

type NewClient struct {
	Client models.ClientResponse
	Stage  string
}

var newClients = map[string]*NewClient{}
var chatStages = map[string]string{}

func EventHandler(event any) {
	switch v := event.(type) {
	case *events.Message:
		if v.Info.IsFromMe || v.Info.IsGroup {
			return
		}
		log.Printf("New message from %s: %+v\n", v.Info.Sender.User, v.Message.GetConversation())
		client, err := database.GetClient(v.Info.Sender.User)
		if err != nil {
			RegisterNewClient(v)
			return
		}
		stage, ok := chatStages[v.Info.Sender.User]
		if !ok {
			chatStages[v.Info.Sender.User] = "normal"
			stage = "normal"
		}
		switch stage {
		case "normal":
			sendText(v, "Olá "+client.Name+"!"+"\n\nO que deseja fazer?\n1 - Novo pedido\n2 - Cancelar pedido\n2 - Conversar com um atendente")
			markRead(v)
		}
	}
}

func RegisterNewClient(m *events.Message) {
	phone := m.Info.Sender.User
	if _, ok := newClients[phone]; !ok {
		newClients[phone] = &NewClient{
			Client: models.ClientResponse{
				Phone:    phone,
				Name:     m.Info.PushName,
				Location: models.Location{},
			},
			Stage: "name",
		}
	}
	client := newClients[phone]
	switch client.Stage {
	case "name":
		sendText(m, "Olá, você é novo(a) por aqui! Qual é o seu nome?")
		client.Stage = "location"
	case "location":
		client.Client.Name = m.Message.GetConversation()
		locations, err := database.GetAllLocations()
		if err != nil {
			sendText(m, "Houve um erro ao buscar as localizações. Tente novamente mais tarde.")
			log.Print("Error getting locations: " + err.Error())
			return
		}
		locationsString := ""
		for _, location := range locations {
			locationsString += fmt.Sprintf("%d - %s\n", location.ID, location.Name)
		}
		locationsString += "0 - Nenhuma das anteriores"
		sendText(m, "Qual é o número da sua localização na lista abaixo?\n"+locationsString)
		client.Stage = "location_id"
	case "location_id":
		locationID := m.Message.GetConversation()
		if locationID == "0" {
			sendText(m, "Sinto muito por isso 😔, agora você pode conversar com nosso atendente para resolver esse problema 😅")
			client.Stage = "chat"
			return
		}
		locationIDInt, err := strconv.Atoi(locationID)
		if err != nil {
			sendText(m, "Opa, parece que o número da localização que vocé escolheu é inválido. Por favor, escolha novamente.")
			break
		}
		location, err := database.GetLocation(locationIDInt)
		if err != nil {
			sendText(m, "O número da localização que você escolheu não foi encontrado no sistema. Por favor, escolha novamente.")
			break
		}
		client.Client.Location = *location
		err = database.CreateClient(&models.ClientDatabase{
			Phone:      phone,
			Name:       client.Client.Name,
			LocationID: client.Client.Location.ID,
		})
		if err == nil {
			sendText(m, "Obrigado por se registrar! 😉")
			delete(newClients, phone)
			EventHandler(m)
		} else {
			sendText(m, "Houve um erro ao criar o seu cadastro. Esse chat passará a ser usado para conversar com um atendente!")
			client.Stage = "chat"
			break
		}

	case "chat":
		return
	}
	markRead(m)
}
