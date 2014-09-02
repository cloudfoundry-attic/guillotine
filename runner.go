package main

import (
	"flag"

	"github.com/cloudfoundry-incubator/guillotine/inquisitor"
	"github.com/cloudfoundry-incubator/guillotine/magistrate"
	"github.com/cloudfoundry-incubator/guillotine/headsman"
	"github.com/cloudfoundry/mariadb_ctrl/os_helper"

)

var haproxyUser = flag.String(
	"mysqlUser",
	"root",
	"Specifies the user name for MySQL",
)

var haproxyPassword = flag.String(
	"mysqlPassword",
	"",
	"Specifies the password for connecting to MySQL",
)

var haproxyIp = flag.String(
	"haproxyIp",
	"",
	"Specifies location of the HAProxy node",
)

var executablePath = flag.String(
	"executablePath",
	"",
	"Specifies the location of the 'execute' executable",
)

var mysqlUser = flag.String(
	"mysqlUser",
	"root",
	"Specifies the user name for MySQL",
)

var mysqlPassword = flag.String(
	"mysqlPassword",
	"",
	"Specifies the password for connecting to MySQL",
)

func main() {
	flag.Parse()

	iq := inquisitor.NewHttpInquisitor(*haproxyUser, *haproxyPassword, *haproxyIp)
	h := headsman.NewMysqlHeadsman(os_helper.NewImpl(), *executablePath, *mysqlUser, *mysqlPassword, *haproxyIp)
	me := magistrate.NewMagistrate(iq, h)

	me.DeliberateAndExecute()
}
