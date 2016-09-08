package config

import (
	"errors"
	"net/url"

	"github.com/stellar/go-stellar-base/keypair"
)

// Config contains config params of the compliance server
type Config struct {
	ExternalPort      *int   `mapstructure:"external_port"`
	InternalPort      *int   `mapstructure:"internal_port"`
	LogFormat         string `mapstructure:"log_format"`
	NeedsAuth         bool   `mapstructure:"needs_auth"`
	NetworkPassphrase string `mapstructure:"network_passphrase"`
	Database          struct {
		Type string
		URL  string
	}
	Keys
	Callbacks
	TLS struct {
		CertificateFile string `mapstructure:"certificate_file"`
		PrivateKeyFile  string `mapstructure:"private_key_file"`
	}
}

// Keys contains values of `keys` config group
type Keys struct {
	SigningSeed   string `mapstructure:"signing_seed"`
	EncryptionKey string `mapstructure:"encryption_key"`
}

// Callbacks contains values of `callbacks` config group
type Callbacks struct {
	Sanctions string
	AskUser   string `mapstructure:"ask_user"`
	FetchInfo string `mapstructure:"fetch_info"`
}

// Validate validates config and returns error if any of config values is incorrect
func (c *Config) Validate() (err error) {
	if c.ExternalPort == nil {
		err = errors.New("external_port param is required")
		return
	}

	if c.InternalPort == nil {
		err = errors.New("internal_port param is required")
		return
	}

	if c.NetworkPassphrase == "" {
		err = errors.New("network_passphrase param is required")
		return
	}

	if c.Keys.SigningSeed == "" || c.Keys.EncryptionKey == "" {
		err = errors.New("keys.signing_seed and keys.encryption_key params are required")
		return
	}

	if c.Keys.SigningSeed != "" {
		_, err = keypair.Parse(c.Keys.SigningSeed)
		if err != nil {
			err = errors.New("keys.signing_seed is invalid")
			return
		}
	}

	if c.Keys.EncryptionKey != "" {
		_, err = keypair.Parse(c.Keys.EncryptionKey)
		if err != nil {
			err = errors.New("keys.encryption_key is invalid")
			return
		}
	}

	var dbURL *url.URL
	dbURL, err = url.Parse(c.Database.URL)
	if err != nil {
		err = errors.New("Cannot parse database.url param")
		return
	}

	switch c.Database.Type {
	case "mysql":
		// Add `parseTime=true` param to mysql url
		query := dbURL.Query()
		query.Set("parseTime", "true")
		dbURL.RawQuery = query.Encode()
		c.Database.URL = dbURL.String()
	case "postgres":
		break
	default:
		err = errors.New("Invalid database.type param")
		return
	}

	if c.Callbacks.Sanctions != "" {
		_, err = url.Parse(c.Callbacks.Sanctions)
		if err != nil {
			err = errors.New("Cannot parse callbacks.sanctions param")
			return
		}
	}

	if c.Callbacks.Sanctions != "" {
		_, err = url.Parse(c.Callbacks.Sanctions)
		if err != nil {
			err = errors.New("Cannot parse callbacks.sanctions param")
			return
		}
	}

	if c.Callbacks.AskUser != "" {
		_, err = url.Parse(c.Callbacks.AskUser)
		if err != nil {
			err = errors.New("Cannot parse callbacks.ask_user param")
			return
		}
	}

	if c.Callbacks.FetchInfo != "" {
		_, err = url.Parse(c.Callbacks.FetchInfo)
		if err != nil {
			err = errors.New("Cannot parse callbacks.fetch_info param")
			return
		}
	}

	return
}
