package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Agents []struct {
	DistinguishedName string      `json:"distinguishedName"`
	Ou                interface{} `json:"ou"`
	Cn                string      `json:"cn"`
	UID               string      `json:"uid"`
	Sn                interface{} `json:"sn"`
	O                 string      `json:"o"`
	UserPassword      string      `json:"userPassword"`
	Parent            interface{} `json:"parent"`
	ParentName        interface{} `json:"parentName"`
	EntryUUID         string      `json:"entryUUID"`
	HasSubordinates   string      `json:"hasSubordinates"`
	Name              string      `json:"name"`
	IconPath          interface{} `json:"iconPath"`
	ExpandedUser      string      `json:"expandedUser"`
	Attributes        struct {
		Owner                 string `json:"owner"`
		EntryUUID             string `json:"entryUUID"`
		StructuralObjectClass string `json:"structuralObjectClass"`
		CreatorsName          string `json:"creatorsName"`
		UserPassword          string `json:"userPassword"`
		SubschemaSubentry     string `json:"subschemaSubentry"`
		Cn                    string `json:"cn"`
		HasSubordinates       string `json:"hasSubordinates"`
		O                     string `json:"o"`
		CreateTimestamp       string `json:"createTimestamp"`
		ModifyTimestamp       string `json:"modifyTimestamp"`
		UID                   string `json:"uid"`
		EntryCSN              string `json:"entryCSN"`
		ModifiersName         string `json:"modifiersName"`
		LiderDeviceOSType     string `json:"liderDeviceOSType"`
		PwdChangedTime        string `json:"pwdChangedTime"`
		EntryDN               string `json:"entryDN"`
	} `json:"attributes"`
	AttributesMultiValues struct {
		Owner                 []string `json:"owner"`
		EntryUUID             []string `json:"entryUUID"`
		StructuralObjectClass []string `json:"structuralObjectClass"`
		CreatorsName          []string `json:"creatorsName"`
		UserPassword          []string `json:"userPassword"`
		SubschemaSubentry     []string `json:"subschemaSubentry"`
		ObjectClass           []string `json:"objectClass"`
		Cn                    []string `json:"cn"`
		HasSubordinates       []string `json:"hasSubordinates"`
		O                     []string `json:"o"`
		CreateTimestamp       []string `json:"createTimestamp"`
		ModifyTimestamp       []string `json:"modifyTimestamp"`
		UID                   []string `json:"uid"`
		EntryCSN              []string `json:"entryCSN"`
		ModifiersName         []string `json:"modifiersName"`
		LiderDeviceOSType     []string `json:"liderDeviceOSType"`
		PwdChangedTime        []string `json:"pwdChangedTime"`
		EntryDN               []string `json:"entryDN"`
	} `json:"attributesMultiValues"`
	Type                string      `json:"type"`
	Priviliges          interface{} `json:"priviliges"`
	ChildEntries        interface{} `json:"childEntries"`
	TelephoneNumber     interface{} `json:"telephoneNumber"`
	HomePostalAddress   interface{} `json:"homePostalAddress"`
	CreateDateStr       string      `json:"createDateStr"`
	Mail                interface{} `json:"mail"`
	SessionList         interface{} `json:"sessionList"`
	AgentListSize       interface{} `json:"agentListSize"`
	OnlineAgentListSize interface{} `json:"onlineAgentListSize"`
	AgentList           interface{} `json:"agentList"`
	OnlineAgentList     interface{} `json:"onlineAgentList"`
	Online              bool        `json:"online"`
}


func (cu *ConsoleUser)NewAgents(agentLdapBaseDn string){

	print_info("Getting all the active computers")
	client := http.Client{Jar: cu.cookieJar}
	
	v := url.Values{}
	v.Set("uid", agentLdapBaseDn)
	v.Set("type", "ORGANIZATIONAL_UNIT")
	v.Set("name", "Agents")
	v.Set("parent", "")

    resp, err := client.PostForm(fmt.Sprintf("http://%s:8080/lider/computer/getOuDetails", TARGET), v)

	if err != nil {
		panic_with_msg("Unable to login somehow. Dunno why", err)
	}
	
	body, err := ioutil.ReadAll(resp.Body)

	agents := new(Agents)

	err2 := json.Unmarshal(body, agents)

	if err2 != nil {
		panic_with_msg("asdsad", err)
	}
}