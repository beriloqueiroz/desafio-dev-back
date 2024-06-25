package factories

import (
	"github.com/beriloqueiroz/desafio-dev-back/configs"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/infra/implements"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/usecase"
	"github.com/redis/go-redis/v9"
)

func NewGetMessageUseCase() *usecase.GetMsgsUseCase {
	configs, err := configs.LoadConfig([]string{"."})
	if err != nil {
		panic(err)
	}
	clientRedis := redis.NewClient(&redis.Options{
		Addr:     configs.CacheUri,
		Password: configs.CachePass,
		DB:       0,
	})
	redisCacheRepository := implements.NewRedisCacheRepository(clientRedis)
	cptecMessageGateway := implements.NewCptecMessageGateway()
	return usecase.NewGetMsgsUseCase(redisCacheRepository, cptecMessageGateway)
}
