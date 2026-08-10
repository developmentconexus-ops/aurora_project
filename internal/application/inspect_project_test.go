package application

import (
	"context"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
)

func TestInspectProjectWithoutAcceptedStateReportsExplicitAbsence(t *testing.T) {
	ctx:=context.Background(); root:=t.TempDir(); store,err:=sqlite.Open(root); if err!=nil{t.Fatal(err)}; defer store.Close()
	svc:=&Service{State:store,Trust:trustfs.New(root),Clock:fixedClock{now:time.Unix(100,0).UTC()}}; pass:=[]byte("fixture-passphrase")
	if _,err:=svc.Initialize(ctx,pass);err!=nil{t.Fatal(err)}
	p,err:=svc.CreateProject(ctx,pass,CreateProjectInput{DisplayLabel:"P",ObjectiveSummary:"Objective"});if err!=nil{t.Fatal(err)}
	view,err:=svc.InspectProject(ctx,pass,p.ProjectID);if err!=nil{t.Fatal(err)}
	if view.Project.ProjectID!=p.ProjectID || view.CurrentState!=nil { t.Fatalf("view=%+v",view) }
}
