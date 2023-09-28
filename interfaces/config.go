package interfaces

type AppConfig struct {
	AWSAccountID     string `yaml:"accountId"`
	AWSProfileName   string `yaml:"awsProfile"`
	AWSProfileRegion string `yaml:"region"`
}
