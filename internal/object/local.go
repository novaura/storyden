package object

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/ftag"

	"github.com/Southclaws/storyden/internal/config"
)

type localStorer struct {
	s    fs.FS
	path string
}

func NewLocalStorer(cfg config.Config) Storer {
	var path string
	if cfg.AssetStorageLocalPath != "" {
		path = cfg.AssetStorageLocalPath
	} else {
		path = "./data"
	}

	s := os.DirFS(path)

	return &localStorer{s, path}
}

func (s *localStorer) Exists(ctx context.Context, path string) (bool, error) {
	_, err := fs.Stat(s.s, path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fault.Wrap(err, fctx.With(ctx))
	}

	return true, nil
}

func (s *localStorer) Read(ctx context.Context, path string) (io.Reader, error) {
	f, err := s.s.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.NotFound))
		}
		return nil, fault.Wrap(err, fctx.With(ctx))
	}
	return f, nil
}

func (s *localStorer) Write(ctx context.Context, path string, r io.Reader) error {
	fullpath := filepath.Join(s.path, path)

	f, err := os.OpenFile(fullpath,
		os.O_CREATE|os.O_WRONLY,
		0o666,
	)
	if err != nil {
		if os.IsNotExist(err) {
			return fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.NotFound))
		}
		return fault.Wrap(err, fctx.With(ctx))
	}

	_, err = io.Copy(f, r)
	if err != nil {
		return fault.Wrap(err, fctx.With(ctx))
	}

	return nil
}