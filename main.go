package noah_explorer_api

import (
	"github.com/noah-blockchain/noah-explorer-api/api"
	"github.com/noah-blockchain/noah-explorer-api/core"
	"github.com/noah-blockchain/noah-explorer-api/database"
)

func main() {
	// init environment
	env := core.NewEnvironment()

	// connect to database
	db := database.Connect(env)
	defer database.Close(db)

	// create explorer
	explorer := core.NewExplorer(db, env)

	// create ws extender
	extender := core.NewExtenderWsClient(explorer)
	defer extender.Close()

	// subscribe to channel and add cache handler
	sub := extender.CreateSubscription(explorer.Environment.WsBlocksChannel)
	sub.OnPublish(explorer.Cache)
	extender.Subscribe(sub)

	// run api
	api.Run(db, explorer)
}
