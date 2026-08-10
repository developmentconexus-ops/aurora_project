package application

import (
	"context"
	"crypto/hmac"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type TrustState string

const (
	TrustNormal               TrustState = "NORMAL"
	TrustAnchorLag            TrustState = "ANCHOR_LAG"
	TrustStateRollback        TrustState = "STATE_ROLLBACK"
	TrustInvalidDBMAC         TrustState = "INVALID_DB_MAC"
	TrustInvalidAnchorMAC     TrustState = "INVALID_ANCHOR_MAC"
	TrustMissingAnchor        TrustState = "MISSING_ANCHOR"
	TrustTimeUntrusted        TrustState = "TIME_UNTRUSTED"
	TrustRevalidationRequired TrustState = "REVALIDATION_REQUIRED"
)

type trustSignals struct {
	dbMACValid, anchorPresent, anchorMACValid bool
	dbGeneration, anchorGeneration            uint64
	now, highWater                            time.Time
	revalidationRequired                      bool
}

func classifyTrust(in trustSignals) TrustState {
	if !in.dbMACValid { return TrustInvalidDBMAC }
	if !in.anchorPresent { return TrustMissingAnchor }
	if !in.anchorMACValid { return TrustInvalidAnchorMAC }
	if in.dbGeneration > in.anchorGeneration { return TrustAnchorLag }
	if in.dbGeneration < in.anchorGeneration { return TrustStateRollback }
	if in.revalidationRequired { return TrustRevalidationRequired }
	if !in.highWater.IsZero() && in.now.Before(in.highWater) { return TrustTimeUntrusted }
	return TrustNormal
}

type TrustAssessment struct { State TrustState; Current ports.CurrentSnapshot; Anchor *ports.Anchor }
type TrustStateError struct{ State TrustState }
func (e TrustStateError) Error() string { return fmt.Sprintf("owner trust state is %s", e.State) }

func (s *Service) AssessTrust(ctx context.Context, pass []byte) (TrustAssessment,error) {
	current,err:=s.State.LoadCurrent(ctx);if err!=nil{return TrustAssessment{},err}
	root,err:=s.Trust.LoadRootEnvelope(ctx);if err!=nil{return TrustAssessment{},err}
	ork,err:=unlockORK(root,pass);if err!=nil{return TrustAssessment{},err};defer clear(ork)
	wantDB,macErr:=governingMAC(ork,governingSnapshot(current));dbValid:=macErr==nil&&hmac.Equal(wantDB,current.GoverningHMAC)
	assessment:=TrustAssessment{Current:current}
	if !dbValid { assessment.State=TrustInvalidDBMAC; return assessment,nil }
	anchor,err:=s.Trust.LoadAnchor(ctx)
	if err!=nil { if errors.Is(err,ports.ErrAnchorNotFound){assessment.State=TrustMissingAnchor;return assessment,nil};assessment.State=TrustInvalidAnchorMAC;return assessment,nil }
	assessment.Anchor=&anchor
	anchorBytes,err:=base64.RawURLEncoding.DecodeString(anchor.HMAC);anchorValid:=err==nil
	if anchorValid { wantAnchor,macErr:=anchorMAC(ork,anchor);anchorValid=macErr==nil&&hmac.Equal(wantAnchor,anchorBytes)&&anchor.RootID==root.RootID&&anchor.AuroraID==current.AuroraID }
	revalidation:=false
	if auth,err:=decodeAuthorityState(current.AuthorityJSON);err==nil { revalidation=auth.RevalidationRequired } else { assessment.State=TrustInvalidDBMAC;return assessment,nil }
	assessment.State=classifyTrust(trustSignals{dbMACValid:true,anchorPresent:true,anchorMACValid:anchorValid,dbGeneration:current.GoverningGeneration,anchorGeneration:anchor.GoverningGeneration,now:s.Clock.Now().UTC(),highWater:anchor.ObservedWallTimeHighWater,revalidationRequired:revalidation})
	return assessment,nil
}

func (s *Service) requireNormalTrust(ctx context.Context,pass []byte)(TrustAssessment,error){assessment,err:=s.AssessTrust(ctx,pass);if err!=nil{return TrustAssessment{},err};if assessment.State!=TrustNormal{return assessment,TrustStateError{State:assessment.State}};return assessment,nil}

func (s *Service) ReconcileAnchor(ctx context.Context,pass []byte)(InspectResult,error){assessment,err:=s.AssessTrust(ctx,pass);if err!=nil{return InspectResult{},err};if assessment.State!=TrustAnchorLag||assessment.Anchor==nil{return InspectResult{},TrustStateError{State:assessment.State}};root,err:=s.Trust.LoadRootEnvelope(ctx);if err!=nil{return InspectResult{},err};ork,err:=unlockORK(root,pass);if err!=nil{return InspectResult{},err};defer clear(ork);anchor:=*assessment.Anchor;anchor.GoverningGeneration=assessment.Current.GoverningGeneration;now:=s.Clock.Now().UTC();if now.After(anchor.ObservedWallTimeHighWater){anchor.ObservedWallTimeHighWater=now};mac,err:=anchorMAC(ork,anchor);if err!=nil{return InspectResult{},err};anchor.HMAC=base64.RawURLEncoding.EncodeToString(mac);if err:=s.Trust.PublishAnchor(ctx,anchor);err!=nil{return InspectResult{},err};return s.Inspect(ctx,pass)
}
