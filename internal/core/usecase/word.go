package usecase

import (
	"context"

	"github.com/joqd/slovo/internal/core/domain"
	"github.com/joqd/slovo/internal/core/port"
)

type wordUsecase struct {
	persistent port.WordPersistent
	cache      port.WordCache
	xlog       port.Logger
}

func NewWordUsecase(persistent port.WordPersistent, cache port.WordCache, xlog port.Logger) port.WordUsecase {
	return &wordUsecase{
		persistent: persistent,
		cache:      cache,
		xlog:       xlog,
	}
}

func (w *wordUsecase) GetByID(ctx context.Context, id string) (*domain.Word, error) {
	word, err := w.cache.Get(ctx, id)
	if err != nil {
		word, err = w.persistent.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}

		if err := w.cache.Set(ctx, word); err != nil {
			w.xlog.Warn("cache set failed, word_id=%s, err=%v", word.ID, err)
		}
	}

	if word.Disable {
		return nil, domain.ErrDataNotFound
	}

	return word, nil
}

func (w *wordUsecase) Create(ctx context.Context, word *domain.Word) (*domain.Word, error) {
	oid, err := w.persistent.Create(ctx, word)
	if err != nil {
		return nil, err
	}

	word.ID = oid
	if err := w.cache.Set(ctx, word); err != nil {
		w.xlog.Warn("cache set failed, word_id=%s, err=%v", word.ID, err)
	}

	return word, nil
}

func (w *wordUsecase) GetByBare(ctx context.Context, bare string) (*domain.Word, error) {
	word, err := w.cache.Get(ctx, bare)
	if err != nil {
		word, err = w.persistent.GetByBare(ctx, bare)
		if err != nil {
			return nil, err
		}

		if err := w.cache.Set(ctx, word); err != nil {
			w.xlog.Warn("cache set failed, word_id=%s, err=%v", word.ID, err)
		}
	}

	if word.Disable {
		return nil, domain.ErrDataNotFound
	}

	return word, nil
}

func (w *wordUsecase) DeleteByID(ctx context.Context, id string) error {
	err := w.persistent.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	if err := w.cache.Del(ctx, id); err != nil {
		w.xlog.Warn("cache delete failed, id=%s, err=%v", id, err)
	}

	return nil
}

func (w *wordUsecase) DeleteByBare(ctx context.Context, bare string) error {
	err := w.persistent.DeleteByBare(ctx, bare)
	if err != nil {
		return err
	}

	if err := w.cache.Del(ctx, bare); err != nil {
		w.xlog.Warn("cache delete failed, bare=%s, err=%v", bare, err)
	}

	return nil
}

func (w *wordUsecase) Update(ctx context.Context, word *domain.Word) error {
	if err := w.persistent.Update(ctx, word); err != nil {
		return err
	}

	if err := w.cache.Set(ctx, word); err != nil {
		w.xlog.Warn("cache set failed, word_id=%s, err=%v", word.ID, err)
	}

	return nil
}
