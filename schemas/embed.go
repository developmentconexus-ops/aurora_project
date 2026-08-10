package schemas

import _ "embed"

// ProjectStateV1 is the accepted M0 state-envelope schema.
//
//go:embed project-state-v1.schema.json
var ProjectStateV1 []byte
