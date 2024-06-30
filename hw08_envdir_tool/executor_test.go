package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	env := Environment{
		"TEST_ENV_VAR": EnvValue{
			Value: "test_value",
		},
	}
	cmd := []string{"./test.sh"}
	returnCode := RunCmd(cmd, env)
	require.Equal(t, returnCode, 0, fmt.Sprintf("Unexpected return code %d", returnCode))
}
