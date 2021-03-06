package simplecontroller

import (
	"strings"
	"testing"

	"github.com/mattermost/mattermost-load-test-ng/loadtest/control"
	"github.com/mattermost/mattermost-load-test-ng/loadtest/user/userentity"
	"github.com/stretchr/testify/require"
)

func TestSetRate(t *testing.T) {
	c := New(1, &userentity.UserEntity{}, make(chan control.UserStatus))

	require.Equal(t, 1.0, c.rate)

	err := c.SetRate(-1.0)
	require.NotNil(t, err)
	require.Equal(t, 1.0, c.rate)

	err = c.SetRate(0.0)
	require.Nil(t, err)
	require.Equal(t, 0.0, c.rate)

	err = c.SetRate(1.5)
	require.Nil(t, err)
	require.Equal(t, 1.5, c.rate)
}

func TestGetErrOrigin(t *testing.T) {
	var origin string
	test := func() {
		origin = getErrOrigin()
	}
	test()
	require.True(t, strings.HasPrefix(origin, "simplecontroller.TestGetErrOrigin"))
}
