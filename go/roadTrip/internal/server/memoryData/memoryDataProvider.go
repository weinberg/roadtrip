package memoryData

import (
  "github.com/brickshot/roadtrip/internal/server"
)

type MemoryProvider struct {
  server.DataProvider
}

type Config struct {
  server.InitConfig
}

func (d MemoryProvider) Init(c Config) MemoryProvider {
  InitCars()
  InitCharacters()
  return d
}

func (d MemoryProvider) Shutdown() {
  ShutdownCars()
  ShutdownCharacters()
}
