package exportage

import (
	"bytes"
	"context"
	"errors"
	"io"

	"filippo.io/age"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type Protection struct{}
func(Protection)Protect(ctx context.Context,plaintext,secret []byte)([]byte,error){if len(secret)==0{return nil,errors.New("export passphrase cannot be empty")};select{case<-ctx.Done():return nil,ctx.Err();default:};r,err:=age.NewScryptRecipient(string(secret));if err!=nil{return nil,err};var out bytes.Buffer;w,err:=age.Encrypt(&out,r);if err!=nil{return nil,err};if _,err:=w.Write(plaintext);err!=nil{return nil,err};if err:=w.Close();err!=nil{return nil,err};return out.Bytes(),nil}
func(Protection)Unprotect(ctx context.Context,ciphertext,secret []byte)([]byte,error){if len(secret)==0{return nil,errors.New("export passphrase cannot be empty")};select{case<-ctx.Done():return nil,ctx.Err();default:};id,err:=age.NewScryptIdentity(string(secret));if err!=nil{return nil,err};id.SetMaxWorkFactor(18);r,err:=age.Decrypt(bytes.NewReader(ciphertext),id);if err!=nil{return nil,err};plain,err:=io.ReadAll(r);if err!=nil{return nil,err};return plain,nil}
var _ ports.ExportProtection=Protection{}
