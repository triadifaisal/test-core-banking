package intrfc

import "context"

// AuthUseCase ...
type AuthUseCase interface {
	Auth(
		ctx context.Context,
	) string
}
