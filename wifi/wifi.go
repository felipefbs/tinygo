package wifi

import (
	"tinygo.org/x/drivers/netdev"
	nl "tinygo.org/x/drivers/netlink"
	link "tinygo.org/x/espradio/netlink"
)

func ConnectWiFi(ssid string, password string) (string, error) {
	// Cria o link usando a interface de rádio nativa do ESP
	espLink := link.Esplink{}
	netdev.UseNetdev(&espLink)

	// Tenta a conexão passando as credenciais
	err := espLink.NetConnect(&nl.ConnectParams{
		Ssid:       ssid,
		Passphrase: password,
	})
	if err != nil {
		return "", err
	}

	ip, err := espLink.Addr()
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}
