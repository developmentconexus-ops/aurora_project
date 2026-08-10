package application

import(
	"context";"encoding/json";"errors";"time"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
)
type ExportResult struct{ExportID string `json:"export_id"`;CreatedAt time.Time `json:"created_at"`;PayloadSHA256 string `json:"payload_sha256"`;Ciphertext []byte `json:"-"`}
func(s *Service)Export(ctx context.Context,ownerPassphrase,exportSecret []byte)(ExportResult,error){if s.ExportProtection==nil{return ExportResult{},errors.New("export protection adapter is not configured")};trusted,err:=s.loadTrustedCurrent(ctx,ownerPassphrase);if err!=nil{return ExportResult{},err};defer zero(trusted.ORK);state,err:=s.State.ExportLogicalState(ctx);if err!=nil{return ExportResult{},err};doc:=portability.Document{Format:portability.FormatV1,Version:1,CreatedAt:s.Clock.Now().UTC(),GoverningGeneration:state.GoverningGeneration,Aurora:state.Aurora,Projects:state.Projects,Authority:state.Authority,Attempts:state.Attempts,Records:state.Records,OwnerRecovery:rootToPortable(trusted.Root)};if err:=portability.Finalize(&doc);err!=nil{return ExportResult{},err};plain,err:=json.Marshal(doc);if err!=nil{return ExportResult{},err};ciphertext,err:=s.ExportProtection.Protect(ctx,plain,exportSecret);if err!=nil{return ExportResult{},err};return ExportResult{ExportID:doc.ExportID,CreatedAt:doc.CreatedAt,PayloadSHA256:doc.Integrity.PayloadSHA256,Ciphertext:ciphertext},nil}
