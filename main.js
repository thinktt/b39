const bip39 = require('bip39')
const crypto = require('crypto')
const hexCode0 = process.env.HEX_ENT0
const hexCode1 = process.env.HEX_ENT1
const hexCodeLong = process.env.HEX_ENT_LONG


const rand0 = crypto.randomBytes(16).toString('hex')
const rand1 = crypto.randomBytes(16).toString('hex')
const randLong = rand0 + rand1


// const nemo3 = getNemo(randLong)
// const nemo0 = getNemo(hexCode0)
// const nemo1 = getNemo(hexCode1)
const nemo2 = getNemo(hexCodeLong)


// console.log(nemo3, '\n')
// console.log(nemo0, '\n') 
// console.log(nemo1, '\n')
console.log(nemo2, '\n')


function getNemo(hexStr) {
  const memStr = bip39.entropyToMnemonic(hexStr)
  // const hexStrBack = bip39.mnemonicToEntropy(memStr)
  // console.log(bip39.validateMnemonic(memStr))
  // console.log(hexStr)
  // console.log(hexStrBack)
  return memStr
}


// let i = 0
// for (word of nemo0.split(' ')) {
//   console.log(word)
//   i++
//   if (i == 4) {
//     i = 0
//     console.log()
//   }
// }