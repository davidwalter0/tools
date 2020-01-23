package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davidwalter0/tools/x/webauthn.io/config"
	yaml "gopkg.in/yaml.v2"
)

// App configuration structure
type App struct {
	WebAuthn *config.Config
}

// Load a file to a byte array
func Load(path string) (content []byte, err error) {
	if content, err = ioutil.ReadFile(path); err == nil {
		if err != nil {
			content = nil
		}
	}
	return
}

// valid tests for empty configuration values
func (app *App) valid() bool {
	return app.WebAuthn.DBName == "" ||
		app.WebAuthn.DBPath == "" ||
		app.WebAuthn.HostAddress == "" ||
		app.WebAuthn.HostPort == "" ||
		app.WebAuthn.Cert == "" ||
		app.WebAuthn.Ca == "" ||
		app.WebAuthn.Key == ""
}

////////// // DumpYAML info from struct
////////// func (app *App) DumpYAML() {
////////// 	cfg.DumpYAML(*app)
////////// }

////////// // DumpJSON info from struct
////////// func (app *App) DumpJSON() {
////////// 	cfg.DumpJSON(*app)
////////// }

// DumpYAML info from struct
func (app *App) DumpYAML() {
	var err error
	var j = []byte{}
	if j, err = yaml.Marshal(app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(j))
}

// DumpJSON info from struct
func (app *App) DumpJSON() {
	var err error
	var j = []byte{}
	if j, err = json.MarshalIndent(app, "", "  "); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(j))
}

///////// // ParseCfg from a file to initialize the certificate definition
///////// func (app *App) Parse() (err error) {
///////// 	var text = []byte{}
///////// 	// if len(app.YamlCfg) > 0 && len(app.JSONCfg) > 0 {
///////// 	// 	return fmt.Errorf("Use --yamlcfg or --jsoncfg not both")
///////// 	// }
///////// 	if len(app.WebAuthn.Config) > 0 {
///////// 		text, err = Load(app.WebAuthn.Config)
///////// 		if err != nil {
///////// 			fmt.Printf("config %+v\n", app.WebAuthn.Config)
///////// 			fmt.Printf("config %+v\n", *app)
///////// 			// return err
///////// 		}
///////// 		err = yaml.Unmarshal(text, &app.WebAuthn)
///////// 		if err != nil {
///////// 			fmt.Printf("yaml? config %+v\n", app.WebAuthn.Config)
///////// 			fmt.Printf("yaml? config %+v\n", *app)
///////// 			err = json.Unmarshal(text, &app.WebAuthn)
///////// 			if err != nil {
///////// 				fmt.Printf("json? config %+v\n", app.WebAuthn.Config)
///////// 				fmt.Printf("json? config %+v\n", *app)
///////// 				return err
///////// 			}
///////// 		}
///////// 		fmt.Printf("config %+v\n", app.WebAuthn.Config)
///////// 		fmt.Printf("config %+v\n", *app)
///////// 		fmt.Printf("%s\n", *app)
///////// 	}
///////// 	return nil
///////// }

// Parse from a file to initialize the certificate definition
func (app *App) Parse() (err error) {
	var text = []byte{}
	fmt.Println("app.WebAuthn.Config", app.WebAuthn.Config)
	if len(app.WebAuthn.Config) > 0 {
		text, err = Load(app.WebAuthn.Config)
		if err == nil {
			err = json.Unmarshal(text, &app.WebAuthn)
			if err == nil {
				return
			}
			err = yaml.Unmarshal(text, &app.WebAuthn)
		}
	}
	return nil
}

/*
// ParseCfg from a file to initialize the certificate definition
func (app *App) ParseCfg() error {
	if len(app.WebAuthn.YamlCfg) > 0 && len(app.WebAuthn.JSONCfg) > 0 {
		return fmt.Errorf("Use --yamlcfg or --jsoncfg not both")
	}
	if len(app.WebAuthn.YamlCfg) > 0 {
		text, err := Load(app.WebAuthn.YamlCfg)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(text, app); err != nil {
			return err
		}
	}

	if len(app.WebAuthn.JSONCfg) > 0 {
		text, err := Load(app.WebAuthn.JSONCfg)
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
