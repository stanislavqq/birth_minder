package bevent

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"testing"
	"time"
)

func TestRepository_GetListWithRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db, _, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}

	database := sqlx.NewDb(db, "sqlmock")
	rep := NewRepository(database, zerolog.Logger{})

	t.Run("Get row tomorrow", func(t *testing.T) {
		const day = time.Hour * 24
		rule := NewTimeRule(day)
		queryCond, params := rep.buildQueryRule(rule)

		tomorrow := time.Now().Add(day)

		assert.Eq(t, queryCond, "day = ? AND month = ?")
		assert.Eq(t, params, []string{tomorrow.Format("2"), tomorrow.Format("1")})
	})

	t.Run("Get row next week", func(t *testing.T) {
		const week = time.Hour * 24 * 7
		rule := NewTimeRule(week)
		_, params := rep.buildQueryRule(rule)

		newWeek := time.Now().Add(week)
		assert.Eq(t, params, []string{newWeek.Format("2"), newWeek.Format("1")})
	})

}
