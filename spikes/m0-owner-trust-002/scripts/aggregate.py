#!/usr/bin/env python3
import argparse
import json
import pathlib
import statistics
import sys

p = argparse.ArgumentParser()
p.add_argument("--root", required=True)
p.add_argument("--json-out", required=True)
p.add_argument("--md-out", required=True)
p.add_argument("--state-out", required=True)
a = p.parse_args()
root = pathlib.Path(a.root)

required_scenarios = [f"S{i:02d}_" for i in range(1, 13)]
required_classification_tests = [
    "DB_generation_equals_anchor",
    "DB_generation_greater_than_anchor",
    "DB_generation_less_than_anchor",
    "invalid_DB_MAC",
    "invalid_anchor_MAC",
    "missing_anchor",
    "missing_wrapped_root",
]
expected_classes = {
    "NORMAL",
    "STATE_ROLLBACK",
    "ANCHOR_LAG",
    "INVALID_DB_MAC",
    "INVALID_ANCHOR_MAC",
    "MISSING_ANCHOR",
    "MISSING_WRAPPED_ROOT",
    "OWNER_UNLOCK_FAILED",
    "TIME_UNTRUSTED",
    "REVALIDATION_REQUIRED",
}


def read_test_events(path: pathlib.Path):
    passed = set()
    failed = set()
    package_failed = False
    for line in path.read_text(encoding="utf-8", errors="replace").splitlines():
        try:
            event = json.loads(line)
        except json.JSONDecodeError:
            continue
        action = event.get("Action")
        test = event.get("Test")
        if action == "pass" and test:
            passed.add(test)
        elif action == "fail" and test:
            failed.add(test)
        elif action == "fail" and not test:
            package_failed = True
    return passed, failed, package_failed


cases = []
for metrics_path in sorted(root.rglob("metrics-*.json")):
    metrics = json.loads(metrics_path.read_text(encoding="utf-8"))
    os_name = "Windows" if metrics["goos"] == "windows" else "Linux"
    tests_path = next(iter(root.rglob(f"tests-{os_name}.json")), None)
    if tests_path is None:
        cases.append({"os": os_name, "metrics": metrics, "complete": False, "reason": "missing test log"})
        continue
    passed, failed, package_failed = read_test_events(tests_path)
    scenario_results = {}
    for scenario in required_scenarios:
        names = [t for t in passed if t.startswith("TestSPK002Scenarios/") and scenario in t]
        scenario_results[scenario[:3]] = bool(names) and not any(scenario in t for t in failed)
    classification_results = {}
    for name in required_classification_tests:
        full = f"TestRecoveryClassificationMatrix/{name}"
        classification_results[name] = full in passed and full not in failed

    diagnostics = set(metrics.get("diagnostic_samples", {}).keys())
    metric_checks = {
        "cgo_disabled": metrics.get("cgo_enabled") == "0",
        "argon2_memory_64mib": metrics.get("kdf", {}).get("memory_kib") == 64 * 1024,
        "argon2_iterations_3": metrics.get("kdf", {}).get("iterations") == 3,
        "argon2_parallelism_4": metrics.get("kdf", {}).get("parallelism") == 4,
        "argon2_peak_observed": metrics.get("kdf", {}).get("observed_peak_heap_delta_bytes", 0) > 0,
        "root_stable_after_rotation": metrics.get("root_fingerprint_stable_after_rotation") is True,
        "restore_fail_closed": metrics.get("restored_classification") == "REVALIDATION_REQUIRED",
        "owner_revalidation_restores_normal": metrics.get("post_revalidation_classification") == "NORMAL" and metrics.get("post_revalidation_permitting") is True,
        "all_diagnostics_present": expected_classes.issubset(diagnostics),
        "positive_root_artifact": metrics.get("root_artifact_bytes", 0) > 0,
        "positive_anchor_artifact": metrics.get("anchor_artifact_bytes", 0) > 0,
        "positive_state_db": metrics.get("state_db_bytes", 0) > 0,
        "positive_recovery_bundle": metrics.get("recovery_bundle_bytes", 0) > 0,
    }
    complete = (
        not package_failed
        and not failed
        and all(scenario_results.values())
        and all(classification_results.values())
        and all(metric_checks.values())
    )
    cases.append({
        "os": os_name,
        "complete": complete,
        "package_failed": package_failed,
        "failed_tests": sorted(failed),
        "scenario_results": scenario_results,
        "classification_results": classification_results,
        "metric_checks": metric_checks,
        "metrics": metrics,
    })

cases_by_os = {c["os"]: c for c in cases}
complete_matrix = set(cases_by_os) == {"Linux", "Windows"} and all(c["complete"] for c in cases)

payload = {
    "spike_id": "SPK-AURORA-M0-OWNER-TRUST-002",
    "execution_candidate": "random ORK + Argon2id KEK + AES-256-GCM + HKDF-SHA-256 + HMAC-SHA-256 + external trust high-water",
    "expected_operating_systems": ["Linux", "Windows"],
    "complete": complete_matrix,
    "cases": cases,
    "decision_rule_note": "Candidate A passes only if S01-S12, all recovery classifications and evidence checks pass cross-platform. Candidate B remains a reference alternative unless Candidate A leaves material rotation/recovery uncertainty.",
}

json_out = pathlib.Path(a.json_out)
json_out.parent.mkdir(parents=True, exist_ok=True)
json_out.write_text(json.dumps(payload, indent=2) + "\n", encoding="utf-8")

lines = [
    "# SPK-002 Evidence Summary",
    "",
    f"- complete cross-platform matrix: `{complete_matrix}`",
    "- candidate: random ORK + Argon2id-wrapped root + HKDF/HMAC + external authenticated trust high-water",
    "",
    "| OS | S01–S12 | Recovery classes | KDF ms | Observed heap MiB | Bootstrap ms | Rotation ms | DB commit ms | Anchor ms | Restore posture | Post-revalidation |",
    "|---|---:|---:|---:|---:|---:|---:|---:|---:|---|---|",
]
for os_name in ("Linux", "Windows"):
    case = cases_by_os.get(os_name)
    if not case:
        lines.append(f"| {os_name} | MISSING | MISSING | — | — | — | — | — | — | — | — |")
        continue
    m = case["metrics"]
    scenarios_ok = all(case["scenario_results"].values())
    classes_ok = all(case["classification_results"].values())
    lines.append(
        "| {os} | {scenarios} | {classes} | {kdf:.1f} | {heap:.1f} | {bootstrap:.1f} | {rotation:.1f} | {db:.2f} | {anchor:.2f} | {restore} | {post} |".format(
            os=os_name,
            scenarios="PASS" if scenarios_ok else "FAIL",
            classes="PASS" if classes_ok else "FAIL",
            kdf=m["kdf"]["duration_ns"] / 1e6,
            heap=m["kdf"]["observed_peak_heap_delta_bytes"] / 1024 / 1024,
            bootstrap=m["bootstrap_ns"] / 1e6,
            rotation=m["passphrase_rotation_ns"] / 1e6,
            db=m["db_commit_ns"] / 1e6,
            anchor=m["anchor_write_ns"] / 1e6,
            restore=m["restored_classification"],
            post=f"{m['post_revalidation_classification']}/{m['post_revalidation_permitting']}",
        )
    )

lines += [
    "",
    "## Security/result checks",
    "",
]
for os_name in ("Linux", "Windows"):
    case = cases_by_os.get(os_name)
    if not case:
        continue
    lines.append(f"### {os_name}")
    lines.append("")
    for key, value in case["metric_checks"].items():
        lines.append(f"- {key}: `{'PASS' if value else 'FAIL'}`")
    lines.append("")

lines += [
    "## Limitations",
    "",
    "- Fault evidence is process-kill based; it does not emulate all physical power-loss/storage-controller failures.",
    "- A purely local trust architecture cannot defeat replay of all trust files together with compromise of owner secrets/runtime.",
    "- Argon2 peak memory is an approximate Go heap observation; configured memory remains the normative 64 MiB parameter.",
    "",
]
md_out = pathlib.Path(a.md_out)
md_out.write_text("\n".join(lines), encoding="utf-8")

state_lines = [
    "# SPK-002 Recovery Classification Model",
    "",
    "These rows are covered by `TestRecoveryClassificationMatrix` plus S07/S08/S12.",
    "",
    "| Observed condition | Required classification | Permitting |",
    "|---|---|---:|",
    "| valid DB generation N / valid anchor N | NORMAL | authority-dependent |",
    "| valid DB generation N+1 / valid anchor N | ANCHOR_LAG | no |",
    "| valid DB generation N / valid anchor N+1 | STATE_ROLLBACK | no |",
    "| invalid governing-state HMAC | INVALID_DB_MAC | no |",
    "| invalid anchor HMAC | INVALID_ANCHOR_MAC | no |",
    "| current anchor missing outside restore flow | MISSING_ANCHOR | no |",
    "| wrapped owner root missing | MISSING_WRAPPED_ROOT | no |",
    "| wall clock behind authenticated high-water | TIME_UNTRUSTED | no |",
    "| historical/fresh-machine restore with authentic root+state but no current anchor | REVALIDATION_REQUIRED | no |",
    "| authenticated owner creates new post-restore authority revision and current anchor | NORMAL | yes when authority ACTIVE/unexpired |",
    "",
]
path_state = pathlib.Path(a.state_out)
path_state.write_text("\n".join(state_lines), encoding="utf-8")

print("\n".join(lines))
if not complete_matrix:
    sys.exit(1)
