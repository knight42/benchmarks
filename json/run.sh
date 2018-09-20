#!/bin/bash

export NODE_ENV=production
export PYTHONOPTIMIZE=2

pushd ./bench-rs
cargo build --release
rustc_version=$(rustc -V)
popd
cp ./bench-rs/target/release/bench-rs ./bench.rs

cat > './README.md' <<EOF
# Hardware
* Processor: 3.1 GHz Intel Core i5
* Memory: 16 GB 2133 MHz LPDDR3

# Python
\`\`\`
$(python ./bench.py ./data.json)
\`\`\`

# Node.js
\`\`\`
$(node ./bench.js ./data.json)
\`\`\`

# Go
\`\`\`
$(go run ./bench.go)
\`\`\`

# Rust ($rustc_version)
\`\`\`
$(./bench.rs ./data.json)
\`\`\`
EOF
