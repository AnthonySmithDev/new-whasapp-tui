package wa

import (
	// "encoding/json"
	// "fmt"
	// "mime"
	"fmt"
	"os"
	// "strings"
	// "sync/atomic"
	"time"

	"go.mau.fi/whatsmeow/appstate"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

var historySyncID int32
var startupTime = time.Now().Unix()

func Print(text string) {
	fmt.Println(text)
}

func (cli Client) eventHandler(rawEvt interface{}) {
	switch evt := rawEvt.(type) {
	case *events.AppStateSyncComplete:
		if len(cli.waclient.Store.PushName) > 0 && evt.Name == appstate.WAPatchCriticalBlock {
			err := cli.waclient.SendPresence(types.PresenceAvailable)
			if err != nil {
				logMain.Warnf("Failed to send available presence: %v", err)
			} else {
				logMain.Infof("Marked self as available")
			}
		}
	case *events.Connected, *events.PushNameSetting:
		cli.Connected <- struct{}{}
		if len(cli.waclient.Store.PushName) == 0 {
			return
		}
		// Send presence available when connecting and when the pushname is changed.
		// This makes sure that outgoing messages always have the right pushname.
		err := cli.waclient.SendPresence(types.PresenceAvailable)
		if err != nil {
			logMain.Warnf("Failed to send available presence: %v", err)
		} else {
			logMain.Infof("Marked self as available")
		}
	case *events.StreamReplaced:
		os.Exit(0)
	case *events.Message:
		cli.db.CreateMessage(evt)
		// metaParts := []string{fmt.Sprintf("pushname: %s", evt.Info.PushName), fmt.Sprintf("timestamp: %s", evt.Info.Timestamp)}
		// if evt.Info.Type != "" {
		// 	metaParts = append(metaParts, fmt.Sprintf("type: %s", evt.Info.Type))
		// }
		// if evt.Info.Category != "" {
		// 	metaParts = append(metaParts, fmt.Sprintf("category: %s", evt.Info.Category))
		// }
		// if evt.IsViewOnce {
		// 	metaParts = append(metaParts, "view once")
		// }
		// if evt.IsViewOnce {
		// 	metaParts = append(metaParts, "ephemeral")
		// }

		// logMain.Infof("Received message %s from %s (%s): %+v", evt.Info.ID, evt.Info.SourceString(), strings.Join(metaParts, ", "), evt.Message)

		// img := evt.Message.GetImageMessage()
		// if img != nil {
		// 	data, err := cli.waclient.Download(img)
		// 	if err != nil {
		// 		logMain.Errorf("Failed to download image: %v", err)
		// 		return
		// 	}
		// 	exts, _ := mime.ExtensionsByType(img.GetMimetype())
		// 	path := fmt.Sprintf("%s%s", evt.Info.ID, exts[0])
		// 	err = os.WriteFile(path, data, 0600)
		// 	if err != nil {
		// 		logMain.Errorf("Failed to save image: %v", err)
		// 		return
		// 	}
		// 	logMain.Infof("Saved image in message to %s", path)
		// }
	case *events.Receipt:
		if evt.Type == events.ReceiptTypeRead || evt.Type == events.ReceiptTypeReadSelf {
			logMain.Infof("%v was read by %s at %s", evt.MessageIDs, evt.SourceString(), evt.Timestamp)
		} else if evt.Type == events.ReceiptTypeDelivered {
			logMain.Infof("%s was delivered to %s at %s", evt.MessageIDs[0], evt.SourceString(), evt.Timestamp)
		}
	case *events.Presence:
		if evt.Unavailable {
			if evt.LastSeen.IsZero() {
				logMain.Infof("%s is now offline", evt.From)
			} else {
				logMain.Infof("%s is now offline (last seen: %s)", evt.From, evt.LastSeen)
			}
		} else {
			logMain.Infof("%s is now online", evt.From)
		}
	case *events.HistorySync:
		cli.db.CreateHistory(evt.Data, cli.waclient)

		// id := atomic.AddInt32(&historySyncID, 1)
		// fileName := fmt.Sprintf("history-%d-%d.json", startupTime, id)
		// file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
		// if err != nil {
		// 	logMain.Errorf("Failed to open file to write history sync: %v", err)
		// 	return
		// }
		// enc := json.NewEncoder(file)
		// enc.SetIndent("", "  ")
		// err = enc.Encode(evt.Data)
		// if err != nil {
		// 	logMain.Errorf("Failed to write history sync: %v", err)
		// 	return
		// }
		// logMain.Infof("Wrote history sync to %s", fileName)
		// _ = file.Close()
	case *events.AppState:
		logMain.Debugf("App state event: %+v / %+v", evt.Index, evt.SyncActionValue)
	case *events.KeepAliveTimeout:
		logMain.Debugf("Keepalive timeout event: %+v", evt)
		if evt.ErrorCount > 3 {
			logMain.Debugf("Got >3 keepalive timeouts, forcing reconnect")
			go func() {
				cli.waclient.Disconnect()
				err := cli.waclient.Connect()
				if err != nil {
					logMain.Errorf("Error force-reconnecting after keepalive timeouts: %v", err)
				}
			}()
		}
	case *events.KeepAliveRestored:
		logMain.Debugf("Keepalive restored")
	}
}