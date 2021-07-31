package external

import (
	"github.com/Nerzal/gocloak/v8"
)

type Keycloak struct {
	BasePath     string
	Realm        string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	Client       gocloak.GoCloak
}

func ConnectKeycloak(basePath, realm, username, password string) *Keycloak {
	k := &Keycloak{
		BasePath: basePath,
		Realm:    realm,
		Username: username,
		Password: password,
	}
	k.Client = gocloak.NewClient(k.BasePath)

	return k
}
