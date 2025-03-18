package config

import (
	"backend_hub/pkg/constant"

	"github.com/spf13/viper"
)

// setDefaultCoreConfigs sets default configs for APP
func setDefaultCoreConfigs(viper *viper.Viper) {
	viper.SetDefault(constant.PORT, "4444")
	viper.SetDefault(constant.TIMEZONE, "Asia/Jakarta")
	viper.SetDefault(constant.TIMEOUT_GRACEFUL_SHUTDOWN, 10)
}
