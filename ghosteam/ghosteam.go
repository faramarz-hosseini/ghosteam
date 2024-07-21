package ghosteam

import (
	"net/http"
	"time"

	"github.com/faramarz-hosseini/ghosteam.git/ghosteam/interfaces"
)

type Ghosteam struct {
	iSteamUser *interfaces.SteamUser
}

func NewClient(apiKey string) *Ghosteam {
	baseSteamInterface := interfaces.NewBase(
		&http.Client{Timeout: 5 * time.Second},
		"https://api.steampowered.com",
		apiKey,
	)
	return &Ghosteam{
		iSteamUser: &interfaces.SteamUser{Base: baseSteamInterface},
	}
}

func (g *Ghosteam) ISteamUser() *interfaces.SteamUser {
	return g.iSteamUser
}
