import {unpack, pack} from 'purepack'

const readPipe = (inp: NodeJS.ReadStream): Promise<Buffer> => {
  return new Promise((resolve, reject) => {
    const chunks: Buffer[] = []
    inp.resume()
    inp.on('data', (data: Buffer) => chunks.push(data))
    inp.on('error', err => reject(err))
    inp.on('end', () => {
      const all = Buffer.concat(chunks)
      resolve(all)
    })
  })
}

const readJSONFromPipe = async (inp: NodeJS.ReadStream): Promise<any> => {
  const buf = await readPipe(inp)
  return JSON.parse(buf.toString('utf8'))
}

export type ChainFile = {
  chain: Link[]
}

type Link = {
  seqno: number
  kid: string | null
  sig: string
  payload_json: string | null
  sig_version: number | null
  merkle_seqno: number | null
  si1: string | null
  si2: string | null
  s2: string | null
}

type NaclSigInfo = {
  body: {
    payload: Buffer
    sig: Buffer
  }
  hash?: {
    type: number
    value: Buffer
  }
}

type OuterLinkV2 = [
  2,
  number, // Seqno
  Buffer | null, // Prev
  Buffer | null, // Curr
  number, // linkType
  number, // seqType
  boolean // Ignore if unsupported
]

const unpackSig = (s: string): {payload: Buffer; sig: Buffer} => {
  const buf = Buffer.from(s, 'base64')
  const sigInfo = unpack(buf) as NaclSigInfo
  const payload = sigInfo.body.payload
  const sig = sigInfo.body.sig
  return {payload, sig}
}

const encode = (o: any): string => pack(o).toString('base64')

// Given a signature that we read out of the database that's a NaCl sig v2 of an outer
// v2 wrapper, strip out all of the inferrable parts and just return the parts the client
// striclty needs to verify. We'll be pruning: kid and hash from the siginfo, and
// prev, curr and seqno from the outer link v2.
const implodeSig2 = (s: string): string => {
  const {payload, sig} = unpackSig(s)
  const outerLink = unpack(payload) as OuterLinkV2
  outerLink[1] = 0 // numbers > 128 take 3 bytes to encode, so save some bytes here, too
  outerLink[2] = null
  outerLink[3] = null
  const retObj = [sig, outerLink, outerLink.length]
  return encode(retObj)
}

const implodeSig1 = (s: string): string => {
  const {sig} = unpackSig(s)
  return sig.toString('base64')
}

const isPGPSig = (s: string): boolean => {
  return !!s.match(/^-{5}BEGIN PGP MESSAGE-{5}/)
}

const implodeLinkPGP = (link: Link): Link => {
  delete link.payload_json
  return link
}

const implodeLinkV1 = (link: Link): Link => {
  const {sig} = link
  delete link.sig
  link.si1 = implodeSig1(sig)
  return link
}

const implodeLinkV2 = (link: Link): Link => {
  const {sig} = link
  delete link.sig
  link.si2 = implodeSig2(sig)
  return link
}

const implodeLink = (link: Link): Link => {
  if (link.s2) {
    return link
  }
  const {seqno, kid, sig, payload_json, merkle_seqno} = link
  let {sig_version} = link
  if (!sig_version) {
    sig_version = 1
  }
  const intermediate = {seqno, kid, sig, payload_json, sig_version, merkle_seqno} as Link
  const pgp = isPGPSig(sig)
  if (pgp) {
    return implodeLinkPGP(intermediate)
  }
  delete intermediate.kid
  switch (sig_version) {
    case 1:
      return implodeLinkV1(intermediate)
    case 2:
      return implodeLinkV2(intermediate)
    default:
      throw new Error('unknown link version')
  }
}

export const implode = (f: ChainFile): ChainFile => {
  f.chain = f.chain.map(implodeLink)
  return f
}

const run = async (): Promise<any> => {
  const inpRaw = await readJSONFromPipe(process.stdin)
  const inp = inpRaw as ChainFile
  return implode(inp)
}

export const main = async (): Promise<void> => {
  let rc = 0
  try {
    const out = await run()
    console.log(JSON.stringify(out))
  } catch (e) {
    console.error(e)
    rc = 2
  }
  process.exit(rc)
  return
}
