package headsman

import (
	"github.com/cloudfoundry/mariadb_ctrl/os_helper"
)

type Headsman interface {
	Chop()
}

type MysqlHeadsman struct {
	oh 				os_helper.OsHelper
	mysqlUsername 	string
	mysqlPassword	string
	executablePath	string
	haproxyIp		string
}

func NewMysqlHeadsman(oh os_helper.OsHelper, mysqlUsername string,
						mysqlPassword string, executablePath string, haproxyIp string) *MysqlHeadsman {
	return &MysqlHeadsman{
		oh : oh,
		mysqlUsername : mysqlUsername,
		mysqlPassword : mysqlPassword,
		executablePath: executablePath,
		haproxyIp : haproxyIp,
	}
}

func (mh *MysqlHeadsman) Chop() {
	mh.oh.RunCommand(mh.executablePath, mh.mysqlUsername, mh.mysqlPassword, mh.haproxyIp)
}
