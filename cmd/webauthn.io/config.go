package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davidwalter0/tools/x/webauthn.io/config"
	yaml "gopkg.in/yaml.v2"
)

type App struct {
	WebAuthn *config.Config
	// DumpCfg  bool `json:"dumpcfg"     doc:"dump config and exit"`
	// JSONCfg  string `json:"jsoncfg"     doc:"file with json definition"`
	// YamlCfg  string `json:"yamlcfg"     doc:"file with yaml definition"`
	// Debug bool `json:"debug"       doc:"be more verbose"`
}

/*
type App struct {
	// DBType       string `json:"db-type"       yaml:"db-name"       doc:"db type: of [sqlite3]`
	// DBPath       string `json:"db-path"       yaml:"db-path"       doc:"path to the db"`
	// HostAddress  string `json:"address"       yaml:"address"       doc:"host address to listen on."`
	// HostPort     string `json:"port"          yaml:"port"          doc:"insecure port redirects to secure"`
	// SecurePort   string `json:"secure-port"   yaml:"secure-port"   doc:"tls secure port"`
	// LogFile      string `json:"log-file"      yaml:"log-file"      doc:"log-file is an optional path"`
	// RelyingParty string `json:"relying-party" yaml:"relying-party" doc"webauthn relying party"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`
	Ca      string `json:"ca"`
	DumpCfg bool   `json:"dumpcfg"     doc:"dump config and exit"`
	JSONCfg string `json:"jsoncfg"     doc:"file with json definition"`
	YamlCfg string `json:"yamlcfg"     doc:"file with yaml definition"`
	Config  string `json:"config"     doc:"json or yaml formatted config file (uses file extension for type)"`
	Debug   bool   `json:"debug"       doc:"be more verbose"`
}
*/

// Load a file to a byte array
func Load(path string) (content []byte, err error) {
	if content, err = ioutil.ReadFile(path); err == nil {
		if err != nil {
			content = nil
		}
	}
	return
}

// DumpYAML info from struct
func (app *App) DumpYAML() {
	var err error
	if app.WebAuthn.Debug || app.WebAuthn.DumpCfg {
		var j = []byte{}
		if j, err = yaml.Marshal(app); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(j))
	}
}

// DumpJSON info from struct
func (app *App) DumpJSON() {
	var err error
	if app.WebAuthn.Debug || app.WebAuthn.DumpCfg {
		var j = []byte{}
		if j, err = json.MarshalIndent(app, "", "  "); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(j))
	}
}

// ParseCfg from a file to initialize the certificate definition
func (app *App) Parse() (err error) {
	var text = []byte{}
	// if len(app.YamlCfg) > 0 && len(app.JSONCfg) > 0 {
	// 	return fmt.Errorf("Use --yamlcfg or --jsoncfg not both")
	// }
	if len(app.WebAuthn.Config) > 0 {
		text, err = Load(app.WebAuthn.Config)
		if err != nil {
			fmt.Printf("config %+v\n", app.WebAuthn.Config)
			fmt.Printf("config %+v\n", *app)
			// return err
		}
		err = yaml.Unmarshal(text, &app.WebAuthn)
		if err != nil {
			fmt.Printf("yaml? config %+v\n", app.WebAuthn.Config)
			fmt.Printf("yaml? config %+v\n", *app)
			err = json.Unmarshal(text, &app.WebAuthn)
			if err != nil {
				fmt.Printf("json? config %+v\n", app.WebAuthn.Config)
				fmt.Printf("json? config %+v\n", *app)
				return err
			}
		}
		fmt.Printf("config %+v\n", app.WebAuthn.Config)
		fmt.Printf("config %+v\n", *app)
		fmt.Printf("%s\n", *app)
	}
	return nil
}

/*
// ParseCfg from a file to initialize the certificate definition
func (app *App) ParseCfg() error {
	if len(app.YamlCfg) > 0 && len(app.JSONCfg) > 0 {
		return fmt.Errorf("Use --yamlcfg or --jsoncfg not both")
	}
	if len(app.YamlCfg) > 0 {
		text, err := Load(app.YamlCfg)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(text, app); err != nil {
			return err
		}
	}

	if len(app.JSONCfg) > 0 {
		text, err := Load(app.JSONCfg)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(text, app); err != nil {
			return err
		}
	}

	return nil
}
*/
