package log

import (
	"strings"

	"github.com/cihub/seelog"
)

const template = `
<seelog>
    <outputs formatid="common">
        <rollingfile type="date" filename="${path}.log" datepattern="2006-01-02" maxrolls="7" />
    </outputs>
    <formats>
        <format id="common" format="%Date %Time [%LEVEL] %Msg%n" />
    </formats>
</seelog>
`

func Setup(path string) error {
	config := strings.Replace(template, "${path}", path, 1)
	logger, err := seelog.LoggerFromConfigAsString(config)

	if err != nil {
		return err
	}

	return seelog.ReplaceLogger(logger)
}
