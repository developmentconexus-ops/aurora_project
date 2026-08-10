package application

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type CreateProjectInput struct{DisplayLabel string `json:"display_label"`;ObjectiveSummary string `json:"objective_summary"`}
type ProjectView struct{Project project.Project `json:"project"`;CurrentState *project.ProjectStateRevision `json:"current_state,omitempty"`}
func(s *Service)CreateProject(ctx context.Context,ownerPassphrase []byte,in CreateProjectInput)(project.Project,error){if strings.TrimSpace(in.DisplayLabel)==""||len(in.DisplayLabel)>200{return project.Project{},errors.New("display label must be 1..200 characters")};if strings.TrimSpace(in.ObjectiveSummary)==""||len(in.ObjectiveSummary)>2000{return project.Project{},errors.New("objective summary must be 1..2000 characters")};trusted,err:=s.loadTrustedCurrent(ctx,ownerPassphrase);if err!=nil{return project.Project{},err};defer zero(trusted.ORK);id,err:=project.NewProjectID();if err!=nil{return project.Project{},err};now:=s.Clock.Now().UTC();p:=project.Project{ProjectID:id,DisplayLabel:in.DisplayLabel,ObjectiveSummary:in.ObjectiveSummary,CreatedAt:now,UpdatedAt:now};newSnap:=trusted.Snapshot;newSnap.Projects=append(append([]project.Project(nil),trusted.Snapshot.Projects...),p);newSnap.GoverningGeneration++;mac,err:=governingMAC(trusted.ORK,newSnap);if err!=nil{return project.Project{},err};opID:="OP-PROJECT-CREATE-"+string(id);created,err:=s.State.CreateProject(ctx,ports.CreateProjectMutation{Project:p,ExpectedGeneration:trusted.Snapshot.GoverningGeneration,NewGeneration:newSnap.GoverningGeneration,GoverningMAC:mac,OperationID:opID});if err!=nil{return project.Project{},fmt.Errorf("persist Project: %w",err)};anchor:=trusted.Anchor;anchor.GoverningGeneration=newSnap.GoverningGeneration;if now.After(anchor.ObservedWallTimeHighWater){anchor.ObservedWallTimeHighWater=now};am,err:=anchorMAC(trusted.ORK,anchor);if err!=nil{return project.Project{},err};anchor.HMAC=encodeMAC(am);if err:=s.Trust.PublishAnchor(ctx,anchor);err!=nil{return project.Project{},fmt.Errorf("publish Project anchor: %w",err)};return created,nil}
func(s *Service)ShowProject(ctx context.Context,ownerPassphrase []byte,id project.ProjectID)(ProjectView,error){return s.InspectProject(ctx,ownerPassphrase,id)}
func(s *Service)InspectProject(ctx context.Context,ownerPassphrase []byte,id project.ProjectID)(ProjectView,error){trusted,err:=s.loadTrustedCurrent(ctx,ownerPassphrase);if err!=nil{return ProjectView{},err};defer zero(trusted.ORK);p,err:=s.State.GetProject(ctx,id);if err!=nil{return ProjectView{},err};state,err:=s.State.GetCurrentProjectState(ctx,id);if err!=nil{return ProjectView{},err};return ProjectView{Project:p,CurrentState:state},nil}
