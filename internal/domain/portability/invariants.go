package portability

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/gowebpki/jcs"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

var(ErrUnsupportedVersion=errors.New("unsupported sovereign export version");ErrDigestMismatch=errors.New("sovereign export digest mismatch"))
const SchemaV1=`{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://aurora.local/schemas/sovereign-export-v1.schema.json","type":"object","additionalProperties":false,"required":["format","version","export_id","created_at","governing_generation","aurora","projects","authority","transition_attempts","records","owner_recovery","integrity"],"properties":{"format":{"const":"aurora-sovereign-export"},"version":{"const":1},"export_id":{"type":"string","minLength":1},"created_at":{"type":"string","format":"date-time"},"governing_generation":{"type":"integer","minimum":1},"aurora":{"type":"object"},"projects":{"type":"array"},"authority":{"type":"object"},"transition_attempts":{"type":"array"},"records":{"type":"array"},"owner_recovery":{"type":"object"},"integrity":{"type":"object","required":["payload_sha256"],"properties":{"payload_sha256":{"type":"string","pattern":"^[0-9a-f]{64}$"}},"additionalProperties":false}}}`
var(schemaOnce sync.Once;compiled *jsonschema.Schema;compileErr error)
func exportSchema()(*jsonschema.Schema,error){schemaOnce.Do(func(){resource,err:=jsonschema.UnmarshalJSON(strings.NewReader(SchemaV1));if err!=nil{compileErr=err;return};c:=jsonschema.NewCompiler();if err:=c.AddResource("sovereign-export-v1.schema.json",resource);err!=nil{compileErr=err;return};compiled,compileErr=c.Compile("sovereign-export-v1.schema.json")});return compiled,compileErr}
func NewExportID()(string,error){var raw[16]byte;if _,err:=rand.Read(raw[:]);err!=nil{return "",err};return "EXP-"+hex.EncodeToString(raw[:]),nil}
func Finalize(doc *Document)error{if doc.Format==""{doc.Format=FormatV1};if doc.Version==0{doc.Version=1};if doc.ExportID==""{id,err:=NewExportID();if err!=nil{return err};doc.ExportID=id};normalizeCollections(doc);digest,err:=Digest(*doc);if err!=nil{return err};doc.Integrity.PayloadSHA256=digest;return validateSchema(*doc)}
func Verify(doc Document)error{if doc.Format!=FormatV1||doc.Version!=1{return ErrUnsupportedVersion};if err:=validateSchema(doc);err!=nil{return err};want,err:=Digest(doc);if err!=nil{return err};if !strings.EqualFold(want,doc.Integrity.PayloadSHA256){return ErrDigestMismatch};return nil}
func normalizeCollections(doc *Document){if doc.Projects==nil{doc.Projects=[]ProjectBundle{}};if doc.Authority.Revisions==nil{doc.Authority.Revisions=[]authority.State{}};if doc.Attempts==nil{doc.Attempts=[]TransitionAttempt{}};if doc.Records==nil{doc.Records=[]Record{}}}
func Digest(doc Document)(string,error){view:=struct{Format string `json:"format"`;Version int `json:"version"`;ExportID string `json:"export_id"`;CreatedAt any `json:"created_at"`;GoverningGeneration uint64 `json:"governing_generation"`;Aurora any `json:"aurora"`;Projects any `json:"projects"`;Authority any `json:"authority"`;Attempts any `json:"transition_attempts"`;Records any `json:"records"`;OwnerRecovery any `json:"owner_recovery"`}{doc.Format,doc.Version,doc.ExportID,doc.CreatedAt,doc.GoverningGeneration,doc.Aurora,doc.Projects,doc.Authority,doc.Attempts,doc.Records,doc.OwnerRecovery};raw,err:=json.Marshal(view);if err!=nil{return "",err};canonical,err:=jcs.Transform(raw);if err!=nil{return "",err};sum:=sha256.Sum256(canonical);return hex.EncodeToString(sum[:]),nil}
func validateSchema(doc Document)error{raw,err:=json.Marshal(doc);if err!=nil{return err};v,err:=jsonschema.UnmarshalJSON(bytes.NewReader(raw));if err!=nil{return err};schema,err:=exportSchema();if err!=nil{return err};if err:=schema.Validate(v);err!=nil{return fmt.Errorf("sovereign export schema: %w",err)};return nil}
