package invoice

import (
	"fmt"

	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/go-centrifuge/centerrors"
	"github.com/centrifuge/go-centrifuge/code"
	"github.com/centrifuge/go-centrifuge/documents"
	clientinvoicepb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/invoice"
	"github.com/ethereum/go-ethereum/common/hexutil"
	logging "github.com/ipfs/go-log"
	"golang.org/x/net/context"
)

var apiLog = logging.Logger("invoice-api")

// grpcHandler handles all the invoice document related actions
// anchoring, sending, finding stored invoice document
type grpcHandler struct {
	service Service
}

// GRPCHandler returns an implementation of invoice.DocumentServiceServer
func GRPCHandler() (clientinvoicepb.DocumentServiceServer, error) {
	invoiceService, err := documents.GetRegistryInstance().LocateService(documenttypes.InvoiceDataTypeUrl)
	if err != nil {
		return nil, err
	}
	return &grpcHandler{
		service: invoiceService.(Service),
	}, nil
}

// Create handles the creation of the invoices and anchoring the documents on chain
func (h *grpcHandler) Create(ctx context.Context, req *clientinvoicepb.InvoiceCreatePayload) (*clientinvoicepb.InvoiceResponse, error) {
	ctxHeader, err := documents.NewContextHeader()
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.New(code.Unknown, fmt.Sprintf("failed to get header: %v", err))
	}

	doc, err := h.service.DeriveFromCreatePayload(req, ctxHeader)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	// validate and persist
	doc, err = h.service.Create(ctx, doc)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	return h.service.DeriveInvoiceResponse(doc)
}

// Update handles the document update and anchoring
func (h *grpcHandler) Update(ctx context.Context, payload *clientinvoicepb.InvoiceUpdatePayload) (*clientinvoicepb.InvoiceResponse, error) {
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

	return h.service.DeriveInvoiceResponse(doc)
}

// GetVersion returns the requested version of the document
func (h *grpcHandler) GetVersion(ctx context.Context, getVersionRequest *clientinvoicepb.GetVersionRequest) (*clientinvoicepb.InvoiceResponse, error) {
	identifier, err := hexutil.Decode(getVersionRequest.Identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.Wrap(err, "identifier is invalid")
	}
	version, err := hexutil.Decode(getVersionRequest.Version)
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.Wrap(err, "version is invalid")
	}
	model, err := h.service.GetVersion(identifier, version)
	if err != nil {
		apiLog.Error(err)
		return nil, centerrors.Wrap(err, "document not found")
	}
	resp, err := h.service.DeriveInvoiceResponse(model)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}
	return resp, nil
}

// Get returns the invoice the latest version of the document with given identifier
func (h *grpcHandler) Get(ctx context.Context, getRequest *clientinvoicepb.GetRequest) (*clientinvoicepb.InvoiceResponse, error) {
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
	resp, err := h.service.DeriveInvoiceResponse(model)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}
	return resp, nil
}
