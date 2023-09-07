package os

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFs(t *testing.T) {
	ret1, err := FS.Open("/tmp/1")
	require.NoError(t, err)

	data, err := ioutil.ReadAll(ret1)
	require.NoError(t, err)

	t.Logf("%v", len(data))

	_, err = FS.Stat("/tmp/1")
	require.NoError(t, err)

	ret2, err := FS.FindProcess(1)
	require.NoError(t, err)
	ret2.Release()

	p, err := FS.StartProcess("/bin/ls", nil, &os.ProcAttr{})
	require.NoError(t, err)
	p.Wait()
}
