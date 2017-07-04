package main

type GoogleServicesJson struct {
	Clients []Client `json:"client"`
}

type Client struct {
	Info         ClientInfo `json:"client_info"`
	OauthClients []OauthClient `json:"oauth_client"`
	ApiKeys      []ApiKey `json:"api_key"`
}

type ClientInfo struct {
	AppId   string `json:"mobilesdk_app_id"`
	Android AndroidClientInfo `json:"android_client_info"`
}

type AndroidClientInfo struct {
	PackageName string `json:"package_name"`
}

type OauthClient struct {
	Id      string `json:"client_id"`
	Type    int `json:"client_type"`
	Android OauthAndroidInfo `json:"android_info"`
}

type OauthAndroidInfo struct {
	PackageName string `json:"package_name"`
	Hash        string `json:"certificate_hash"`
}

type ApiKey struct {
	CurrentKey string `json:"current_key"`
}

func (client Client) GetApiKey() string {
	return client.ApiKeys[0].CurrentKey
}

func (client Client) GetAppId() string {
	return client.Info.AppId
}

func (servics GoogleServicesJson) GetClientBy(packageName string) *Client {
	for _, client := range servics.Clients {
		if client.Info.Android.PackageName == packageName {
			return &client
		}
	}

	return nil
}
