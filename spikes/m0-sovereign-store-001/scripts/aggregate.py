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
    cases.setdefault(key, {})["metrics"] = data
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
for key in expected:
    item = cases.get(key, {})
    tests = item.get("tests")
    build = item.get("build")
    metrics = item.get("metrics")
    if not tests or not build or not metrics:
        complete = False
    passed = bool(tests and tests.get("passed"))
    all_correct = all_correct and passed
    rows.append({
        "runner_os": key[0],
        "candidate": key[1],
        "correctness_passed": passed,
        "failed_tests": [] if not tests else tests.get("failed_tests", []),
        "build": build,
        "metrics": metrics,
    })

by_candidate = {}
for candidate in ("modernc", "mattn"):
    candidate_rows = [r for r in rows if r["candidate"] == candidate and r["build"] and r["metrics"]]
    if candidate_rows:
        by_candidate[candidate] = {
            "all_platforms_correct": all(r["correctness_passed"] for r in rows if r["candidate"] == candidate),
            "cgo_values": sorted({r["build"]["cgo_enabled"] for r in candidate_rows}),
            "build_seconds_mean": statistics.mean(r["build"]["build_seconds"] for r in candidate_rows),
            "binary_bytes_mean": statistics.mean(r["build"]["binary_bytes"] for r in candidate_rows),
            "dependency_packages_mean": statistics.mean(r["build"]["go_dependency_packages"] for r in candidate_rows),
            "transition_p95_ns_mean": statistics.mean(r["metrics"]["transition_p95_ns"] for r in candidate_rows),
        }

payload = {
    "spike_id": "SPK-AURORA-M0-SOVEREIGN-STORE-001",
    "candidate_count": 2,
    "matrix_cases_expected": 4,
    "matrix_cases_complete": sum(1 for r in rows if r["build"] and r["metrics"] and (r["correctness_passed"] or r["failed_tests"])),
    "complete": complete,
    "all_correctness_passed": all_correct,
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
    "",
    "| OS | Candidate | Correctness | CGO | SQLite | Build s | Binary MiB | Go deps | Tx p95 ms | Backup ms | Restore ms |",
    "|---|---|---:|---:|---|---:|---:|---:|---:|---:|---:|",
]
for r in rows:
    b, m = r["build"], r["metrics"]
    if not b or not m:
        lines.append(f"| {r['runner_os']} | {r['candidate']} | {'PASS' if r['correctness_passed'] else 'FAIL/INCOMPLETE'} | — | — | — | — | — | — | — | — |")
        continue
    lines.append(
        "| {os} | {candidate} | {correct} | {cgo} | {sqlite} | {build:.3f} | {binary:.2f} | {deps} | {p95:.3f} | {backup:.3f} | {restore:.3f} |".format(
            os=r["runner_os"], candidate=r["candidate"], correct="PASS" if r["correctness_passed"] else "FAIL",
            cgo=b["cgo_enabled"], sqlite=m["sqlite_version"], build=b["build_seconds"],
            binary=b["binary_bytes"] / 1024 / 1024, deps=b["go_dependency_packages"],
            p95=m["transition_p95_ns"] / 1e6, backup=m["backup_ns"] / 1e6, restore=m["restore_ns"] / 1e6,
        )
    )
lines += ["", "## Candidate summaries", "", "```json", json.dumps(by_candidate, indent=2), "```", ""]
path_md = pathlib.Path(a.md_out)
path_md.write_text("\n".join(lines), encoding="utf-8")
print("\n".join(lines))

if not complete or not all_correct:
    sys.exit(1)
