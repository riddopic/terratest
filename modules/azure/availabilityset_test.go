// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	// subscriptionID is overridden by the environment variable "ARM_SUBSCRIPTION_ID"
	subscriptionID = ""
	rgName         = "terratest-rg"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods to create and delete network resources are added, these tests can be extended.
*/

func TestGetAvailabilitySetClientE(t *testing.T) {
	t.Parallel()

	client, err := GetAvailabilitySetClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestGetAvailabilitySetE(t *testing.T) {
	t.Parallel()

	avsName := ""

	_, err := GetAvailabilitySetE(t, avsName, rgName, subscriptionID)

	require.Error(t, err)
}

func TestGetAvailabilitySetFaultDomainCountE(t *testing.T) {
	t.Parallel()

	avsName := ""

	_, err := GetAvailabilitySetFaultDomainCountE(t, avsName, rgName, subscriptionID)

	require.Error(t, err)
}

func TestGetAvailabilitySetVMsE(t *testing.T) {
	t.Parallel()

	avsName := ""

	_, err := GetAvailabilitySetVMsE(t, avsName, rgName, subscriptionID)

	require.Error(t, err)
}

func TestAvailabilitySetExistsE(t *testing.T) {
	t.Parallel()

	avsName := ""

	_, err := AvailabilitySetExistsE(t, avsName, rgName, subscriptionID)

	require.Error(t, err)
}