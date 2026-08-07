#!/usr/bin/env python3
import argparse
import json
import os
import pathlib
import subprocess
import time

p = argparse.ArgumentParser()
p.add_argument("--driver", required=True)
p.add_argument("--out", required=True)
p.add_argument("--binary", required=True)
a = p.parse_args()

out = pathlib.Path(a.out)
out.parent.mkdir(parents=True, exist_ok=True)
binary = pathlib.Path(a.binary)
binary.parent.mkdir(parents=True, exist_ok=True)

start = time.perf_counter()
subprocess.run(
    ["go", "build", "-tags", a.driver, "-o", str(binary), "./cmd/spike-runner"],
    check=True,
)
build_seconds = time.perf_counter() - start
modules = subprocess.check_output(["go", "list", "-deps", "-tags", a.driver, "./..."], text=True).splitlines()

payload = {
    "candidate": a.driver,
    "runner_os": os.environ.get("RUNNER_OS", "unknown"),
    "runner_arch": os.environ.get("RUNNER_ARCH", "unknown"),
    "cgo_enabled": os.environ.get("CGO_ENABLED", ""),
    "build_seconds": build_seconds,
    "binary_bytes": binary.stat().st_size,
    "go_dependency_packages": len([x for x in modules if x.strip()]),
}
out.write_text(json.dumps(payload, indent=2) + "\n", encoding="utf-8")
print(json.dumps(payload, indent=2))
