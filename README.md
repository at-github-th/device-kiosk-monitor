# Device/Kiosk Monitor

Language: Go (Fiber)

## How to run

API
```bash
cd device-kiosk-monitor-native/api && go run .
```

Web
```bash
cd device-kiosk-monitor-native/web && python3 -m http.server 5503
```

Open http://localhost:5503

## Endpoints
- Devices: GET /api/devices\n- Update: POST /api/devices/:id/status
