package addr

import (
	"fmb/addr/header"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func checkAddr(n string) bool {
	var found bool = false
	filepath.Walk("./addr", func(p string, i os.FileInfo, e error) error {
		if i.Name() == n+".miho" {
			found = true
		}
		return nil
	})
	return found
}

func Routing(app *fiber.App) {
	Navs := header.GetNavs()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home/home", fiber.Map{
			"Navs": Navs,
		}, "addr")
	})

	app.Get("/:addr", func(c *fiber.Ctx) error {
		p := c.Params("addr")
		for _, arr := range Navs {
			if arr.Menu == p && checkAddr(p) {
				fmt.Println("passed")
				return c.Render(p+"/"+p, fiber.Map{
					"Navs": Navs,
				}, "addr")
			}
		}
		return c.Render("nil/nil", fiber.Map{
			"Navs": Navs,
		}, "addr")
	})
}
