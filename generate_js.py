#! /usr/bin/python3

# Requires Python 3.4.
from pathlib import Path

root = Path(__file__).parent
js_dir = root / "js"
js_file = js_dir / "main.js"
chains_dir = root / "chains"
chain_files = list(chains_dir.glob("*.json"))

if not js_dir.is_dir():
    js_dir.mkdir()

with js_file.open("w") as f:
    f.write("exports.chain_tests = require('../chain_tests.json');\n")
    f.write("exports.chain_test_inputs = {};\n")
    for chain_file in chain_files:
        f.write("exports.chain_test_inputs['{0}'] = "
                "require('../chains/{0}');\n"
                .format(chain_file.name))
