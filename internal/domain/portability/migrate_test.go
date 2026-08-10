package portability

import (
	"errors"
	"os"
	"testing"
)

func TestMigrateV0ToV1PreservesIdentityStateAuthorityAndProvenance(t *testing.T) {
	raw, err := os.ReadFile("../../../testdata/migration/v0-valid.json")
	if err != nil { t.Fatal(err) }
	result, err := Migrate(raw)
	if err != nil { t.Fatal(err) }
	if result.SourceVersion != 0 || result.TargetVersion != 1 { t.Fatalf("versions=%+v", result) }
	doc := result.Document
	if err := Verify(doc); err != nil { t.Fatal(err) }
	if doc.Aurora.AuroraID != "AUR-fixture-v0" { t.Fatalf("Aurora identity changed: %s", doc.Aurora.AuroraID) }
	if len(doc.Projects) != 1 || doc.Projects[0].Project.ProjectID != "PRJ-fixture-v0" { t.Fatalf("Projects=%+v", doc.Projects) }
	if len(doc.Projects[0].Revisions) != 1 || doc.Projects[0].Revisions[0].State.Summary != "Legacy accepted state" { t.Fatalf("state=%+v", doc.Projects[0].Revisions) }
	if doc.Authority.CurrentRevision != 1 || len(doc.Authority.Revisions) != 1 || doc.Authority.Revisions[0].Grants[0].Provenance != "fixture-v0-owner" { t.Fatalf("authority=%+v", doc.Authority) }
	if len(doc.Attempts) != 1 || len(doc.Records) != 1 || doc.Records[0].OperationID != "ATT-v0-1" { t.Fatalf("provenance attempts=%+v records=%+v", doc.Attempts, doc.Records) }
}

func TestMigrateRejectsSemanticMutationFixture(t *testing.T) {
	raw, err := os.ReadFile("../../../testdata/migration/v0-semantic-mutation.json")
	if err != nil { t.Fatal(err) }
	_, err = Migrate(raw)
	if !errors.Is(err, ErrMigrationSemanticInvariant) { t.Fatalf("err=%v want ErrMigrationSemanticInvariant", err) }
}

func TestMigrateRejectsUnsupportedVersion(t *testing.T) {
	_, err := Migrate([]byte(`{"format":"aurora-sovereign-export","version":99}`))
	if !errors.Is(err, ErrUnsupportedVersion) { t.Fatalf("err=%v want ErrUnsupportedVersion", err) }
}
