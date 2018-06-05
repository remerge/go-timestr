package timestr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeStr(t *testing.T) {
	Stop()
	ts := time.Date(1982, 4, 3, 12, 00, 5, 1234, time.UTC)
	day := time.Date(1982, 4, 3, 00, 00, 0, 0, time.UTC)
	SetTimeStr(ts)
	require.Equal(t, Now(), ts)
	require.Equal(t, Today(), day)
	require.Equal(t, ISO8601(), "1982-04-03T12:00:05Z")
	require.Equal(t, URLSafe(), "1982-04-03T12-00-05Z")
}
