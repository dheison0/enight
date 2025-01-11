package bot

import (
	"fmt"
	"log"
	"server/database"
	"server/extra"
	"server/models"
	"server/tokens"
	"strconv"

	"go.mau.fi/whatsmeow/types/events"
)

type NewClient struct {
	Client models.ClientResponse
	Stage  string
}

var newClients = map[string]*NewClient{}
var chatStages = map[string]string{}

// EventHandler handles all incoming events from whatsapp and choose
// what to do with it
func EventHandler(event any) {
	switch v := event.(type) {
	case *events.Message:
		if v.Info.IsFromMe || v.Info.IsGroup {
			return
		}
		user := v.Info.Sender.User

		if chatStages[user] == "new" { // this is being used to avoid unneeded database query
			RegisterNewClient(v)
			return
		}

		// check if it's a new client or a old one
		client, err := database.GetClient(user)
		if err == nil && chatStages[user] == "" {
			chatStages[user] = "normal"
		} else if err != nil {
			chatStages[user] = "new"
			RegisterNewClient(v)
			return
		}

		CommandByStage(v, user, client)
	}
}

func CommandByStage(m *events.Message, user string, client *models.ClientResponse) {
	switch chatStages[user] {
	case "normal":
		sendText(
			m, true,
			extra.Dedent(fmt.Sprintf(`
        Olá %s!
        O que deseja fazer?
        
        1 - Novo pedido
        2 - Cancelar pedido
        3 - Conversar com um atendente`,
				client.Name,
			)),
		)
		chatStages[user] = "command"
	case "command":
		menuCommand(m, user, client)
	case "chat":
		return
	}
}

func menuCommand(m *events.Message, user string, client *models.ClientResponse) {
	cmd := m.Message.GetConversation()
	switch cmd {
	case "1":
		token := tokens.Create(user)
		sendText(m, true,
			fmt.Sprintf(
				"Acesse o link a seguir para escolher os produtos! http://0.0.0.0:8080/b/%s",
				token,
			),
		)
	case "2":
		sendText(m, true, "TODO!")
	case "3":
		sendText(m, false, "Ok, agora você vai falar diretamento com o atendente!")
		chatStages[user] = "chat"
	default:
		chatStages[user] = "normal"
		CommandByStage(m, user, client)
	}
}

func RegisterNewClient(m *events.Message) {
	user := m.Info.Sender.User
	if _, ok := newClients[user]; !ok {
		newClients[user] = &NewClient{
			Client: models.ClientResponse{
				Phone:    user,
				Name:     m.Info.PushName,
				Location: models.Location{},
			},
			Stage: "name",
		}
	}
	client := newClients[user]
	switch client.Stage {
	case "name":
		sendText(m, true, "Olá, você é novo(a) por aqui! Qual é o seu nome?")
		client.Stage = "location"
	case "location":
		client.Client.Name = m.Message.GetConversation()
		locations, err := database.GetAllLocations()
		if err != nil {
			sendText(m, true, "Houve um erro ao buscar as localizações. Tente novamente mais tarde.")
			log.Print("Error getting locations: " + err.Error())
			return
		}
		locationsString := ""
		for _, location := range locations {
			locationsString += fmt.Sprintf("%d - %s\n", location.ID, location.Name)
		}
		locationsString += "0 - Nenhuma das anteriores"
		sendText(m, true, "Qual é o número da sua localização na lista abaixo?\n"+locationsString)
		client.Stage = "location_id"
	case "location_id":
		locationID := m.Message.GetConversation()
		if locationID == "0" {
			sendText(m, false, "Sinto muito por isso 😔, agora você pode conversar com nosso atendente para resolver esse problema 😅")
			client.Stage = "chat"
		}
		locationIDInt, err := strconv.Atoi(locationID)
		if err != nil {
			sendText(m, true, "Opa, parece que o número da localização que vocé escolheu é inválido. Por favor, escolha novamente.")
			break
		}
		location, err := database.GetLocation(locationIDInt)
		if err != nil {
			sendText(m, true, "O número da localização que você escolheu não foi encontrado no sistema. Por favor, escolha novamente.")
			break
		}
		client.Client.Location = *location
		err = database.CreateClient(&models.ClientDatabase{
			Phone:      user,
			Name:       client.Client.Name,
			LocationID: client.Client.Location.ID,
		})
		delete(newClients, user)
		if err == nil {
			sendText(m, true, "Obrigado por se registrar! 😉")
			chatStages[user] = "normal"
			EventHandler(m)
		} else {
			sendText(m, false, "Houve um erro ao criar o seu cadastro. Esse chat passará a ser usado para conversar com um atendente!")
			chatStages[user] = "chat"
		}
	}
}
