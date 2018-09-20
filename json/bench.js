#!/usr/bin/env node

'use strict'

const fs = require('fs')

const data = fs.readFileSync(process.argv[2])

const N = 1e5

console.time('JSON.parse')
for (let i = 0; i < N; i++) {
  JSON.parse(data)
}
console.timeEnd('JSON.parse')
