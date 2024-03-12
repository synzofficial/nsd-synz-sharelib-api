package basemodel_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	basemodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/base-model"
)

func TestSetLastPage(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		p := basemodel.PaginationWithLastPageResponse{
			PageId:        1,
			PageSize:      10,
			TotalElements: 100,
		}
		p.SetLastPage()

		assert.Equal(t, 10, p.LastPage)
	})

	t.Run("2", func(t *testing.T) {
		p := basemodel.PaginationWithLastPageResponse{
			PageId:        1,
			PageSize:      5,
			TotalElements: 2,
		}
		p.SetLastPage()

		assert.Equal(t, 1, p.LastPage)
	})

	t.Run("3", func(t *testing.T) {
		p := basemodel.PaginationWithLastPageResponse{
			PageId:        1,
			PageSize:      5,
			TotalElements: 0,
		}
		p.SetLastPage()

		assert.Equal(t, 0, p.LastPage)
	})
}
