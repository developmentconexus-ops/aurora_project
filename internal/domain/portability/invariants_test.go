package portability

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
)

func TestFinalizeAndVerifyLogicalExportDigest(t *testing.T) {
	doc := Document{Format:FormatV1, Version:1, CreatedAt:time.Unix(1,0).UTC(), GoverningGeneration:7, Aurora:identity.AuroraIdentity{AuroraID:"AUR-1",OwnerOperatorID:"OWNER-LEANDRO",CreatedAt:time.Unix(1,0).UTC(),IdentityRevision:1,CapabilityContractVersion:"0.2.0"}, Authority:authority.State{Revision:3,Grants:[]authority.Grant{},ChangedBy:"OWNER-LEANDRO",ChangedAt:time.Unix(1,0).UTC()}, OwnerRecovery:RootEnvelope{Version:1,RootID:"ROOT-1",KDF:"argon2id",MemoryKiB:65536,Iterations:3,Parallelism:4,Salt:"AA",Nonce:"BB",WrappedORK:"CC"}}
	if err := Finalize(&doc); err != nil { t.Fatal(err) }
	if doc.Integrity.PayloadSHA256 == "" { t.Fatal("missing export digest") }
	if err := Verify(doc); err != nil { t.Fatal(err) }
	mutated := doc
	mutated.Aurora.AuroraID = "AUR-tampered"
	if err := Verify(mutated); err == nil { t.Fatal("semantic mutation unexpectedly passed digest verification") }
	unsupported := doc; unsupported.Version=99
	if err := Verify(unsupported); err == nil { t.Fatal("unsupported export version accepted") }
	raw, _ := json.Marshal(doc)
	if len(raw)==0 { t.Fatal("export JSON empty") }
}
