#!/usr/bin/env python3

# Requires Python 3.4.
from pathlib import Path

root = Path(__file__).parent
chains_dir = root / "chains"
chain_tests_file = root / "chain_tests.json"
chain_files = sorted(list(chains_dir.glob("*.json")))

# Generate JS

js_dir = root / "js"
if not js_dir.is_dir():
    js_dir.mkdir()
js_file = js_dir / "main.js"

with js_file.open("w") as f:
    f.write("exports.chain_tests = require('../chain_tests.json');\n")
    f.write("exports.chain_test_inputs = {};\n")
    for chain_file in chain_files:
        f.write("exports.chain_test_inputs['{0}'] = "
                "require('../chains/{0}');\n"
                .format(chain_file.name))

# Generate Go -- now done via bin/generate (in TS)
