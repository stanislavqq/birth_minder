package notify

import (
	"github.com/gookit/goutil/testutil/assert"
	"testing"
	"time"
)

func Test_ParseFormatMessage(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		msgFormat := "Напоминание: \n\n {fullname} {soon_time}"
		paramsMap := map[string]string{"fullname": "Тест Тестов", "soon_time": "завтра"}
		formatedMsg := parseFormatMessage(msgFormat, paramsMap)
		assert.Eq(t, "Напоминание: \n\n Тест Тестов завтра", formatedMsg)
	})

	t.Run("test 2", func(t *testing.T) {
		paramsMap := map[string]string{"firstname": "Танюхи", "soon_time": "Завтра"}
		formatedMsg := parseFormatMessage("{soon_time} будет днюха у {firstname}!", paramsMap)
		assert.Eq(t, "Завтра будет днюха у Танюхи!", formatedMsg)
	})

}

func Test_durationToStringFormat(t *testing.T) {
	var res string

	t.Run("1 день = завтра", func(t *testing.T) {
		res = durationToStringFormat(time.Hour * 24)
		assert.Eq(t, "завтра", res)
	})

	t.Run("2 дня = после завтра", func(t *testing.T) {
		res = durationToStringFormat(time.Hour * 24 * 2)
		assert.Eq(t, "после завтра", res)
	})

	t.Run("7 дней = неделя", func(t *testing.T) {
		res = durationToStringFormat(time.Hour * 24 * 7)
		assert.Eq(t, "через неделю", res)
	})

	t.Run("5 дней = через 5 дней", func(t *testing.T) {
		res = durationToStringFormat(time.Hour * 24 * 5)
		assert.Eq(t, "через 5 дней", res)
	})

	t.Run("14 дней = через 2 недели", func(t *testing.T) {
		res = durationToStringFormat(time.Hour * 24 * 14)
		assert.Eq(t, "через 2 недели", res)
	})

	t.Run("14 дней = через 2 недели", func(t *testing.T) {
		res = durationToStringFormat(time.Hour * 24 * 14)
		assert.Eq(t, "через 2 недели", res)
	})

	t.Run("56 дней = через 56 дней", func(t *testing.T) {
		res = durationToStringFormat(time.Hour * 24 * 56)
		assert.Eq(t, "через 56 дней", res)
	})
}
