package models

type ExecutableProvider string

type Archive struct {
	URL             string
	SHA256          string
	CanSkipDownload bool
}

type Distribution struct {
	Targets          []string   `yaml:"targets"`
	Version          string     `yaml:"version"`
	Executable       Executable `yaml:"executable"`
	SkipVerification bool       `yaml:"skipVerification"`
}

type Executable struct {
	// Provider specifies how the Address will be used. (Required)
	Provider ExecutableProvider `yaml:"provider"`

	// Provider specific data
	ProviderData map[string]string `yaml:"providerData"`

	// Address is the location of the executable archive. (Required)
	Address string `yaml:"address"`

	// SHA256 is the SHA256 checksum of the executable archive. (Optional)
	//
	// Required for the URI provider.
	SHA256 string `yaml:"sha256"`

	// Entrypoint specifies which file will be executed in the archive. Support templates. (Required)
	Entrypoint string `yaml:"entrypoint"`

	// Archive specifies which file will be downloaded. Support templates. (Required)
	Archive string `yaml:"archive"`
}
