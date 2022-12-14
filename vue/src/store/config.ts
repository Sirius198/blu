import { blocks, env, wallet } from '../starportvuex'

import generated from './generated'
export default function init(store) {
  for (const moduleInit of Object.values(generated)) {
    moduleInit(store)
  }
  blocks(store)
  env(store)
  wallet(store)
}
