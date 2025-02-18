# MMORPG Shared Module

Reusable components for MMORPG servers.

## Features

- gRPC client/server utilities
- Configuration (Viper)
- Logging (Zap)
- PostgreSQL/Redis clients
- Protobuf definitions

## Usage

```go
import (
	"github.com/silkroadbeyond/mmorpg-shared/config"
	"github.com/silkroadbeyond/mmorpg-shared/grpc"
)

func main() {
	config.Load()
	grpcServer := grpc.NewServer()
}
```
