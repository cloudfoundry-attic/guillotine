package main

import (
	"flag"
	"time"
	"os"
	"strconv"
	"io/ioutil"

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

var pidfile = flag.String(
	"pidfile",
	"",
	"Specifies the location of the file to write the PID to",
)

func main() {
	flag.Parse()

	err := ioutil.WriteFile(*pidfile, []byte(strconv.Itoa(os.Getpid())), 0644)
	if err != nil {
		panic(err)
	}

	for {
		iq := inquisitor.NewHttpInquisitor(*haproxyUser, *haproxyPassword, *haproxyIp)
		h := headsman.NewMysqlHeadsman(os_helper.NewImpl(), *executablePath, *mysqlUser, *mysqlPassword, *haproxyIp)
		me := magistrate.NewMagistrate(iq, h)

		me.DeliberateAndExecute()
		time.Sleep(.5)
	}
}
