package domain

type Channel struct {
	Login           string `json:"login,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	ProfileImageURL string `json:"profileImageURL,omitempty"`
}

func NewChannel(login, displayName, profileImageURL string) Channel {
	return Channel{
		login,
		displayName,
		profileImageURL,
	}
}