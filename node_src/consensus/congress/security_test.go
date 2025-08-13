// Copyright 2024 The go-ethereum Authors
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

package congress

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// TestMinimumValidatorThreshold tests that the consensus engine enforces
// the minimum validator threshold to prevent single validator control
func TestMinimumValidatorThreshold(t *testing.T) {
	// Test cases for validator set validation
	testCases := []struct {
		name        string
		validators  []common.Address
		shouldError bool
		errorMsg    string
	}{
		{
			name:        "Valid validator set with minimum threshold",
			validators:  generateValidators(4),
			shouldError: false,
		},
		{
			name:        "Valid validator set above minimum threshold",
			validators:  generateValidators(10),
			shouldError: false,
		},
		{
			name:        "Invalid validator set below minimum threshold",
			validators:  generateValidators(3),
			shouldError: true,
			errorMsg:    "insufficient validators: minimum threshold not met: got 3, need at least 4",
		},
		{
			name:        "Invalid validator set with single validator",
			validators:  generateValidators(1),
			shouldError: true,
			errorMsg:    "insufficient validators: minimum threshold not met: got 1, need at least 4",
		},
		{
			name:        "Invalid validator set with zero validators",
			validators:  []common.Address{},
			shouldError: true,
			errorMsg:    "insufficient validators: minimum threshold not met: got 0, need at least 4",
		},
		{
			name:        "Invalid validator set above maximum threshold",
			validators:  generateValidators(22),
			shouldError: true,
			errorMsg:    "Invalid validators length: got 22, max allowed 21",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Congress instance
			config := &params.ChainConfig{
				Congress: &params.CongressConfig{
					Epoch:  30000,
					Period: 3,
				},
			}

			congress := New(config, nil)

			// Test the validateValidatorSet function
			err := congress.validateValidatorSet(tc.validators)

			if tc.shouldError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tc.errorMsg != "" && err.Error() != tc.errorMsg {
					t.Errorf("Expected error message '%s' but got '%s'", tc.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

// TestValidatorSetSecurityInFunctions tests that all critical functions
// enforce the minimum validator threshold
func TestValidatorSetSecurityInFunctions(t *testing.T) {
	config := &params.ChainConfig{
		Congress: &params.CongressConfig{
			Epoch:  30000,
			Period: 3,
		},
	}

	congress := New(config, nil)

	// Test with insufficient validators
	insufficientValidators := generateValidators(2)

	// Test updateValidators function
	err := congress.updateValidators(insufficientValidators, nil, nil, nil)
	if err == nil {
		t.Error("Expected updateValidators to fail with insufficient validators")
	}

	// Test getTopValidators would fail validation (though we can't easily test the full flow)
	// The validation is now enforced in the function
}

// TestMinimumValidatorConstants tests that the constants are properly defined
func TestMinimumValidatorConstants(t *testing.T) {
	if minValidators != 4 {
		t.Errorf("Expected minValidators to be 4, got %d", minValidators)
	}

	if maxValidators != 21 {
		t.Errorf("Expected maxValidators to be 21, got %d", maxValidators)
	}

	if minValidators >= maxValidators {
		t.Error("minValidators should be less than maxValidators")
	}
}

// Helper function to generate test validators
func generateValidators(count int) []common.Address {
	validators := make([]common.Address, count)
	for i := 0; i < count; i++ {
		// Generate deterministic addresses for testing
		addr := common.BytesToAddress([]byte{byte(i + 1), byte(i + 2), byte(i + 3)})
		validators[i] = addr
	}
	return validators
}

// TestSecurityErrorMessages tests that security-related error messages are properly defined
func TestSecurityErrorMessages(t *testing.T) {
	// Test that the error contains the expected base message
	if !strings.Contains(errInsufficientValidators.Error(), "insufficient validators: minimum threshold not met") {
		t.Errorf("Error message should contain 'insufficient validators: minimum threshold not met', got: %s", errInsufficientValidators.Error())
	}

	if !strings.Contains(errInvalidValidatorsLength.Error(), "Invalid validators length") {
		t.Errorf("Error message should contain 'Invalid validators length', got: %s", errInvalidValidatorsLength.Error())
	}
}
