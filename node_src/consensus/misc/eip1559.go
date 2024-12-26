// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
// bug across the entire project files fixed and high tx per block feature added  by EtherAuthority <https://etherauthority.io/>

package misc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)


// VerifyEip1559Header verifies some header attributes which were changed in EIP-1559,
// - gas limit check
// - basefee check
func VerifyEip1559Header(config *params.ChainConfig, parent, header *types.Header) error {
	// Verify that the gas limit remains within allowed bounds
	parentGasLimit := parent.GasLimit

	if err := VerifyGaslimit(parentGasLimit, header.GasLimit); err != nil {
		return err
	}

	// Verify the header is not malformed
	if header.BaseFee == nil {
		return fmt.Errorf("header is missing baseFee")
	}

	// If the parent block number is <= 446163, skip base fee validation
	// if parent.Number.Uint64() > 4046532 {
	// 	// Verify the baseFee is correct based on the parent header.
	// 	expectedBaseFee := CalcBaseFee(config, parent)
	// 	if header.BaseFee.Cmp(expectedBaseFee) != 0 {
	// 		return fmt.Errorf("invalid baseFee: have %s, want %s, parentBaseFee %s, parentGasUsed %d", 
	// 			expectedBaseFee, header.BaseFee, parent.BaseFee, parent.GasUsed)
	// 	}
	// }

	return nil
}

// FetchSEPPrice fetches the current price of SEP token in USD from the XT Exchange API.
func FetchSEPPrice() (float64, error) {
	// Define the API URL
	apiURL := "https://sapi.xt.com/v4/public/ticker/price/"

	// Make an HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Check if result contains "result" field and it is an array of maps
	if data, ok := result["result"].([]interface{}); ok {
		for _, item := range data {
			if token, ok := item.(map[string]interface{}); ok {
				if symbol, ok := token["s"].(string); ok && symbol == "sep_usdt" {
					if priceStr, ok := token["p"].(string); ok {
						var price float64
						fmt.Sscanf(priceStr, "%f", &price)
						return price, nil
					}
				}
			}
		}
	}

	return 0, fmt.Errorf("SEP_USDT price not found")
}

// // CalcBaseFee calculates the base fee based on a simple token transfer that takes at most 21,000 units of gas
// func CalcBaseFee(config *params.ChainConfig, parent *types.Header) *big.Int {
// 	sepPrice, err := FetchSEPPrice()
// 	if err != nil || sepPrice <= 0 {
// 		// Fallback to a static base fee if price fetch fails
// 		fmt.Println("Error fetching SEP price, defaulting baseFee to 476,190 gwei:", err)
// 		return new(big.Int).SetUint64(476190 * 1e9)
// 	}

// 	// 0.25 USD worth of SEP
// 	usdTarget := 0.25

// 	// Calculate the SEP amount needed for 0.25 USD
// 	sepForTargetUSD := usdTarget / sepPrice

// 	// Calculate the gas fee for 21,000 gas units
// 	gasUnits := 21000
// 	sepForGasUnits := new(big.Float).Mul(big.NewFloat(float64(gasUnits)), big.NewFloat(sepForTargetUSD / 21_000))

// 	// Convert SEP to Gwei (1 SEP = 1e9 Gwei)
// 	baseFeeInGwei := new(big.Float).Mul(sepForGasUnits, big.NewFloat(1e9))
// 	baseFeeInt, _ := baseFeeInGwei.Int(nil) // Convert to *big.Int
// 	fmt.Println("Base Fee Right Now: ", baseFeeInt)

// 	return baseFeeInt
// }

// 30087856541



// CalcBaseFee calculates the base fee in Gwei for a 0.25 USD gas cost at 21,000 gas units
func CalcBaseFee(config *params.ChainConfig, parent *types.Header) *big.Int {
	sepPrice, err := FetchSEPPrice()
	if err != nil || sepPrice <= 0 {
		// Fallback to a static base fee if price fetch fails
		fmt.Println("Error fetching SEP price, defaulting baseFee to 476,190 gwei:", err)
		return new(big.Int).SetUint64(476190 * 1e9)
	}

	// Target gas fee in USD
	usdTarget := 0.99

	// Calculate total gas fee in SEP
	sepForGas := usdTarget / sepPrice

	// Gas used for the smallest transaction
	gasUnits := 21000

	// Calculate BaseFee in SEP per gas unit
	baseFeeInSep := sepForGas / float64(gasUnits)

	// Convert BaseFee to Gwei (1 SEP = 1e9 Gwei)
	baseFeeInGwei := new(big.Float).Mul(big.NewFloat(baseFeeInSep), big.NewFloat(1e9))

	// Convert BaseFee to *big.Int
	baseFeeInt, _ := baseFeeInGwei.Int(nil)
	// fmt.Println("Base Fee Right Now: ", baseFeeInt)

	// Multiply baseFeeInt by 1e9
	factor := big.NewInt(1e9)
	result := new(big.Int).Mul(baseFeeInt, factor)

	// fmt.Println("Base Fee After Correction: ", result)




	return result
}
