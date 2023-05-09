package console

import (
	"Server/application/core/gamemap"
	"Server/application/core/room"
	hasproto "Server/application/proto"
	"errors"
	"github.com/desertbit/grumble"
)

var App = grumble.New(&grumble.Config{
	Name:        "hasserver",
	Description: "A HideAndSike Server Application",
	Flags: func(f *grumble.Flags) {
		f.String("d", "directory", "DEFAULT", "set an alternative root directory path")
		f.Bool("v", "verbose", false, "enable verbose mode")
	},
})

func init() {
	App.AddCommand(&grumble.Command{
		Name: "display",
		Help: "display runtime message",
		Args: func(a *grumble.Args) {
			a.String("args", "display you want to see")
			a.Int("player", "select player to display info")
		},
		Run: func(c *grumble.Context) error {
			player := room.RoomMgrObj.GetPlayerByPid(int32(c.Args.Int("player")))
			if player != nil {
				msg := player.PlayerMessage
				if msg != nil {
					switch c.Args.String("args") {
					case "location":
						c.App.Println("location:  x:", msg.Info.Trans.L.X, "y:", msg.Info.Trans.L.Y, "z:", msg.Info.Trans.L.Z)
						break
					case "rotation":
						c.App.Println("rotation:  Pitch:", msg.Info.Trans.R.Pitch, "Roll", msg.Info.Trans.R.Roll, "Yaw", msg.Info.Trans.R.Yaw)
						break
					case "scale":
						c.App.Println("scale:  x:", msg.Info.Trans.S.X, "y:", msg.Info.Trans.S.Y, "z:", msg.Info.Trans.S.Z)
						break
					case "transform":
						c.App.Println("location:  x:", msg.Info.Trans.L.X, "y:", msg.Info.Trans.L.Y, "z:", msg.Info.Trans.L.Z)
						c.App.Println("rotation:  Pitch:", msg.Info.Trans.R.Pitch, "Roll", msg.Info.Trans.R.Roll, "Yaw", msg.Info.Trans.R.Yaw)
						c.App.Println("scale:     x:", msg.Info.Trans.S.X, "y:", msg.Info.Trans.S.Y, "z:", msg.Info.Trans.S.Z)
						break

					case "all":
						c.App.Println("UserID:   ", msg.Token.Pid)
						c.App.Println("Username: ", msg.Token.Name)
						c.App.Println("Token:    ", msg.Token.Token.Token)
						c.App.Println("location:  x:", msg.Info.Trans.L.X, "y:", msg.Info.Trans.L.Y, "z:", msg.Info.Trans.L.Z)
						c.App.Println("rotation:  Pitch:", msg.Info.Trans.R.Pitch, "Roll", msg.Info.Trans.R.Roll, "Yaw", msg.Info.Trans.R.Yaw)
						c.App.Println("scale:     x:", msg.Info.Trans.S.X, "y:", msg.Info.Trans.S.Y, "z:", msg.Info.Trans.S.Z)
						c.App.Println("crotation: Pitch:", msg.Info.CRotation.Pitch, "Roll", msg.Info.CRotation.Roll, "Yaw", msg.Info.CRotation.Yaw)
						c.App.Println("Action:   ", msg.ActionData)
					}
				}
			} else {
				errors.New("error：player not found")
			}
			return nil
		},
	})
	App.AddCommand(&grumble.Command{
		Name: "players",
		Help: "display all players id of game",
		Run: func(c *grumble.Context) error {
			players := room.RoomMgrObj.GetAllPlayers()
			for _, v := range players {
				c.App.Println("Room Player")
				c.App.Println("UserID: ", v.Pid, "    Name:  ", v.PlayerMessage.Token.Name)
			}
			Waiting := room.RoomMgrObj.GetAllWaitingPlayer()
			for _, v := range Waiting {
				c.App.Println("Waiting Player")
				c.App.Println("UserID: ", v.Pid, "    Name:  ", v.Name)
			}
			return nil
		},
	})
	App.AddCommand(&grumble.Command{
		Name: "say",
		Help: "Say message to player",
		Args: func(a *grumble.Args) {
			a.String("message", "display you want to see")
			a.Int("player", "select player to display info")
		},
		Run: func(c *grumble.Context) error {
			player := room.RoomMgrObj.GetPlayerByPid(int32(c.Args.Int("player")))
			if player != nil {
				has := hasproto.PlayerChat{
					Type: 0,
					Chat: c.Args.String("message"),
				}
				has.Pid = append(has.Pid, 0)
				player.SendChat(&has)
			}
			return nil
		},
	})
	App.AddCommand(&grumble.Command{
		Name: "map",
		Help: "display map message",
		Args: func(a *grumble.Args) {
			a.String("args", "display actor message")
		},
		Run: func(c *grumble.Context) error {
			Actorname := c.Args.String("args")
			if Actorname == "test" {
				Origin := gamemap.Maputils.NewVector(-825.536, -402.800, 305.900)
				Direction := gamemap.Maputils.NewVector(0.423, 0.906, 0.000)
				Ray := gamemap.Maputils.NewRay(Origin, Direction)
				isCollsion := gamemap.Maputils.ActorRayIntersect("Acote_father", Ray)
				if isCollsion {
					c.App.Println("射线射中了Actor!")
				} else {
					c.App.Println("射线没有射中Actor!")
				}
			} else {
				c.App.Println(gamemap.Maputils.Entities[Actorname])
			}
			return nil
		},
	})
}
