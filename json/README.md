# Hardware
* Processor: 3.1 GHz Intel Core i5
* Memory: 16 GB 2133 MHz LPDDR3

# Python
10K:
```
json.loads: 11.282286676s
simplejson.loads: 14.774538988999998s
```

130K:
```
json.loads: 194.552027692s
simplejson.loads: 165.078162637s
```

# Node.js
10K:
```
JSON.parse: 5968.064ms
```

130K:
```
JSON.parse: 120296.183ms
```

# Go
10K:
```
json.Unmarshal: 17.705681462s
```

130K:
```
json.Unmarshal: 3m48.724326223s
```

# Rust (rustc 1.29.0 (aa3ca1994 2018-09-11))
10K:
```
serde_json::from_str: 7.156s
```

130K:
```
serde_json::from_str: 143.110s
```
