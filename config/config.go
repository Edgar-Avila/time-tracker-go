package config

import (
	"log"
	"os"
	"path"
)

var cfg *Config

type Config struct {
    DbPath string
}

func Get() *Config{
    if cfg == nil {
        configDir, err := os.UserConfigDir()
        if err != nil {
            log.Fatal(err)
        }
        dbdir := path.Join(configDir, "time-tracker-go")
        dbpath := path.Join(configDir, "time-tracker-go", "time-tracker.db")
        os.MkdirAll(dbdir, os.ModePerm)
        cfg = &Config{
            DbPath: dbpath,
        }
    }
    return cfg
}
