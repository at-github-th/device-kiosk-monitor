package main
import (
  "log"; "os"; "time"
  "github.com/gofiber/fiber/v2"; "github.com/gofiber/fiber/v2/middleware/cors"
)
type Device struct{ ID,Type,Status string; LastSeen int64; Battery int }
var devices = []Device{
  {ID:"POS-01", Type:"POS",   Status:"online",  LastSeen:time.Now().Unix(), Battery:98},
  {ID:"KIOSK-02", Type:"Kiosk", Status:"offline", LastSeen:time.Now().Unix()-120, Battery:77},
  {ID:"LOCKER-03", Type:"Locker", Status:"online", LastSeen:time.Now().Unix(), Battery:65},
}
func auth(c *fiber.Ctx) bool { return c.Get("X-Api-Key") == os.Getenv("API_TOKEN") }
func main(){
  port := os.Getenv("API_PORT"); if port==""{ port="5103" }
  app:=fiber.New(); app.Use(cors.New(cors.Config{AllowOrigins:"*",AllowMethods:"GET,POST,OPTIONS",AllowHeaders:"Content-Type,Authorization,X-Api-Key"}))
  app.Get("/", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"name":"Real-time Device/Kiosk Monitor API","ok":true}) })
  app.Get("/api/devices", func(c *fiber.Ctx) error {
    if !auth(c){ return c.Status(401).JSON(fiber.Map{"error":"unauthorized"}) }
    now:=time.Now().Unix(); for i:=range devices{ devices[i].LastSeen=now }
    return c.JSON(devices)
  })
  app.Post("/api/devices/:id/status", func(c *fiber.Ctx) error {
    if !auth(c){ return c.Status(401).JSON(fiber.Map{"error":"unauthorized"}) }
    id:=c.Params("id"); var req struct{ Status string `json:"status"` }
    if err:=c.BodyParser(&req); err!=nil{ return c.Status(400).JSON(fiber.Map{"error":"bad_request"}) }
    for i:=range devices{ if devices[i].ID==id{ devices[i].Status=req.Status; devices[i].LastSeen=time.Now().Unix(); return c.JSON(fiber.Map{"id":id,"status":devices[i].Status,"updated":true}) } }
    return c.Status(404).JSON(fiber.Map{"error":"not_found"})
  })
  log.Printf("Listening on http://127.0.0.1:%s", port); app.Listen("127.0.0.1:"+port)
}
