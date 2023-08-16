package storage_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
	"github.com/vkolev/locmock/storage"
	"testing"
)

func TestNewOsStorageWithMemMapFs(t *testing.T) {
	t.Parallel()
	memFs := afero.Afero{Fs: afero.NewMemMapFs()}
	_ = storage.NewOsStorage(memFs)
}

func TestNewOsStorageWithRealPath(t *testing.T) {
	t.Parallel()
	realFs := afero.Afero{Fs: afero.NewOsFs()}
	_ = storage.NewOsStorage(realFs)
}

func TestStorage_GetServiceNames(t *testing.T) {
	t.Parallel()
	memFs := afero.Afero{Fs: afero.NewMemMapFs()}
	s := storage.NewOsStorage(memFs)
	// When
	memFs.Mkdir("test", 0666)
	memFs.Mkdir("service", 0666)
	memFs.MkdirAll("with/sub/dirs", 0666)

	wantLen := 3

	gotLen := len(s.GetServiceNames("/"))

	if wantLen != gotLen {
		t.Errorf("want len: %v, got len %v", wantLen, gotLen)
	}
}

func TestStorage_GetActionsForService(t *testing.T) {
	t.Parallel()
	memFs := afero.Afero{Fs: afero.NewMemMapFs()}
	s := storage.NewOsStorage(memFs)
	//When
	memFs.MkdirAll("/test/example/test", 0666)
	_, _ = memFs.Create("/test/example/test/action.yml")
	memFs.MkdirAll("/test/one/more/test", 0666)
	_, _ = memFs.Create("/test/one/more/action.yml")
	_, _ = memFs.Create("/test/one/more/test/action.yml")

	wantLen := 3
	want := []string{
		"example/test",
		"one/more",
		"one/more/test",
	}

	got := s.GetActionsForService("/", "test")
	gotLen := len(got)

	if wantLen != gotLen {
		t.Errorf("want len: %v, got len %v", wantLen, gotLen)
	}

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
