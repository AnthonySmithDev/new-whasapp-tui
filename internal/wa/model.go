package wa

import (
	"context"

	"github.com/knipferrc/bubbletea-starter/internal/repository"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var logMain waLog.Logger

type Client struct {
	waclient       *whatsmeow.Client
	eventHandlerID uint32

	QRChannel      <-chan whatsmeow.QRChannelItem
	MessageChannel chan events.Message
	HistorySync    chan events.Message

	db *repository.DB
}

func NewClient(db *repository.DB) *Client {
	logMain = Stdout("Main", "", true)
	msgChannel := make(chan events.Message)
	return &Client{
		db:             db,
		MessageChannel: msgChannel,
	}
}

func (cli *Client) GetQRChannel() {
	clientLog := Stdout("Client", "", true)
	cli.waclient = whatsmeow.NewClient(cli.GetDevice(), clientLog)
	ch, err := cli.waclient.GetQRChannel(context.Background())
	if err != nil {
		logMain.Errorf("Failed to get QR channel: %v", err)
	} else {
		cli.QRChannel = ch
	}
}

func (cli *Client) Connect() {
	cli.eventHandlerID = cli.waclient.AddEventHandler(cli.eventHandler)
	if err := cli.waclient.Connect(); err != nil {
		logMain.Errorf("Failed to connect: %v", err)
		return
	}
}

func (cli *Client) GetContact(jid string) types.ContactInfo {
	userJID, err := types.ParseJID(jid)
	if err != nil {
		return types.ContactInfo{FullName: jid}
	}
	contact, err := cli.waclient.Store.Contacts.GetContact(userJID)
	if err != nil {
		return types.ContactInfo{FullName: jid}
	}
	return contact
}
