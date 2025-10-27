
package main
import ("github.com/gofiber/fiber/v2")
func main() {
    app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"name":"Real-time Device/Kiosk Monitor API","ok":true}) })
    app.Get("/api/devices", func(c *fiber.Ctx) error {
        return c.JSON([]fiber.Map{{"id":"D-1","type":"RFID","alive":true},{"id":"D-2","type":"Locker","alive":false}})
    })
    app.Listen(":5103")
}
