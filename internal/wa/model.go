package wa

import (
	"go.mau.fi/whatsmeow"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var logMain waLog.Logger

type Client struct {
	WAClient       *whatsmeow.Client
	eventHandlerID uint32
}

func NewClient() Client {
	logMain = Stdout("Main", "", true)
	return Client{}
}

func (cli Client) Connect() {
	clientLog := Stdout("Client", "", true)
	cli.WAClient = whatsmeow.NewClient(cli.GetDevice(), clientLog)
	if cli.WAClient.Store.ID == nil {
		cli.GetQR()
	}
	cli.eventHandlerID = cli.WAClient.AddEventHandler(cli.eventHandler)
	if err := cli.WAClient.Connect(); err != nil {
		logMain.Errorf("Failed to connect: %v", err)
		return
	}
}