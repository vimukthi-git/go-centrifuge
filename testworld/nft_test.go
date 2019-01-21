// +build testworld

package testworld

import (
	"net/http"
	"testing"

	"github.com/centrifuge/go-centrifuge/config"

	"github.com/stretchr/testify/assert"
)

func TestPaymentObligationMint_invoice_successful(t *testing.T) {
	t.Parallel()
	paymentObligationMint(t, typeInvoice)
}

/* TODO: testcase not stable
func TestPaymentObligationMint_po_successful(t *testing.T) {
	t.Parallel()
	paymentObligationMint(t, typePO)
}
*/

func paymentObligationMint(t *testing.T, documentType string) {
	alice := doctorFord.getHostTestSuite(t, "Alice")
	bob := doctorFord.getHostTestSuite(t, "Bob")

	// Alice shares document with Bob
	res := createDocument(alice.httpExpect, alice.id.String(), documentType, http.StatusOK, defaultNFTPayload(documentType, []string{bob.id.String()}))
	txID := getTransactionID(t, res)

	waitTillStatus(t, alice.httpExpect, alice.id.String(), txID, "success")

	docIdentifier := getDocumentIdentifier(t, res)
	if docIdentifier == "" {
		t.Error("docIdentifier empty")
	}

	params := map[string]interface{}{
		"document_id": docIdentifier,
		"currency":    "USD",
	}
	getDocumentAndCheck(alice.httpExpect, alice.id.String(), documentType, params)
	getDocumentAndCheck(bob.httpExpect, bob.id.String(), documentType, params)

	proofPrefix := documentType
	if proofPrefix == typePO {
		proofPrefix = poPrefix
	}

	// mint an NFT
	test := struct {
		httpStatus int
		payload    map[string]interface{}
	}{
		http.StatusOK,
		map[string]interface{}{

			"identifier":      docIdentifier,
			"registryAddress": doctorFord.getHost("Alice").config.GetContractAddress(config.PaymentObligation).String(),
			"depositAddress":  "0x186158a678f4623ca1620bc933874dda6b8f7ed8", // dummy address
			"proofFields":     []string{proofPrefix + ".gross_amount", proofPrefix + ".currency", proofPrefix + ".due_date", "collaborators[0]"},
		},
	}

	response, err := alice.host.mintNFT(alice.httpExpect, alice.id.String(), test.httpStatus, test.payload)
	txID = getTransactionID(t, response)
	waitTillStatus(t, alice.httpExpect, alice.id.String(), txID, "success")

	assert.Nil(t, err, "mintNFT should be successful")
	assert.True(t, len(response.Value("token_id").String().Raw()) > 0, "successful tokenId should have length 77")

}

func TestPaymentObligationMint_errors(t *testing.T) {
	t.Parallel()
	alice := doctorFord.getHostTestSuite(t, "Alice")
	tests := []struct {
		errorMsg   string
		httpStatus int
		payload    map[string]interface{}
	}{
		{

			"RegistryAddress is not a valid Ethereum address",
			http.StatusInternalServerError,
			map[string]interface{}{

				"registryAddress": "0x123",
			},
		},
		{
			"DepositAddress is not a valid Ethereum address",
			http.StatusInternalServerError,
			map[string]interface{}{

				"registryAddress": "0xf72855759a39fb75fc7341139f5d7a3974d4da08", //dummy address
				"depositAddress":  "abc",
			},
		},
		{
			"document not found in the system database",
			http.StatusInternalServerError,
			map[string]interface{}{

				"identifier":      "0x12121212",
				"registryAddress": "0xf72855759a39fb75fc7341139f5d7a3974d4da08", //dummy address
				"depositAddress":  "0xf72855759a39fb75fc7341139f5d7a3974d4da08", //dummy address
			},
		},
	}
	for _, test := range tests {
		t.Run(test.errorMsg, func(t *testing.T) {
			t.Parallel()
			response, err := alice.host.mintNFT(alice.httpExpect, alice.id.String(), test.httpStatus, test.payload)
			assert.Nil(t, err, "it should be possible to call the API endpoint")
			response.Value("error").String().Contains(test.errorMsg)
		})
	}
}
