package memoryData

import (
  "github.com/brickshot/roadtrip/internal/server"
)

type MemoryProvider struct {
  playerServer.DataProvider
}

type Config struct {
  playerServer.InitConfig
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
