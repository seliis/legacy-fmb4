package addr

import (
	"os"
	"path/filepath"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func isExist(p string) bool {
	var r bool = false
	filepath.Walk("./addr", func(
		d string, i os.FileInfo, e error) error {
		if i.Name() == p+".miho" {
			r = true
		}
		return nil
	})
	return r
}

func getMap(p string) interface{} {
	return nil
}

func Routing(app *fiber.App, rdb *redis.Client) {
	app.Get("/:addr?", func(c *fiber.Ctx) error {
		p := c.Params("addr")

		if p == "" {
			return c.Render(
				"home/home", fiber.Map{}, "addr",
			)
		} else if isExist(p) {
			return c.Render(
				p+"/"+p, fiber.Map{}, "addr",
			)
		} else {
			return c.Render(
				"denied/denied", fiber.Map{}, "addr",
			)
		}
	})
}
