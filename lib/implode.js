"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const purepack_1 = require("purepack");
const readPipe = (inp) => {
    return new Promise((resolve, reject) => {
        const chunks = [];
        inp.resume();
        inp.on('data', (data) => chunks.push(data));
        inp.on('error', err => reject(err));
        inp.on('end', () => {
            const all = Buffer.concat(chunks);
            resolve(all);
        });
    });
};
const readJSONFromPipe = async (inp) => {
    const buf = await readPipe(inp);
    return JSON.parse(buf.toString('utf8'));
};
const unpackSig = (s) => {
    const buf = Buffer.from(s, 'base64');
    const sigInfo = purepack_1.unpack(buf);
    const payload = sigInfo.body.payload;
    const sig = sigInfo.body.sig;
    return { payload, sig };
};
const encode = (o) => purepack_1.pack(o).toString('base64');
// Given a signature that we read out of the database that's a NaCl sig v2 of an outer
// v2 wrapper, strip out all of the inferrable parts and just return the parts the client
// striclty needs to verify. We'll be pruning: kid and hash from the siginfo, and
// prev, curr and seqno from the outer link v2.
const implodeSig2 = (s) => {
    const { payload, sig } = unpackSig(s);
    const outerLink = purepack_1.unpack(payload);
    outerLink[1] = 0; // numbers > 128 take 3 bytes to encode, so save some bytes here, too
    outerLink[2] = null;
    outerLink[3] = null;
    const retObj = [sig, outerLink, outerLink.length];
    return encode(retObj);
};
const implodeSig1 = (s) => {
    const { sig } = unpackSig(s);
    return sig.toString('base64');
};
const isPGPSig = (s) => {
    return !!s.match(/^-{5}BEGIN PGP MESSAGE-{5}/);
};
const implodeLinkPGP = (link) => {
    delete link.payload_json;
    return link;
};
const implodeLinkV1 = (link) => {
    const { sig } = link;
    delete link.sig;
    link.si1 = implodeSig1(sig);
    return link;
};
const implodeLinkV2 = (link) => {
    const { sig } = link;
    delete link.sig;
    link.si2 = implodeSig2(sig);
    return link;
};
const implodeLink = (link) => {
    if (link.s2) {
        return link;
    }
    const { seqno, kid, sig, payload_json, merkle_seqno } = link;
    let { sig_version } = link;
    if (!sig_version) {
        sig_version = 1;
    }
    const intermediate = { seqno, kid, sig, payload_json, sig_version, merkle_seqno };
    const pgp = isPGPSig(sig);
    if (pgp) {
        return implodeLinkPGP(intermediate);
    }
    delete intermediate.kid;
    switch (sig_version) {
        case 1:
            return implodeLinkV1(intermediate);
        case 2:
            return implodeLinkV2(intermediate);
        default:
            throw new Error('unknown link version');
    }
};
exports.implode = (f) => {
    f.chain = f.chain.map(implodeLink);
    return f;
};
const run = async () => {
    const inpRaw = await readJSONFromPipe(process.stdin);
    const inp = inpRaw;
    return exports.implode(inp);
};
exports.main = async () => {
    let rc = 0;
    try {
        const out = await run();
        console.log(JSON.stringify(out));
    }
    catch (e) {
        console.error(e);
        rc = 2;
    }
    process.exit(rc);
    return;
};
//# sourceMappingURL=implode.js.map