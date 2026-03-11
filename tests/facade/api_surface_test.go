package facade_test

import (
	"testing"

	NoKV "github.com/feichai0017/NoKV/engine"
	"github.com/stretchr/testify/require"
)

var (
	_ NoKV.UserKV     = (*NoKV.DB)(nil)
	_ NoKV.MVCCStore  = (*NoKV.DB)(nil)
	_ NoKV.EngineMeta = (*NoKV.DB)(nil)
)

func TestPublicFacadeBasicCRUD(t *testing.T) {
	opt := NoKV.NewDefaultOptions()
	opt.WorkDir = t.TempDir()

	db := NoKV.Open(opt)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	require.NoError(t, db.Set([]byte("alpha"), []byte("one")))
	require.NoError(t, db.Set([]byte("beta"), []byte("two")))

	got, err := db.Get([]byte("alpha"))
	require.NoError(t, err)
	require.Equal(t, []byte("one"), got.Value)

	iter := db.NewIterator(nil)
	t.Cleanup(func() {
		require.NoError(t, iter.Close())
	})
	iter.Rewind()
	require.True(t, iter.Valid())
}
