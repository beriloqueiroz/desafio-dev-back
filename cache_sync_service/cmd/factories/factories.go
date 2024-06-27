package factories

import (
	implements2 "github.com/beriloqueiroz/desafio-dev-back/cache_sync_service/internal/infra/implements"
	"github.com/beriloqueiroz/desafio-dev-back/cache_sync_service/internal/usecase"
	"github.com/beriloqueiroz/desafio-dev-back/configs"
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
	redisCacheRepository := implements2.NewRedisCacheRepository(clientRedis)
	cptecMessageGateway := implements2.NewCptecMessageGateway()
	return usecase.NewGetMsgsUseCase(redisCacheRepository, cptecMessageGateway)
}
