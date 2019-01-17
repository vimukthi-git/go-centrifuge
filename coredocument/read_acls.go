package coredocument

import (
	"bytes"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/crypto/secp256k1"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ErrZeroCollaborators error when no collaborators are passed
	ErrZeroCollaborators = errors.Error("require at least one collaborator")

	// ErrPeerNotFound error when peer is not found in the read rules
	ErrPeerNotFound = errors.Error("peer not found")

	// nftByteCount is the length of combined bytes of registry and tokenID
	nftByteCount = 52
)

// TokenRegistry defines NFT retrieval functions.
type TokenRegistry interface {
	// OwnerOf to retrieve owner of the tokenID
	OwnerOf(registry common.Address, tokenID []byte) (common.Address, error)
}

// initReadRules initiates the read rules for a given coredocument.
// Collaborators are given Read_Sign action.
// if the rules are created already, this is a no-op.
func initReadRules(cd *coredocumentpb.CoreDocument, collabs []identity.CentID) error {
	if len(cd.Roles) > 0 && len(cd.ReadRules) > 0 {
		return nil
	}

	if len(collabs) < 1 {
		return ErrZeroCollaborators
	}

	addCollaboratorsToReadSignRules(cd, collabs)
	return nil
}

func addCollaboratorsToReadSignRules(cd *coredocumentpb.CoreDocument, collabs []identity.CentID) {
	if len(collabs) == 0 {
		return
	}

	// create a role for given collaborators
	role := new(coredocumentpb.Role)
	for _, c := range collabs {
		c := c
		role.Collaborators = append(role.Collaborators, c[:])
	}

	addNewRule(cd, role, coredocumentpb.Action_ACTION_READ_SIGN)
}

// addNewRule creates a new rule as per the role and action.
func addNewRule(cd *coredocumentpb.CoreDocument, role *coredocumentpb.Role, action coredocumentpb.Action) {
	roleKey := uint32(len(cd.Roles))
	cd.Roles = append(cd.Roles, &coredocumentpb.RoleEntry{
		RoleKey: roleKey,
		Role:    role,
	})

	rule := new(coredocumentpb.ReadRule)
	rule.Roles = append(rule.Roles, roleKey)
	rule.Action = action
	cd.ReadRules = append(cd.ReadRules, rule)
}

// addNFTToReadRules adds NFT token to the read rules of core document.
func addNFTToReadRules(cd *coredocumentpb.CoreDocument, registry common.Address, tokenID []byte) error {
	nft, err := constructNFT(registry, tokenID)
	if err != nil {
		return errors.New("failed to construct NFT: %v", err)
	}

	role := new(coredocumentpb.Role)
	role.Nfts = append(role.Nfts, nft)
	addNewRule(cd, role, coredocumentpb.Action_ACTION_READ)
	return nil
}

// constructNFT appends registry and tokenID to byte slice
func constructNFT(registry common.Address, tokenID []byte) ([]byte, error) {
	var nft []byte
	// first 20 bytes of registry
	nft = append(nft, registry.Bytes()...)

	// next 32 bytes of the tokenID
	nft = append(nft, tokenID...)

	if len(nft) != nftByteCount {
		return nil, errors.New("byte length mismatch")
	}

	return nft, nil
}

// ReadAccessValidator defines validator functions for peer.
type ReadAccessValidator interface {
	PeerCanRead(cd *coredocumentpb.CoreDocument, peer identity.CentID) error
	NFTOwnerCanRead(
		cd *coredocumentpb.CoreDocument,
		registry common.Address,
		tokenID []byte,
		signature string,
		peer identity.CentID) error
}

// readAccessValidator implements ReadAccessValidator.
type readAccessValidator struct {
	tokenRegistry TokenRegistry
}

// PeerCanRead validate if the core document can be read by the peer.
// Returns an error if not.
func (r readAccessValidator) PeerCanRead(cd *coredocumentpb.CoreDocument, peer identity.CentID) error {
	// loop though read rules
	ch := roleIterator(cd, coredocumentpb.Action_ACTION_READ_SIGN)
	for role := range ch {
		if isPeerInRole(role, peer) {
			return nil
		}
	}

	return ErrPeerNotFound
}

func getRole(key uint32, roles []*coredocumentpb.RoleEntry) (*coredocumentpb.Role, error) {
	for _, roleEntry := range roles {
		if roleEntry.RoleKey == key {
			return roleEntry.Role, nil
		}
	}

	return nil, errors.New("role %d not found", key)
}

// isPeerInRole returns true if peer is in the given role as collaborators.
func isPeerInRole(role *coredocumentpb.Role, peer identity.CentID) bool {
	for _, id := range role.Collaborators {
		if bytes.Equal(id, peer[:]) {
			return true
		}
	}

	return false
}

// peerValidator returns the ReadAccessValidator tp verify peer.
func peerValidator() ReadAccessValidator {
	return readAccessValidator{}
}

// nftValidator returns the ReadAccessValidator for nft owner verification.
func nftValidator(tr TokenRegistry) ReadAccessValidator {
	return readAccessValidator{tokenRegistry: tr}
}

// NFTOwnerCanRead checks if the nft owner/peer can read the document
// Note: signature should be calculated from the hash which is calculated as
// keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
func (r readAccessValidator) NFTOwnerCanRead(
	cd *coredocumentpb.CoreDocument,
	registry common.Address,
	tokenID []byte,
	signature string,
	peer identity.CentID) error {

	// check if the peer can read the doc
	if err := r.PeerCanRead(cd, peer); err == nil {
		return nil
	}

	// check if the nft is present in read rules
	ch := roleIterator(cd, coredocumentpb.Action_ACTION_READ)
	found := false
	for role := range ch {
		if isNFTInRole(role, registry, tokenID) {
			found = true
		}
	}

	if !found {
		return errors.New("nft missing")
	}

	// get the owner of the NFT
	owner, err := r.tokenRegistry.OwnerOf(registry, tokenID)
	if err != nil {
		return errors.New("failed to get NFT owner: %v", err)
	}

	msg, err := constructNFT(registry, tokenID)
	if err != nil {
		return err
	}

	if !secp256k1.VerifySignatureWithAddress(owner.String(), signature, msg) {
		return errors.New("peer(%s) doesn't own NFT", peer.String())
	}

	return nil
}

// roleIterator iterates through each role present in read rule
func roleIterator(cd *coredocumentpb.CoreDocument, action coredocumentpb.Action) <-chan *coredocumentpb.Role {
	ch := make(chan *coredocumentpb.Role)
	go func() {
		for _, rule := range cd.ReadRules {
			if rule.Action != action {
				continue
			}

			for _, rk := range rule.Roles {
				role, err := getRole(rk, cd.Roles)
				if err != nil {
					// seems like roles and rules are not in sync
					// skip to next one
					continue
				}

				ch <- role
			}
		}

		close(ch)
	}()

	return ch
}

// isNFTInRole checks if the given nft(registry + token) is part of the core document role.
func isNFTInRole(role *coredocumentpb.Role, registry common.Address, tokenID []byte) bool {
	enft, err := constructNFT(registry, tokenID)
	if err != nil {
		return false
	}

	for _, n := range role.Nfts {
		if bytes.Equal(n, enft) {
			return true
		}
	}

	return false
}
