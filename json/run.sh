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
10K:
\`\`\`
$(python ./bench.py ./data.json)
\`\`\`

130K:
\`\`\`
$(python ./bench.py ./large-data.json)
\`\`\`

# Node.js
10K:
\`\`\`
$(node ./bench.js ./data.json)
\`\`\`

130K:
\`\`\`
$(node ./bench.js ./large-data.json)
\`\`\`

# Go
10K:
\`\`\`
$(go run ./bench.go ./data.json)
\`\`\`

130K:
\`\`\`
$(go run ./bench.go ./large-data.json)
\`\`\`

# Rust ($rustc_version)
10K:
\`\`\`
$(./bench.rs ./data.json)
\`\`\`

130K:
\`\`\`
$(./bench.rs ./large-data.json)
\`\`\`
EOF
