package refreshrepo

import (
	"context"
	"fmt"
)

// Delete removes a refresh session for the provided composite key.
func (r *RefreshRepository) Delete(ctx context.Context, userID, app, device string) error {
	key := makeKey(userID, app, device)

	_, err := r.rdb.Del(ctx, key).Result()
	if err != nil {
		if isUnavailable(err) {
			return ErrUnavailable
		}
		return fmt.Errorf("%w: delete refresh session", ErrInternal)
	}

	return nil
}
