from pathlib import Path

path = Path("spikes/m0-owner-trust-002/internal/trust/trust.go")
text = path.read_text(encoding="utf-8")

old_errors = '''\tErrStateRollback               = errors.New("governing state rollback detected")
\tErrStaleGeneration             = errors.New("stale expected generation")
\tErrOwnerAuthenticationRequired = errors.New("authenticated owner required")
\tErrRevalidationNotRequired     = errors.New("restore revalidation is not pending")
'''
new_errors = '''\tErrStateRollback               = errors.New("governing state rollback detected")
\tErrAnchorLag                   = errors.New("owner trust anchor lags authenticated governing state")
\tErrTimeUntrusted               = errors.New("wall time is behind authenticated high-water")
\tErrRevalidationRequired        = errors.New("restored historical state requires owner revalidation")
\tErrStaleGeneration             = errors.New("stale expected generation")
\tErrOwnerAuthenticationRequired = errors.New("authenticated owner required")
\tErrRevalidationNotRequired     = errors.New("restore revalidation is not pending")
'''
if text.count(old_errors) != 1:
    raise SystemExit("error declaration patch target mismatch")
text = text.replace(old_errors, new_errors, 1)

old_advance = '''func Advance(layout Layout, session *OwnerSession, mutation Mutation, now time.Time, hook FaultHook) (AdvanceMetrics, error) {
\tanchor, err := loadAnchor(layout.AnchorPath)
\tif err != nil {
\t\treturn AdvanceMetrics{}, err
\t}
\tif !verifyAnchor(session, anchor) {
\t\treturn AdvanceMetrics{}, ErrInvalidAnchorMAC
\t}

\tstartDB := time.Now()
'''
new_advance = '''func Advance(layout Layout, session *OwnerSession, mutation Mutation, now time.Time, hook FaultHook) (AdvanceMetrics, error) {
\t_, anchor, err := preflightOperationalMutation(layout, session, now)
\tif err != nil {
\t\treturn AdvanceMetrics{}, err
\t}

\tstartDB := time.Now()
'''
if text.count(old_advance) != 1:
    raise SystemExit("Advance patch target mismatch")
text = text.replace(old_advance, new_advance, 1)

insert_before = '''func Advance(layout Layout, session *OwnerSession, mutation Mutation, now time.Time, hook FaultHook) (AdvanceMetrics, error) {
'''
preflight = '''func preflightOperationalMutation(layout Layout, session *OwnerSession, now time.Time) (dbState, anchorRecord, error) {
\tif _, err := os.Stat(layout.RestoreMarkerPath); err == nil {
\t\treturn dbState{}, anchorRecord{}, ErrRevalidationRequired
\t} else if !os.IsNotExist(err) {
\t\treturn dbState{}, anchorRecord{}, fmt.Errorf("inspect restore marker: %w", err)
\t}

\tdb, err := openDB(layout.DBPath, false)
\tif err != nil {
\t\treturn dbState{}, anchorRecord{}, err
\t}
\tstate, err := readState(db)
\t_ = db.Close()
\tif err != nil {
\t\treturn dbState{}, anchorRecord{}, err
\t}
\tif !verifyState(session, state) {
\t\treturn dbState{}, anchorRecord{}, ErrInvalidDBMAC
\t}

\tanchor, err := loadAnchor(layout.AnchorPath)
\tif err != nil {
\t\treturn dbState{}, anchorRecord{}, err
\t}
\tif !verifyAnchor(session, anchor) {
\t\treturn dbState{}, anchorRecord{}, ErrInvalidAnchorMAC
\t}
\tif state.Generation < anchor.Generation {
\t\treturn dbState{}, anchorRecord{}, ErrStateRollback
\t}
\tif state.Generation > anchor.Generation {
\t\treturn dbState{}, anchorRecord{}, ErrAnchorLag
\t}
\tif now.Unix() < anchor.WallTimeHighWaterUnix {
\t\treturn dbState{}, anchorRecord{}, ErrTimeUntrusted
\t}
\treturn state, anchor, nil
}

'''
if text.count(insert_before) != 1:
    raise SystemExit("preflight insertion target mismatch")
text = text.replace(insert_before, preflight + insert_before, 1)

path.write_text(text, encoding="utf-8")
