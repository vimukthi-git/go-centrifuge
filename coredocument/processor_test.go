// +build unit

package coredocument

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/signatures"
	"github.com/centrifuge/go-centrifuge/testingutils/commons"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/ed25519"
)

var dp defaultProcessor

func TestCoreDocumentProcessor_SendNilDocument(t *testing.T) {
	err := dp.Send(nil, nil, [identity.CentIDLength]byte{})
	assert.Error(t, err, "should have thrown an error")
}

type mockModel struct {
	mock.Mock
	documents.Model
}

func (m mockModel) PackCoreDocument() (*coredocumentpb.CoreDocument, error) {
	args := m.Called()
	cd, _ := args.Get(0).(*coredocumentpb.CoreDocument)
	return cd, args.Error(1)
}

func (m mockModel) UnpackCoreDocument(cd *coredocumentpb.CoreDocument) error {
	args := m.Called(cd)
	return args.Error(0)
}

func TestDefaultProcessor_PrepareForSignatureRequests(t *testing.T) {
	// pack failed
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("error")).Once()
	err := dp.PrepareForSignatureRequests(model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to pack core document")

	cd := new(coredocumentpb.CoreDocument)
	model = mockModel{}

	// failed to get id
	pub, _ := config.Config().GetSigningKeyPair()
	config.Config().Set("keys.signing.publicKey", "wrong path")
	cd = New()
	cd.DataRoot = utils.RandomSlice(32)
	cd.EmbeddedData = &any.Any{
		TypeUrl: "some type",
		Value:   []byte("some data"),
	}
	assert.Nil(t, FillSalts(cd))
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = dp.PrepareForSignatureRequests(model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get keys for signing")
	config.Config().Set("keys.signing.publicKey", pub)

	// failed unpack
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	model.On("UnpackCoreDocument", cd).Return(fmt.Errorf("error")).Once()
	err = dp.PrepareForSignatureRequests(model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unpack the core document")

	// success
	cd.Signatures = nil
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	model.On("UnpackCoreDocument", cd).Return(nil).Once()
	err = dp.PrepareForSignatureRequests(model)
	model.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, cd.Signatures)
	assert.Len(t, cd.Signatures, 1)
	sig := cd.Signatures[0]
	id, err := identity.GetIdentityConfig()
	assert.Nil(t, err)
	assert.True(t, ed25519.Verify(id.Keys[identity.KeyPurposeSigning].PublicKey, cd.SigningRoot, sig.Signature))
}

type p2pClient struct {
	mock.Mock
	client
}

func (p p2pClient) GetSignaturesForDocument(ctx context.Context, doc *coredocumentpb.CoreDocument) error {
	args := p.Called(ctx, doc)
	return args.Error(0)
}

func TestDefaultProcessor_RequestSignatures(t *testing.T) {
	// pack failed
	ctx := context.Background()
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("error")).Once()
	err := dp.RequestSignatures(ctx, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to pack core document")

	// validations failed
	cd := new(coredocumentpb.CoreDocument)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(4)
	err = dp.RequestSignatures(ctx, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate model for signature request")

	// failed signature collection
	cd = New()
	cd.DataRoot = utils.RandomSlice(32)
	cd.EmbeddedData = &any.Any{
		TypeUrl: "some type",
		Value:   []byte("some data"),
	}
	assert.Nil(t, FillSalts(cd))
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	model.On("UnpackCoreDocument", cd).Return(nil).Once()
	err = dp.PrepareForSignatureRequests(model)
	assert.Nil(t, err)
	model.AssertExpectations(t)
	c := p2pClient{}
	c.On("GetSignaturesForDocument", ctx, cd).Return(fmt.Errorf("error")).Once()
	dp.P2PClient = c
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(4)
	err = dp.RequestSignatures(ctx, model)
	model.AssertExpectations(t)
	c.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to collect signatures from the collaborators")

	// unpack fail
	c = p2pClient{}
	c.On("GetSignaturesForDocument", ctx, cd).Return(nil).Once()
	dp.P2PClient = c
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(4)
	model.On("UnpackCoreDocument", cd).Return(fmt.Errorf("error")).Once()
	err = dp.RequestSignatures(ctx, model)
	model.AssertExpectations(t)
	c.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unpack core document")

	// success
	c = p2pClient{}
	c.On("GetSignaturesForDocument", ctx, cd).Return(nil).Once()
	dp.P2PClient = c
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(4)
	model.On("UnpackCoreDocument", cd).Return(nil).Once()
	err = dp.RequestSignatures(ctx, model)
	model.AssertExpectations(t)
	c.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestDefaultProcessor_PrepareForAnchoring(t *testing.T) {
	// pack failed
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("error")).Once()
	err := dp.PrepareForAnchoring(model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to pack core document")

	// failed validations
	cd := new(coredocumentpb.CoreDocument)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(4)
	err = dp.PrepareForAnchoring(model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate signatures")

	// failed unpack
	cd = New()
	cd.DataRoot = utils.RandomSlice(32)
	cd.EmbeddedData = &any.Any{
		TypeUrl: "some type",
		Value:   []byte("some data"),
	}
	assert.Nil(t, FillSalts(cd))
	err = CalculateSigningRoot(cd)
	assert.Nil(t, err)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(4)
	model.On("UnpackCoreDocument", cd).Return(fmt.Errorf("error")).Once()
	c, err := identity.GetIdentityConfig()
	assert.Nil(t, err)
	s := signatures.Sign(c, identity.KeyPurposeSigning, cd.SigningRoot)
	cd.Signatures = []*coredocumentpb.Signature{s}
	pubkey, err := utils.SliceToByte32(c.Keys[identity.KeyPurposeSigning].PublicKey)
	assert.Nil(t, err)
	idkey := &identity.EthereumIdentityKey{
		Key:       pubkey,
		Purposes:  []*big.Int{big.NewInt(identity.KeyPurposeSigning)},
		RevokedAt: big.NewInt(0),
	}
	id := &testingcommons.MockID{}
	srv := &testingcommons.MockIDService{}
	srv.On("LookupIdentityForID", c.ID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	err = dp.PrepareForAnchoring(model)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	id.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unpack core document")

	// success
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(4)
	model.On("UnpackCoreDocument", cd).Return(nil).Once()
	id = &testingcommons.MockID{}
	srv = &testingcommons.MockIDService{}
	srv.On("LookupIdentityForID", c.ID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	err = dp.PrepareForAnchoring(model)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	id.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, cd.DocumentRoot)
}

type mockRepo struct {
	mock.Mock
	anchors.AnchorRepository
}

func (m mockRepo) CommitAnchor(anchorID anchors.AnchorID, documentRoot anchors.DocumentRoot, centrifugeID identity.CentID, documentProofs [][32]byte, signature []byte) (<-chan *anchors.WatchCommit, error) {
	args := m.Called(anchorID, documentRoot, centrifugeID, documentProofs, signature)
	c, _ := args.Get(0).(chan *anchors.WatchCommit)
	return c, args.Error(1)
}

func (m mockRepo) GetDocumentRootOf(anchorID anchors.AnchorID) (anchors.DocumentRoot, error) {
	args := m.Called(anchorID)
	docRoot, _ := args.Get(0).(anchors.DocumentRoot)
	return docRoot, args.Error(1)
}

func TestDefaultProcessor_AnchorDocument(t *testing.T) {
	// pack failed
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("error")).Once()
	err := dp.AnchorDocument(model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to pack core document")

	// validations failed
	model = mockModel{}
	cd := new(coredocumentpb.CoreDocument)
	model.On("PackCoreDocument").Return(cd, nil).Times(5)
	err = dp.AnchorDocument(model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "pre anchor validation failed")

	// get ID failed
	cd = New()
	cd.DataRoot = utils.RandomSlice(32)
	cd.EmbeddedData = &any.Any{
		TypeUrl: "some type",
		Value:   []byte("some data"),
	}
	assert.Nil(t, FillSalts(cd))
	assert.Nil(t, CalculateSigningRoot(cd))
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(5)
	c, err := identity.GetIdentityConfig()
	assert.Nil(t, err)
	s := signatures.Sign(c, identity.KeyPurposeSigning, cd.SigningRoot)
	cd.Signatures = []*coredocumentpb.Signature{s}
	assert.Nil(t, CalculateDocumentRoot(cd))
	pubkey, err := utils.SliceToByte32(c.Keys[identity.KeyPurposeSigning].PublicKey)
	assert.Nil(t, err)
	idkey := &identity.EthereumIdentityKey{
		Key:       pubkey,
		Purposes:  []*big.Int{big.NewInt(identity.KeyPurposeSigning)},
		RevokedAt: big.NewInt(0),
	}
	id := &testingcommons.MockID{}
	srv := &testingcommons.MockIDService{}
	srv.On("LookupIdentityForID", c.ID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	oldID := config.Config().GetString("identityId")
	config.Config().Set("identityId", "wrong id")
	err = dp.AnchorDocument(model)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	id.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get self cent ID")
	config.Config().Set("identityId", "0x0102030405060708")

	// wrong ID
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(5)
	srv.On("LookupIdentityForID", c.ID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	err = dp.AnchorDocument(model)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	id.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "centID invalid")
	config.Config().Set("identityId", oldID)

	// missing eth keys
	oldPth := config.Config().Get("keys.ethauth.publicKey")
	config.Config().Set("keys.ethauth.publicKey", "wrong path")
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(5)
	srv.On("LookupIdentityForID", c.ID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	err = dp.AnchorDocument(model)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	id.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get eth keys")
	config.Config().Set("keys.ethauth.publicKey", oldPth)

	// failed anchor commit
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(5)
	srv.On("LookupIdentityForID", c.ID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	repo := mockRepo{}
	repo.On("CommitAnchor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error")).Once()
	dp.AnchorRepository = repo
	err = dp.AnchorDocument(model)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	id.AssertExpectations(t)
	repo.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to commit anchor")

	// success
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(5)
	srv.On("LookupIdentityForID", c.ID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	repo = mockRepo{}
	ch := make(chan *anchors.WatchCommit, 1)
	ch <- new(anchors.WatchCommit)
	repo.On("CommitAnchor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ch, nil).Once()
	dp.AnchorRepository = repo
	err = dp.AnchorDocument(model)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	id.AssertExpectations(t)
	repo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestDefaultProcessor_SendDocument(t *testing.T) {
	ctx := context.Background()

	// pack failed
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("error")).Once()
	err := dp.SendDocument(ctx, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to pack core document")

	// failed validations
	model = mockModel{}
	cd := new(coredocumentpb.CoreDocument)
	model.On("PackCoreDocument").Return(cd, nil).Times(6)
	err = dp.SendDocument(ctx, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "post anchor validations failed")

	// failed send
	cd = New()
	cd.DataRoot = utils.RandomSlice(32)
	cd.EmbeddedData = &any.Any{
		TypeUrl: "some type",
		Value:   []byte("some data"),
	}
	cd.Collaborators = [][]byte{[]byte("some id")}
	assert.Nil(t, FillSalts(cd))
	assert.Nil(t, CalculateSigningRoot(cd))
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Times(6)
	c, err := identity.GetIdentityConfig()
	assert.Nil(t, err)
	s := signatures.Sign(c, identity.KeyPurposeSigning, cd.SigningRoot)
	cd.Signatures = []*coredocumentpb.Signature{s}
	assert.Nil(t, CalculateDocumentRoot(cd))
	pubkey, err := utils.SliceToByte32(c.Keys[identity.KeyPurposeSigning].PublicKey)
	assert.Nil(t, err)
	idkey := &identity.EthereumIdentityKey{
		Key:       pubkey,
		Purposes:  []*big.Int{big.NewInt(identity.KeyPurposeSigning)},
		RevokedAt: big.NewInt(0),
	}
	id := &testingcommons.MockID{}
	srv := &testingcommons.MockIDService{}
	srv.On("LookupIdentityForID", c.ID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	docRoot, err := anchors.ToDocumentRoot(cd.DocumentRoot)
	assert.Nil(t, err)
	repo := mockRepo{}
	repo.On("GetDocumentRootOf", mock.Anything).Return(docRoot, nil).Once()
	dp.AnchorRepository = repo
	err = dp.SendDocument(ctx, model)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	id.AssertExpectations(t)
	repo.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid length byte slice provided for centID")
}
