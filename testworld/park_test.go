// +build testworld

package testworld

import (
	"testing"

	"fmt"
	"os"
)

// TODO remember to cleanup config files generated

var robert *manager

func TestMain(m *testing.M) {
	// TODO start POA geth here
	//runSmartContractMigrations()
	contractAddresses := getSmartContractAddresses()
	robert = newManager("ws://127.0.0.1:9546",
		"keystore", "", "testing", true, contractAddresses)
	err := robert.init()
	if err != nil {
		panic(err)
	}
	fmt.Printf("contract addresses %+v\n", contractAddresses)
	result := m.Run()
	robert.stop()
	os.Exit(result)
}

func TestHost_Happy(t *testing.T) {
	alice := robert.getHost("Alice")
	bob := robert.getHost("Bob")
	charlie := robert.getHost("Charlie")
	eAlice := alice.createHttpExpectation(t)
	eBob := bob.createHttpExpectation(t)
	eCharlie := charlie.createHttpExpectation(t)

	b, err := bob.id()
	if err != nil {
		t.Error(err)
	}

	c, err := charlie.id()
	if err != nil {
		t.Error(err)
	}
	res, err := alice.createInvoice(eAlice, map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "40",
			"currency":       "GBP",
			"net_amount":     "40",
		},
		"collaborators": []string{b.String(), c.String()},
	})
	if err != nil {
		t.Error(err)
	}
	docIdentifier := res.Value("header").Path("$.document_id").String().NotEmpty().Raw()
	if docIdentifier == "" {
		t.Error("docIdentifier empty")
	}
	getInvoiceAndCheck(eBob, docIdentifier, "GBP")
	getInvoiceAndCheck(eCharlie, docIdentifier, "GBP")
	fmt.Println("Host test success")
}
