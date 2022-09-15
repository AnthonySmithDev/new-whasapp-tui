package wa

import (
	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

const (
	dbDialect = "sqlite3"
	dbAddress = "file:store.db?_foreign_keys=on"
)

func (cli Client) GetDevice() *store.Device {
	dbLog := Stdout("Database", "", true)
	container, err := sqlstore.New(dbDialect, dbAddress, dbLog)
	if err != nil {
		logMain.Errorf("Failed to connect to database: %v", err)
		return nil
	}
	device, err := container.GetFirstDevice()
	if err != nil {
		logMain.Errorf("Failed to get device: %v", err)
		return nil
	}
	return device
}
