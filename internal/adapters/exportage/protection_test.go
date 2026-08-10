package exportage

import (
	"bytes"
	"context"
	"testing"
)

func TestProtectRoundTripWrongSecretAndTamper(t *testing.T) {
	p:=Protection{};ctx:=context.Background();plain:=[]byte(`{"format":"aurora-sovereign-export"}`);secret:=[]byte("export-secret")
	ciphertext,err:=p.Protect(ctx,plain,secret);if err!=nil{t.Fatal(err)}
	if bytes.Contains(ciphertext,plain){t.Fatal("ciphertext contains plaintext")}
	got,err:=p.Unprotect(ctx,ciphertext,secret);if err!=nil{t.Fatal(err)};if !bytes.Equal(got,plain){t.Fatalf("got=%q want=%q",got,plain)}
	if _,err:=p.Unprotect(ctx,ciphertext,[]byte("wrong"));err==nil{t.Fatal("wrong export secret decrypted package")}
	tampered:=append([]byte(nil),ciphertext...);tampered[len(tampered)-1]^=0xff;if _,err:=p.Unprotect(ctx,tampered,secret);err==nil{t.Fatal("tampered age package decrypted")}
}
