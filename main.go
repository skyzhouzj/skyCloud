package skyCloud

import (
	"fmt"
	"github.com/skyzhouzj/skyCloud/configs"
	"github.com/skyzhouzj/skyCloud/pkg/env"
	"github.com/skyzhouzj/skyCloud/pkg/logger"
)

func main() {
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.Get().SkyCloud.ProjectName, env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.Get().SkyCloud.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(accessLogger)
}
