package application

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/exportage"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
)

func TestMigrateProtectedPackageV0ToV1(t *testing.T) {
	ctx := context.Background()
	raw, err := os.ReadFile("../../testdata/migration/v0-valid.json")
	if err != nil { t.Fatal(err) }
	secret := []byte("migration-export-secret")
	protection := exportage.Protection{}
	ciphertext, err := protection.Protect(ctx, raw, secret)
	if err != nil { t.Fatal(err) }
	svc := &Service{ExportProtection: protection}
	result, err := svc.MigratePackage(ctx, ciphertext, secret)
	if err != nil { t.Fatal(err) }
	plain, err := protection.Unprotect(ctx, result.Ciphertext, secret)
	if err != nil { t.Fatal(err) }
	var doc portability.Document
	if err := json.Unmarshal(plain, &doc); err != nil { t.Fatal(err) }
	if err := portability.Verify(doc); err != nil { t.Fatal(err) }
	if result.SourceVersion != 0 || result.TargetVersion != 1 || result.PayloadSHA256 != doc.Integrity.PayloadSHA256 { t.Fatalf("result=%+v doc=%+v", result, doc) }
}
