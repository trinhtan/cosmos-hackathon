/* eslint-disable no-console */

export function generateNewMnemonic() {
  try {
    const bip39 = require('bip39');
    const mnemonic = bip39.generateMnemonic();
    return mnemonic;
  } catch (err) {
    console.log(err);
  }
}

export function getSunchainAddress(mnemonic) {
  try {
    const chainId = 'band-consumer';
    const cosmosjs = require('@cosmostation/cosmosjs');
    const sunchain = cosmosjs.network(process.env.VUE_APP_API_BACKEND, chainId);
    const address = sunchain.getAddress(mnemonic);
    return address;
  } catch (error) {
    console.log(error);
  }
}
