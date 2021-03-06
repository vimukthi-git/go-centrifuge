package invoice

import (
	"errors"
	"fmt"

	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/coredocument"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/p2p"
)

type Bootstrapper struct{}

// Bootstrap sets the required storage and registers
func (*Bootstrapper) Bootstrap(context map[string]interface{}) error {
	if _, ok := context[bootstrap.BootstrappedLevelDb]; !ok {
		return errors.New("initializing LevelDB repository failed")
	}

	// register service
	srv := DefaultService(getRepository(), coredocument.DefaultProcessor(identity.IDService, p2p.NewP2PClient(), anchors.GetAnchorRepository()), anchors.GetAnchorRepository())
	err := documents.GetRegistryInstance().Register(documenttypes.InvoiceDataTypeUrl, srv)
	if err != nil {
		return fmt.Errorf("failed to register invoice service: %v", err)
	}

	return nil
}
