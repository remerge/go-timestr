package timestr

import (
	"fmt"
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
	require.Equal(t, "1982-04-03T12:00:05Z", ISO8601())
	require.Equal(t, "1982-04-03T12-00-05Z", URLSafe())

	loc, err := time.LoadLocation("Europe/Berlin")
	require.NoError(t, err)
	tsDE := ts.In(loc)
	SetTimeStr(tsDE)
	_, offset := tsDE.Zone()
	require.Equal(t, fmt.Sprintf("1982-04-03T%02d:00:05+%02d:00", tsDE.Hour(), offset/3600), ISO8601())
	require.Equal(t, fmt.Sprintf("1982-04-03T%02d-00-05+%02d00", tsDE.Hour(), offset/3600), URLSafe())
	require.Equal(t, "1982-04-03T12:00:05Z", ISO8601inUTC())
	require.Equal(t, "1982-04-03T12-00-05Z", URLSafeinUTC())
}
