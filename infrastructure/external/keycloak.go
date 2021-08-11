package external

import (
	"github.com/Nerzal/gocloak/v8"
)

type Keycloak struct {
	Realm    string
	Username string
	Password string
	Client   gocloak.GoCloak
}

func NewKeycloak(basePath, realm, username, password string) *Keycloak {
	k := &Keycloak{
		Realm:    realm,
		Username: username,
		Password: password,
	}
	k.Client = gocloak.NewClient(basePath)

	return k
}
