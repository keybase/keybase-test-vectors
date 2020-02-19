"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs_1 = require("fs");
const path_1 = require("path");
const fs_2 = require("fs");
const glob_1 = __importDefault(require("glob"));
const implode_1 = require("./implode");
const makeDir = (d) => {
    return new Promise((resolve, reject) => {
        fs_1.access(d, err => {
            if (err && err.code === 'ENOENT') {
                fs_1.mkdir(d, err => {
                    if (err) {
                        reject(err);
                    }
                    else {
                        resolve();
                    }
                });
            }
            else if (err) {
                reject(err);
            }
            else {
                resolve();
            }
        });
    });
};
const globFiles = async (pattern) => {
    return new Promise((resolve, reject) => {
        glob_1.default(pattern, (err, matches) => {
            if (err) {
                reject(err);
            }
            else {
                resolve(matches.sort());
            }
        });
    });
};
class Buf {
    constructor() {
        this.b = [];
    }
    str(s) {
        this.buf(Buffer.from(s, 'utf8'));
    }
    buf(b) {
        this.b.push(b);
    }
    finish() {
        return Buffer.concat(this.b);
    }
}
const toGoStringLiteral = (f) => '`' + JSON.stringify(f).replace('`', '`+"`"+`') + '`';
const run = async () => {
    const root = path_1.dirname(__dirname);
    const goDir = root + '/go';
    await makeDir(goDir);
    const tests = await fs_2.promises.readFile(root + '/chain_tests.json');
    const b = new Buf();
    b.str('package testvectors\n\nconst ChainTests = `\n');
    b.buf(tests);
    b.str('`\n');
    b.str('\n');
    b.str('var ChainTestInputs = map[string]string{\n');
    const chains = await globFiles(root + '/chains/*.json');
    for (const file of chains) {
        const raw = await fs_2.promises.readFile(file);
        const data = JSON.parse(raw.toString('utf8'));
        const imploded = implode_1.implode(data);
        b.str(`\t"${path_1.basename(file)}" : ${toGoStringLiteral(imploded)},\n`);
    }
    b.str('}\n');
    await fs_2.promises.writeFile(goDir + '/testvectors.go', b.finish());
    return;
};
exports.main = async () => {
    let rc = 0;
    try {
        await run();
    }
    catch (e) {
        console.error(e);
        rc = 2;
    }
    process.exit(rc);
};
//# sourceMappingURL=generate.js.map