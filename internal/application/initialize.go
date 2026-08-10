package application

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

const ownerOperatorID = "OWNER-LOCAL"
var ErrIncompleteBootstrap = errors.New("incomplete Aurora bootstrap requires explicit recovery")

type InitializeResult struct { AuroraID string `json:"aurora_id"`; OwnerOperatorID string `json:"owner_operator_id"`; GoverningGeneration uint64 `json:"governing_generation"` }
type initialAuthorityState struct { Revision uint64 `json:"revision"`; PredecessorRevision *uint64 `json:"predecessor_revision"`; Grants []initialAuthorityGrant `json:"grants"`; RevalidationRequired bool `json:"revalidation_required"`; ChangedBy string `json:"changed_by"`; ChangedAt time.Time `json:"changed_at"` }
type initialAuthorityGrant struct { AuthorityID string `json:"authority_id"`; AuthorityRevision uint64 `json:"authority_revision"`; SubjectOperatorID string `json:"subject_operator_id"`; ActorID string `json:"actor_id"`; ProjectScope []string `json:"project_scope"`; PermittedActionClasses []string `json:"permitted_action_classes"`; Conditions map[string]string `json:"conditions"`; ValidFrom *time.Time `json:"valid_from"`; ValidUntil *time.Time `json:"valid_until"`; LifecycleStatus string `json:"lifecycle_status"`; IssuedAt time.Time `json:"issued_at"`; RevokedAt *time.Time `json:"revoked_at"`; Supersedes string `json:"supersedes"`; Provenance string `json:"provenance"` }
type bootstrapGoverningSnapshot struct { Version int `json:"version"`; Aurora bootstrapAurora `json:"aurora"`; GoverningGeneration uint64 `json:"governing_generation"`; Projects []governingProjectSnapshot `json:"projects"`; Authority json.RawMessage `json:"authority"` }
type bootstrapAurora struct { AuroraID string `json:"aurora_id"`; OwnerOperatorID string `json:"owner_operator_id"` }

func (s *Service) Initialize(ctx context.Context, passphrase []byte) (InitializeResult,error) {
	if s.State==nil||s.Trust==nil||s.Clock==nil{return InitializeResult{},errors.New("service dependencies are incomplete")}
	current,stateErr:=s.State.LoadCurrent(ctx);root,rootErr:=s.Trust.LoadRootEnvelope(ctx);stateExists:=stateErr==nil;rootExists:=rootErr==nil
	if stateExists||rootExists{if stateExists&&rootExists{return InitializeResult{},ports.ErrAlreadyInitialized};return InitializeResult{},ErrIncompleteBootstrap}
	if !errors.Is(stateErr,ports.ErrNotInitialized){return InitializeResult{},stateErr};if !errors.Is(rootErr,ports.ErrRootNotFound){return InitializeResult{},rootErr};_ = current;_ = root
	env,ork,err:=newRootEnvelope(passphrase);if err!=nil{return InitializeResult{},err};defer clear(ork)
	if err:=s.Trust.StoreRootEnvelope(ctx,env);err!=nil{return InitializeResult{},fmt.Errorf("publish owner root: %w",err)}
	auroraID,err:=identity.NewAuroraID();if err!=nil{return InitializeResult{},err};now:=s.Clock.Now().UTC();authorityID,err:=randomIdentifier("AUTH-");if err!=nil{return InitializeResult{},err}
	authority:=initialAuthorityState{Revision:1,Grants:[]initialAuthorityGrant{{AuthorityID:authorityID,AuthorityRevision:1,SubjectOperatorID:ownerOperatorID,ActorID:"AURORA_CORE",ProjectScope:[]string{},PermittedActionClasses:[]string{"PROJECT_CREATE","PROJECT_STATE_MUTATE","AUTHORITY_ADMIN","EXPORT","RESTORE","MIGRATE"},Conditions:map[string]string{},LifecycleStatus:"ACTIVE",IssuedAt:now,Provenance:"OWNER_BOOTSTRAP"}},ChangedBy:ownerOperatorID,ChangedAt:now}
	authorityJSON,err:=json.Marshal(authority);if err!=nil{return InitializeResult{},err}
	snapshot:=bootstrapGoverningSnapshot{Version:1,Aurora:bootstrapAurora{AuroraID:string(auroraID),OwnerOperatorID:ownerOperatorID},GoverningGeneration:1,Projects:[]governingProjectSnapshot{},Authority:authorityJSON}
	mac,err:=governingMAC(ork,snapshot);if err!=nil{return InitializeResult{},err};operationID,err:=randomIdentifier("OP-");if err!=nil{return InitializeResult{},err}
	if _,err:=s.State.Bootstrap(ctx,ports.BootstrapMutation{OperationID:operationID,AuroraID:string(auroraID),OwnerOperatorID:ownerOperatorID,CreatedAt:now,AuthorityRevision:1,AuthorityJSON:authorityJSON,GoverningGeneration:1,GoverningDescriptorHMAC:mac});err!=nil{return InitializeResult{},err}
	anchor:=ports.Anchor{Version:1,RootID:env.RootID,AuroraID:string(auroraID),GoverningGeneration:1,ObservedWallTimeHighWater:now};anchorBytes,err:=anchorMAC(ork,anchor);if err!=nil{return InitializeResult{},err};anchor.HMAC=base64.RawURLEncoding.EncodeToString(anchorBytes)
	if err:=s.Trust.PublishAnchor(ctx,anchor);err!=nil{return InitializeResult{},fmt.Errorf("publish owner anchor: %w",err)}
	return InitializeResult{AuroraID:string(auroraID),OwnerOperatorID:ownerOperatorID,GoverningGeneration:1},nil
}
