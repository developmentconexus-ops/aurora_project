package portability

import (
	"testing"
	"time"
)

func TestFinalizeAndVerifyExportDetectsSemanticCorruption(t *testing.T) {
	doc:=ExportDocument{Format:"aurora-sovereign-export",Version:1,CreatedAt:time.Date(2026,8,10,1,0,0,0,time.UTC),Aurora:AuroraExport{AuroraID:"AUR-F",OwnerOperatorID:"OWNER-LOCAL",GoverningGeneration:4},Projects:[]ProjectExport{{ProjectID:"PRJ-F",DisplayLabel:"Original"}},Authority:AuthorityExport{CurrentRevision:2}}
	final,err:=Finalize(doc);if err!=nil{t.Fatal(err)}
	if final.Integrity.PayloadSHA256==""{t.Fatal("missing digest")}
	if err:=Verify(final);err!=nil{t.Fatalf("valid export rejected: %v",err)}
	final.Projects[0].DisplayLabel="Tampered"
	if err:=Verify(final);err==nil{t.Fatal("tampered logical export unexpectedly verified")}
}

func TestVerifyExportRejectsUnsupportedVersion(t *testing.T){doc:=ExportDocument{Format:"aurora-sovereign-export",Version:2};if err:=Verify(doc);err==nil{t.Fatal("unsupported export version accepted")}}
