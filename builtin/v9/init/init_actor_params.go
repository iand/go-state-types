package init

import (
	addr "github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
)

type ConstructorParams struct {
	NetworkName string
}

type ExecParams struct {
	CodeCID           cid.Cid `checked:"true"` // invalid CIDs won't get committed to the state tree
	ConstructorParams []byte
}

type ExecReturn struct {
	IDAddress     addr.Address // The canonical ID-based address for the actor.
	RobustAddress addr.Address // A more expensive but re-org-safe address for the newly created actor.
}

type InstallParams struct {
	Code []byte
}

type InstallReturn struct {
	CodeCid   cid.Cid
	Installed bool
}
