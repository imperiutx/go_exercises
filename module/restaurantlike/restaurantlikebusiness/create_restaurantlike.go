package restaurantlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

// CreateRestaurantLikeStorage create
type CreateRestaurantLikeStorage interface {
	CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.RestaurantLikeCreate) error
}

type createRestaurantLike struct {
	store CreateRestaurantLikeStorage
}

// NewCreateRestaurantLikeBiz create
func NewCreateRestaurantLikeBiz(store CreateRestaurantLikeStorage) *createRestaurantLike {
	return &createRestaurantLike{store: store}
}

func (biz *createRestaurantLike) CreateNewRestaurantLike(ctx context.Context, data *restaurantlikemodel.RestaurantLikeCreate) error {
	if err := biz.store.CreateRestaurantLike(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantlikemodel.EntityName, err)
	}

	return nil
}