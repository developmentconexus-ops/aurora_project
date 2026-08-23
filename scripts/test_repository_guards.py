#!/usr/bin/env python3
from __future__ import annotations

import tempfile
import unittest
from pathlib import Path

import generate_docs
import validate_docs


class RepositoryGuardRegressionTests(unittest.TestCase):
    def test_generated_blueprint_uses_repository_relative_source_paths(self) -> None:
        content = generate_docs.generate_blueprint(generate_docs.ROOT)
        self.assertNotIn(str(generate_docs.ROOT), content)
        self.assertIn("docs/product/blueprint/01-product-vision.md", content)

    def test_generated_blueprint_is_stable_across_lf_and_crlf_source_bytes(self) -> None:
        with tempfile.TemporaryDirectory() as td:
            root = Path(td)
            source_dir = root / "docs/product/blueprint"
            source_dir.mkdir(parents=True)
            for i in range(1, 16):
                name = source_dir / f"{i:02d}-section.md"
                text = f"---\nid: TEST-{i:02d}\n---\n\n# Section {i}\n\nBody\n"
                name.write_bytes(text.encode("utf-8"))
            lf = generate_docs.generate_blueprint(root)
            for path in source_dir.glob("*.md"):
                path.write_bytes(path.read_bytes().replace(b"\n", b"\r\n"))
            crlf = generate_docs.generate_blueprint(root)
            self.assertEqual(lf, crlf)

    def test_link_validation_includes_index_and_roadmap(self) -> None:
        with tempfile.TemporaryDirectory() as td:
            root = Path(td)
            (root / "docs").mkdir()
            (root / "docs/index.md").write_text("[missing index target](missing-index.md)\n", encoding="utf-8")
            (root / "docs/roadmap.md").write_text("[missing roadmap target](missing-roadmap.md)\n", encoding="utf-8")
            errors: list[str] = []
            validate_docs.validate_links(root, errors)
            self.assertTrue(any("docs/index.md broken link" in error for error in errors), errors)
            self.assertTrue(any("docs/roadmap.md broken link" in error for error in errors), errors)

    def test_program_owner_guard_rejects_duplicate_owner(self) -> None:
        with tempfile.TemporaryDirectory() as td:
            root = Path(td)
            (root / "docs").mkdir()
            (root / "docs/roadmap.md").write_text("program_status_authority: true\n", encoding="utf-8")
            (root / "docs/other.md").write_text("program_status_authority: true\n", encoding="utf-8")
            errors: list[str] = []
            validate_docs.guard_program_owner(root, errors)
            self.assertTrue(errors)

    def test_program_owner_guard_rejects_missing_owner(self) -> None:
        with tempfile.TemporaryDirectory() as td:
            root = Path(td)
            (root / "docs").mkdir()
            (root / "docs/roadmap.md").write_text("no owner flag\n", encoding="utf-8")
            errors: list[str] = []
            validate_docs.guard_program_owner(root, errors)
            self.assertTrue(errors)

    def test_generated_staleness_guard_rejects_perturbed_projection(self) -> None:
        with tempfile.TemporaryDirectory() as td_root, tempfile.TemporaryDirectory() as td_generated:
            root = Path(td_root)
            generated = Path(td_generated)
            (root / "docs/product").mkdir(parents=True)
            (generated / "docs/product").mkdir(parents=True)
            (root / "docs/product/PRODUCT-BLUEPRINT.md").write_text("actual\n", encoding="utf-8")
            (generated / "docs/product/PRODUCT-BLUEPRINT.md").write_text("expected\n", encoding="utf-8")
            errors: list[str] = []
            validate_docs.validate_generated(root, generated, errors)
            self.assertIn("stale generated Product Blueprint", errors)

    def test_router_guard_rejects_missing_required_route(self) -> None:
        with tempfile.TemporaryDirectory() as td:
            root = Path(td)
            (root / "docs").mkdir()
            (root / "docs/index.md").write_text("roadmap.md\n", encoding="utf-8")
            errors: list[str] = []
            validate_docs.validate_router(root, errors)
            self.assertTrue(any("docs/index missing required route" in error for error in errors), errors)


if __name__ == "__main__":
    unittest.main()
