package config

import (
	httproute "backend_hub/internal/adapter/http"
	"backend_hub/internal/adapter/http/controller"
	"backend_hub/internal/adapter/repository"
	"backend_hub/internal/adapter/validator"
	"backend_hub/internal/infrastructure/database"
	"backend_hub/internal/infrastructure/logger"
	"backend_hub/internal/usecase"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *database.Kind[*gorm.DB]
	MasterDB  *database.Kind[*gorm.DB]
	App       *echo.Echo
	Logger    *logger.Logger
	Validator *validator.CustomValidator
	Config    *viper.Viper
	// Publisher        *messaging.Publisher[any, any]
}

func (c *BootstrapConfig) SetDefaultConfigs() {
	setDefaultCoreConfigs(c.Config)
	setDefaultDatabaseConfig(c.Config)
	setMySQLDefaultConfig(c.Config)
}

func (c *BootstrapConfig) WatchConfig() {
	c.Config.OnConfigChange(func(fsnotify.Event) {
		c.SetDefaultConfigs()
	})
	c.Config.WatchConfig()
}

// Bootstrap bootstrap app
func Bootstrap(config *BootstrapConfig) error {

	// Init repositories
	// repo := repository.NewRepository(config.DB)
	// trxRepo := repository.NewTransactionRepositoryImplementation(config.DB.Write)
	tesRepositoory := repository.NewRepositoryTes(config.Logger, config.DB)
	warehouseRepository := repository.NewProductRepository(config.Logger, config.DB)
	// use this fpr repo and trxRepo

	// Init usecases
	healtcheckUseCase := usecase.NewHealthCheckUseCase(
		config.Logger,
	)
	tesUseCase := usecase.NewUseCaseTes(config.Logger, config.Validator, tesRepositoory)
	warehouseUsecase := usecase.NewWarehouseUseCase(
		config.Logger,
		config.Validator,
		warehouseRepository,
	)
	// Init Controller
	healtcheckController := controller.NewHealthCheckController(config.Logger, healtcheckUseCase)
	logController := controller.NewLogController(config.Logger, config.Validator)
	tesController := controller.NewTesController(config.Logger, config.Validator, tesUseCase)
	warehouseControler := controller.NewWarehouseController(config.Logger, warehouseUsecase)

	route := httproute.Route{
		App:                   config.App,
		HealthCheckController: healtcheckController,
		LogController:         logController,
		TesController:         tesController,
		WarehouseController:   warehouseControler,
	}

	route.Setup()

	return nil
}
