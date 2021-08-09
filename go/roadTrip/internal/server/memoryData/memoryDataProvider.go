package memoryData

import "github.com/brickshot/roadtrip/internal/server/types"

type MemoryProvider struct {
  types.DataProvider
}

type Config struct {
  types.InitConfig
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
