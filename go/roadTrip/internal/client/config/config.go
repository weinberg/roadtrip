package config

import (
  "errors"
  "github.com/google/uuid"
  "github.com/kirsle/configdir"
  "gopkg.in/yaml.v2"
  "os"
)

type ClientConfig struct {
  Characters []CharacterInfo
  Server     string
  Port       string
}

type CharacterInfo struct {
  UUID string
}

// LoadConfig loads the client config from an OS specific location.
// If no config file is found it attempts to create one.
func LoadConfig() (ClientConfig, error) {
  conf := ClientConfig{}
  configFilePath, err := getConfigFilePath()
  if err != nil {
    return conf, err
  }

  data, err := os.ReadFile(configFilePath)
  if os.IsNotExist(err) {
    err = initConfig()
    if err != nil {
      return conf, err
    }
    data, err = os.ReadFile(configFilePath)
  }
  if err != nil {
    return conf, err
  }

  err = yaml.Unmarshal(data, &conf)
  if err != nil {
    return conf, err
  }

  return conf, nil
}

// SaveConfig stores the given ClientConfig in an OS specific location.
func SaveConfig(conf ClientConfig) error {
  configFilePath, err := getConfigFilePath()
  if err != nil {
    return err
  }

  d, err := yaml.Marshal(&conf)
  if err != nil {
    return err
  }
  writeErr := os.WriteFile(configFilePath, d, 0666)
  if writeErr != nil {
    return writeErr
  }

  return nil
}

// initConfig writes a new config file.
func initConfig() error {
  newConfig := ClientConfig{
    Server: "localhost",
    Port:   "9066",
    Characters: []CharacterInfo{
      CharacterInfo{
        UUID: uuid.NewString(),
      }},
  }

  return SaveConfig(newConfig)
}

// getConfigFilePath returns the path to the config file.
// It will create the directory to hold it if it does not exist.
func getConfigFilePath() (string, error) {
  var err error
  configPath := configdir.LocalConfig("road-trip")
  err = configdir.MakePath(configPath) // Ensure it exists.
  if err != nil {
    return "", errors.New("Cannot access folder: '" + configPath + "' to store config file.")
  }

  return configPath + string(os.PathSeparator) + "player.yaml", nil
}
