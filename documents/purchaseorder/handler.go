package purchaseorder

import (
	"fmt"

	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/go-centrifuge/centerrors"
	"github.com/centrifuge/go-centrifuge/code"
	"github.com/centrifuge/go-centrifuge/documents"
	clientpurchaseorderpb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/purchaseorder"
	"github.com/ethereum/go-ethereum/common/hexutil"
	logging "github.com/ipfs/go-log"
	"golang.org/x/net/context"
)

var apiLog = logging.Logger("purchaseorder-api")

// grpcHandler handles all the purchase order document related actions
// anchoring, sending, finding stored purchase order document
type grpcHandler struct {
	service Service
}

// GRPCHandler returns an implementation of the purchaseorder DocumentServiceServer
func GRPCHandler() (clientpurchaseorderpb.DocumentServiceServer, error) {
	srv, err := documents.GetRegistryInstance().LocateService(documenttypes.PurchaseOrderDataTypeUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch purchase order service")
	}

	return grpcHandler{
		service: srv.(Service),
	}, nil
}

// Create validates the purchase order, persists it to DB, and anchors it the chain
func (h grpcHandler) Create(ctx context.Context, req *clientpurchaseorderpb.PurchaseOrderCreatePayload) (*clientpurchaseorderpb.PurchaseOrderResponse, error) {
	ctxh, err := documents.NewContextHeader()
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.New(code.Unknown, err.Error())
	}

	doc, err := h.service.DeriveFromCreatePayload(req, ctxh)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	// validate, persist, and anchor
	doc, err = h.service.Create(ctx, doc)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	return h.service.DerivePurchaseOrderResponse(doc)
}

// Update handles the document update and anchoring
func (h grpcHandler) Update(ctx context.Context, payload *clientpurchaseorderpb.PurchaseOrderUpdatePayload) (*clientpurchaseorderpb.PurchaseOrderResponse, error) {
	ctxHeader, err := documents.NewContextHeader()
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.New(code.Unknown, fmt.Sprintf("failed to get header: %v", err))
	}

	doc, err := h.service.DeriveFromUpdatePayload(payload, ctxHeader)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	doc, err = h.service.Update(ctx, doc)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	return h.service.DerivePurchaseOrderResponse(doc)
}

// GetVersion returns the requested version of a purchase order
func (h grpcHandler) GetVersion(ctx context.Context, req *clientpurchaseorderpb.GetVersionRequest) (*clientpurchaseorderpb.PurchaseOrderResponse, error) {
	identifier, err := hexutil.Decode(req.Identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.Wrap(err, "identifier is invalid")
	}
	version, err := hexutil.Decode(req.Version)
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.Wrap(err, "version is invalid")
	}
	model, err := h.service.GetVersion(identifier, version)
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.Wrap(err, "document not found")
	}
	resp, err := h.service.DerivePurchaseOrderResponse(model)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}
	return resp, nil
}

// Get returns the purchase order the latest version of the document with given identifier
func (h grpcHandler) Get(ctx context.Context, getRequest *clientpurchaseorderpb.GetRequest) (*clientpurchaseorderpb.PurchaseOrderResponse, error) {
	identifier, err := hexutil.Decode(getRequest.Identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.Wrap(err, "identifier is an invalid hex string")
	}
	model, err := h.service.GetCurrentVersion(identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.Wrap(err, "document not found")
	}
	resp, err := h.service.DerivePurchaseOrderResponse(model)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}
	return resp, nil
}
