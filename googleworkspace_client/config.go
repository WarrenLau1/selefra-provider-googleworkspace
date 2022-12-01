package googleworkspace_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	CredentialFile        string `yaml:"credential_file"  mapstructure:"credential_file"`
	Credentials           string `yaml:"credentials"  mapstructure:"credentials"`
	ImpersonatedUserEmail string `yaml:"impersonated_user_email"  mapstructure:"impersonated_user_email"`
	TokenPath             string `yaml:"token_path"  mapstructure:"token_path"`
}
