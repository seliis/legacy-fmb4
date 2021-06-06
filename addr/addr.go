package addr

import (
	"fmb/addr/header"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func isExist(n string) bool {
	var r bool = false
	filepath.Walk("./addr", func(p string, i os.FileInfo, e error) error {
		if i.Name() == n+".miho" {
			r = true
		}
		return nil
	})
	return r
}

func Routing(app *fiber.App) {
	arrNavs := header.GetNavs()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home/home", fiber.Map{
			"Navs": arrNavs,
		}, "addr")
	})

	app.Get("/:addr", func(c *fiber.Ctx) error {
		p := c.Params("addr")
		for _, v := range arrNavs {
			if v.Menu == p && isExist(p) {
				return c.Render(p+"/"+p, fiber.Map{
					"Navs": arrNavs,
				}, "addr")
			}
		}
		return c.Render("deny/deny", fiber.Map{
			"Navs": arrNavs,
		}, "addr")
	})
}
