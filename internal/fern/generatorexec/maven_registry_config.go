package generatorexec

type MavenRegistryConfig struct {
	RegistryUrl string `json:"registryUrl"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Group       string `json:"group"`
}