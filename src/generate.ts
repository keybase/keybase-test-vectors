import {mkdir, access} from 'fs'
import {dirname, basename} from 'path'
import {promises as fsPromises} from 'fs'
import glob from 'glob'
import {ChainFile, implode} from './implode'

const makeDir = (d: string): Promise<void> => {
  return new Promise((resolve, reject) => {
    access(d, err => {
      if (err && err.code === 'ENOENT') {
        mkdir(d, err => {
          if (err) {
            reject(err)
          } else {
            resolve()
          }
        })
      } else if (err) {
        reject(err)
      } else {
        resolve()
      }
    })
  })
}

const globFiles = async (pattern: string): Promise<string[]> => {
  return new Promise((resolve, reject) => {
    glob(pattern, (err, matches) => {
      if (err) {
        reject(err)
      } else {
        resolve(matches.sort())
      }
    })
  })
}

class Buf {
  b: Buffer[]
  constructor() {
    this.b = []
  }
  str(s: string): void {
    this.buf(Buffer.from(s, 'utf8'))
  }
  buf(b: Buffer): void {
    this.b.push(b)
  }
  finish(): Buffer {
    return Buffer.concat(this.b)
  }
}

const toGoStringLiteral = (f: ChainFile): string => '`' + JSON.stringify(f).replace('`', '`+"`"+`') + '`'

const run = async (): Promise<void> => {
  const root = dirname(__dirname)
  const goDir = root + '/go'
  await makeDir(goDir)
  const tests = await fsPromises.readFile(root + '/chain_tests.json')
  const b = new Buf()
  b.str('package testvectors\n\nconst ChainTests = `\n')
  b.buf(tests)
  b.str('`\n')
  b.str('\n')
  b.str('var ChainTestInputs = map[string]string{\n')
  const chains = await globFiles(root + '/chains/*.json')
  for (const file of chains) {
    const raw = await fsPromises.readFile(file)
    const data = JSON.parse(raw.toString('utf8')) as ChainFile
    const imploded = implode(data)
    b.str(`\t"${basename(file)}" : ${toGoStringLiteral(imploded)},\n`)
  }
  b.str('}\n')
  await fsPromises.writeFile(goDir + '/testvectors.go', b.finish())
  return
}

export const main = async (): Promise<void> => {
  let rc = 0
  try {
    await run()
  } catch (e) {
    console.error(e)
    rc = 2
  }
  process.exit(rc)
}
