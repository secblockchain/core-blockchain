# Light Client Disablement - Security Fix

Light clients using the Congress consensus engine are susceptible to multiple security vulnerabilities:
- Flawed signature recency validation
- Single validator signature authority issues
- Long-range/reorg attacks
- Validator set manipulation


## Files Modified

### 1. `node_src/cmd/utils/flags.go`

#### **Functions Modified:**
- **`RegisterEthService()`**: Prevents light client initialization with clear error message
- **Flag Definitions**: All light client flags disabled and hidden from help

#### **Security Changes:**
- **Light Client Mode**: Returns fatal error when attempting to use light sync
- **Light Client Serving**: Disabled with warning message
- **All Light Client Flags**: Hidden from help and forced to safe values

#### **Flag Modifications:**
- **`--light.serve`**: Disabled (forced to 0)
- **`--light.ingress`**: Disabled (forced to 0)  
- **`--light.egress`**: Disabled (forced to 0)
- **`--light.maxpeers`**: Disabled (forced to 0)
- **`--ulc.servers`**: Disabled (forced to empty)
- **`--ulc.fraction`**: Disabled (forced to 0)
- **`--ulc.onlyannounce`**: Disabled
- **`--light.nopruning`**: Disabled
- **`--light.nosyncserve`**: Disabled
- **`--syncmode light`**: Disabled with clear error message
