// +build integration

package testingbootstrap

import (
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/context/testlogging"
	"github.com/centrifuge/go-centrifuge/documents/invoice"
	"github.com/centrifuge/go-centrifuge/documents/purchaseorder"
	"github.com/centrifuge/go-centrifuge/ethereum"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/nft"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/centrifuge/go-centrifuge/storage"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("context")

var bootstappers = []bootstrap.TestBootstrapper{
	&testlogging.TestLoggingBootstrapper{},
	&config.Bootstrapper{},
	&storage.Bootstrapper{},
	&ethereum.Bootstrapper{},
	&anchors.Bootstrapper{},
	&identity.Bootstrapper{},
	&invoice.Bootstrapper{},
	&purchaseorder.Bootstrapper{},
	&nft.Bootstrapper{},
	&queue.Bootstrapper{},
}

func TestFunctionalEthereumBootstrap() {
	contextval := map[string]interface{}{}
	for _, b := range bootstappers {
		err := b.TestBootstrap(contextval)
		if err != nil {
			log.Error("Error encountered while bootstrapping", err)
			panic(err)
		}
	}
}
func TestFunctionalEthereumTearDown() {
	for _, b := range bootstappers {
		err := b.TestTearDown()
		if err != nil {
			log.Error("Error encountered while bootstrapping", err)
			panic(err)
		}
	}
}
