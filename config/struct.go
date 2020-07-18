package config

type layout struct {
	Username string `json:"username"`
	Password string `json:"password"`
	States   []struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"states"`
}
