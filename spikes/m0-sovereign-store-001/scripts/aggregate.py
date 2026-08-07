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
a = p.parse_args()
root = pathlib.Path(a.root)

cases = {}

def load_json(path):
    return json.loads(path.read_text(encoding="utf-8"))

for path in root.rglob("build-*.json"):
    data = load_json(path)
    key = (data["runner_os"], data["candidate"])
    cases.setdefault(key, {})["build"] = data

for path in root.rglob("metrics-*.json"):
    data = load_json(path)
    key = ("Windows" if data["goos"] == "windows" else "Linux", data["candidate"])
    slot = "metrics_medium" if path.name.startswith("metrics-medium-") else "metrics_small"
    cases.setdefault(key, {})[slot] = data

for path in root.rglob("tests-*.json"):
    parts = path.stem.split("-")
    # tests-<RunnerOS>-<driver>
    if len(parts) < 3:
        continue
    key = (parts[1], parts[2])
    failed = []
    package_fail = False
    for line in path.read_text(encoding="utf-8", errors="replace").splitlines():
        try:
            event = json.loads(line)
        except json.JSONDecodeError:
            continue
        if event.get("Action") == "fail":
            if event.get("Test"):
                failed.append(event["Test"])
            else:
                package_fail = True
    cases.setdefault(key, {})["tests"] = {
        "passed": not failed and not package_fail,
        "failed_tests": sorted(set(failed)),
        "package_fail": package_fail,
    }

expected = [("Linux", "modernc"), ("Linux", "mattn"), ("Windows", "modernc"), ("Windows", "mattn")]
rows = []
complete = True
all_correct = True
all_receipts = True
for key in expected:
    item = cases.get(key, {})
    tests = item.get("tests")
    build = item.get("build")
    small = item.get("metrics_small")
    medium = item.get("metrics_medium")
    case_complete = bool(tests and build and small and medium)
    complete = complete and case_complete
    passed = bool(tests and tests.get("passed"))
    all_correct = all_correct and passed
    receipts = bool(
        small and medium
        and small.get("integrity_check") == "ok"
        and medium.get("integrity_check") == "ok"
        and small.get("backup_sha256")
        and small.get("restored_sha256")
        and medium.get("backup_sha256")
        and medium.get("restored_sha256")
        and small.get("recovered_snapshot", {}).get("aurora_id") == "AURORA-METRICS"
        and medium.get("recovered_snapshot", {}).get("aurora_id") == "AURORA-METRICS"
    )
    all_receipts = all_receipts and receipts
    rows.append({
        "runner_os": key[0],
        "candidate": key[1],
        "correctness_passed": passed,
        "evidence_receipts_complete": receipts,
        "failed_tests": [] if not tests else tests.get("failed_tests", []),
        "build": build,
        "metrics_small": small,
        "metrics_medium": medium,
    })

by_candidate = {}
for candidate in ("modernc", "mattn"):
    candidate_rows = [r for r in rows if r["candidate"] == candidate and r["build"] and r["metrics_small"] and r["metrics_medium"]]
    if candidate_rows:
        by_candidate[candidate] = {
            "all_platforms_correct": all(r["correctness_passed"] for r in rows if r["candidate"] == candidate),
            "all_receipts_complete": all(r["evidence_receipts_complete"] for r in rows if r["candidate"] == candidate),
            "cgo_values": sorted({r["build"]["cgo_enabled"] for r in candidate_rows}),
            "build_seconds_mean": statistics.mean(r["build"]["build_seconds"] for r in candidate_rows),
            "binary_bytes_mean": statistics.mean(r["build"]["binary_bytes"] for r in candidate_rows),
            "dependency_packages_mean": statistics.mean(r["build"]["go_dependency_packages"] for r in candidate_rows),
            "small_transition_p95_ns_mean": statistics.mean(r["metrics_small"]["transition_p95_ns"] for r in candidate_rows),
            "medium_backup_ns_mean": statistics.mean(r["metrics_medium"]["backup_ns"] for r in candidate_rows),
            "medium_restore_ns_mean": statistics.mean(r["metrics_medium"]["restore_ns"] for r in candidate_rows),
        }

payload = {
    "spike_id": "SPK-AURORA-M0-SOVEREIGN-STORE-001",
    "candidate_count": 2,
    "matrix_cases_expected": 4,
    "matrix_cases_complete": sum(1 for r in rows if r["build"] and r["metrics_small"] and r["metrics_medium"] and r["correctness_passed"]),
    "complete": complete,
    "all_correctness_passed": all_correct,
    "all_evidence_receipts_complete": all_receipts,
    "medium_fixture_bytes": 4194304,
    "cases": rows,
    "candidate_summary": by_candidate,
    "decision_rule_note": "Correctness gates dominate. If both pass, prefer lower operational/build burden and clearer cross-platform reproduction; performance is secondary.",
}
path_json = pathlib.Path(a.json_out)
path_json.parent.mkdir(parents=True, exist_ok=True)
path_json.write_text(json.dumps(payload, indent=2) + "\n", encoding="utf-8")

lines = [
    "# SPK-001 Matrix Summary",
    "",
    f"- complete: `{complete}`",
    f"- all correctness cases passed: `{all_correct}`",
    f"- all required evidence receipts complete: `{all_receipts}`",
    "- medium deterministic fixture: `4 MiB` retained state revision",
    "",
    "| OS | Candidate | Correctness | Receipts | CGO | SQLite | Build s | Binary MiB | Go deps | Small Tx p95 ms | Medium MiB | Medium Backup ms | Medium Restore ms |",
    "|---|---|---:|---:|---:|---|---:|---:|---:|---:|---:|---:|---:|",
]
for r in rows:
    b, s, m = r["build"], r["metrics_small"], r["metrics_medium"]
    if not b or not s or not m:
        lines.append(f"| {r['runner_os']} | {r['candidate']} | {'PASS' if r['correctness_passed'] else 'FAIL/INCOMPLETE'} | {'YES' if r['evidence_receipts_complete'] else 'NO'} | — | — | — | — | — | — | — | — | — |")
        continue
    lines.append(
        "| {os} | {candidate} | {correct} | {receipts} | {cgo} | {sqlite} | {build:.3f} | {binary:.2f} | {deps} | {p95:.3f} | {medium_mib:.2f} | {backup:.3f} | {restore:.3f} |".format(
            os=r["runner_os"], candidate=r["candidate"], correct="PASS" if r["correctness_passed"] else "FAIL",
            receipts="YES" if r["evidence_receipts_complete"] else "NO", cgo=b["cgo_enabled"], sqlite=s["sqlite_version"],
            build=b["build_seconds"], binary=b["binary_bytes"] / 1024 / 1024, deps=b["go_dependency_packages"],
            p95=s["transition_p95_ns"] / 1e6, medium_mib=m["database_bytes"] / 1024 / 1024,
            backup=m["backup_ns"] / 1e6, restore=m["restore_ns"] / 1e6,
        )
    )
lines += ["", "## Candidate summaries", "", "```json", json.dumps(by_candidate, indent=2), "```", ""]
path_md = pathlib.Path(a.md_out)
path_md.write_text("\n".join(lines), encoding="utf-8")
print("\n".join(lines))

if not complete or not all_correct or not all_receipts:
    sys.exit(1)
