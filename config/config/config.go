package config

import (
	"fmt"
	"os"
	"syscall"

	"msp-git.connext.com.cn/connext-go-common/common-util/pencrypt"
	"msp-git.connext.com.cn/connext-go-core/core-config/pconfig"
	"msp-git.connext.com.cn/connext-go-core/core-util/prpc"
	"msp-git.connext.com.cn/connext-go-third/third-db/pdb"
	"msp-git.connext.com.cn/connext-go-third/third-log/plog"
)

// ConfClient config client
var ConfClient *pconfig.ConfigClient

// RetryAction Retry Action
type RetryAction func(t int, serviceName ...string)

// Retry Retry function
func Retry(t int, serviceName string, opts ...RetryAction) {
	for _, o := range opts {
		o(t, serviceName)
	}
}

// SetConfig init config client
func SetConfig(t int, serviceName ...string) {
	if len(serviceName) != 1 {
		plog.Assert(fmt.Errorf("serviceName should be only one"))
	}
	m := make(map[string][]string, 0)
	m[pconfig.SERVICENAMESTR] = []string{serviceName[0]}
	m[pconfig.VERSIONTYPE] = []string{""}

	defer func() {
		if e := recover(); e != nil {
			plog.Error("config init", "fail %d time: %s", t, e.(error).Error())
			if t < RETRYTIMES {
				Retry(t+1, serviceName[0], SetConfig)
			} else {
				os.Exit(int(syscall.SIGQUIT))
			}
		}
	}()
	// rsp := pconfig.QueryCfg(m)
	rsp, err := prpc.NewHTTPClient().NewRequest(pconfig.FRAMECONFIG, pconfig.VAULTINFOURL).Params(m).GET()
	plog.Assert(err)
	data := pconfig.GetCfgInfo(rsp.Body)
	ConfClient = pconfig.SetCfgInfo(ConfClient, data)
	pconfig.CheckCfg(ConfClient)
}

// SetDB init db config
func SetDB(t int, serviceName ...string) {
	defer func() {
		if e := recover(); e != nil {
			plog.Error("DB init", "fail %d time: %s", t, e.(error).Error())
			if t < RETRYTIMES {
				Retry(t+1, serviceName[0], SetDB)
			} else {
				os.Exit(int(syscall.SIGQUIT))
			}
		}
	}()
	host, err := ConfClient.Get(pconfig.MYSQLHOSTSTR)
	plog.Assert(err)
	port, err := ConfClient.Get(pconfig.MYSQLPORTSTR)
	plog.Assert(err)
	user, err := ConfClient.Get(pconfig.MYSQLUSERSTR)
	plog.Assert(err)
	pass, err := ConfClient.Get(pconfig.MYSQLPWDSTR)
	plog.Assert(err)
	database, err := ConfClient.Get(pconfig.MYSQLDBSTR)
	plog.Assert(err)
	pwd := pencrypt.AesDecrypt(pass.(string), pconfig.ENCRYPTIONVALUESTR)

	err = pdb.SetDB(&pdb.DBConfig{
		Host:     host.(string),
		Port:     port.(string),
		User:     user.(string),
		Pass:     pwd,
		Database: database.(string),
	})
	plog.Assert(err)
	_, err = pdb.GetDBInstance()
	plog.Assert(err)
}
