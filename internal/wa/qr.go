package wa

import (
	"context"
	"errors"
	"fmt"

	"github.com/knipferrc/bubbletea-starter/internal/util/qr"
	"go.mau.fi/whatsmeow"
)

func (cli Client) GetQR(done chan struct{}) {
	ch, err := cli.WAClient.GetQRChannel(context.Background())
	if err != nil {
		if !errors.Is(err, whatsmeow.ErrQRStoreContainsID) {
			logMain.Errorf("Failed to get QR channel: %v", err)
		}
	} else {
		go func() {
			for evt := range ch {
				if evt.Event == "code" {
					fmt.Print(qr.Generate(evt.Code))
				} else {
					logMain.Infof("QR channel result: %s", evt.Event)
				}
			}
			done <- struct{}{}
		}()
	}
}
