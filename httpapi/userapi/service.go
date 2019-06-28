package userapi

import (
	"context"

	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/extensions"
	"github.com/centrifuge/go-centrifuge/httpapi"
	"github.com/centrifuge/go-centrifuge/jobs"
)

// Service provides functionality for User APIs.
type Service struct {
	coreService            httpapi.CoreService
	transferDetailsService extensions.TransferDetailService
}

// TODO: this can be refactored into a generic Service which handles all kinds of custom attributes

// CreateTransferDetail creates and anchors a Transfer Detail
func (s Service) CreateTransferDetail(ctx context.Context, req extensions.CreateTransferDetailRequest) (documents.Model, jobs.JobID, error) {
	return s.transferDetailsService.CreateTransferDetail(ctx, req)
}

// UpdateTransferDetail updates and anchors a Transfer Detail
func (s Service) UpdateTransferDetail(ctx context.Context, req extensions.UpdateTransferDetailRequest) (documents.Model, jobs.JobID, error) {
	return s.transferDetailsService.UpdateTransferDetail(ctx, req)
}

// GetCurrentTransferDetail returns the current version on a Transfer Detail
func (s Service) GetCurrentTransferDetail(ctx context.Context, docID, transferID []byte) (*extensions.TransferDetail, documents.Model, error) {
	model, err := s.coreService.GetDocument(ctx, docID)
	if err != nil {
		return nil, nil, err
	}
	data, model, err := s.transferDetailsService.DeriveTransferDetail(ctx, model, transferID)
	if err != nil {
		return nil, nil, err
	}

	return data, model, nil
}

// GetCurrentTransferDetailsList returns a list of Transfer Details on the current version of a document
func (s Service) GetCurrentTransferDetailsList(ctx context.Context, docID []byte) (*extensions.TransferDetailList, documents.Model, error) {
	model, err := s.coreService.GetDocument(ctx, docID)
	if err != nil {
		return nil, nil, err
	}

	data, model, err := s.transferDetailsService.DeriveTransferList(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return data, model, nil
}

// GetVersionTransferDetail returns a Transfer Detail on a particular version of a Document
func (s Service) GetVersionTransferDetail(ctx context.Context, docID, versionID, transferID []byte) (*extensions.TransferDetail, documents.Model, error) {
	model, err := s.coreService.GetDocumentVersion(ctx, docID, versionID)
	if err != nil {
		return nil, nil, err
	}

	data, model, err := s.transferDetailsService.DeriveTransferDetail(ctx, model, transferID)
	if err != nil {
		return nil, nil, err
	}

	return data, model, nil
}

// GetVersionTransferDetailsList returns a list of Transfer Details on a particular version of a Document
func (s Service) GetVersionTransferDetailsList(ctx context.Context, docID, versionID []byte) (*extensions.TransferDetailList, documents.Model, error) {
	model, err := s.coreService.GetDocumentVersion(ctx, docID, versionID)
	if err != nil {
		return nil, nil, err
	}

	data, model, err := s.transferDetailsService.DeriveTransferList(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return data, model, nil
}
