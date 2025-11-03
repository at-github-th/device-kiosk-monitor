# Device/Kiosk Monitor

**Stack:** Go (Fiber)  
**API:** http://127.0.0.1:5103  
**Web:** http://localhost:5503

## Run (local)

### API
cd device-kiosk-monitor-native/api && go run .

### Web (static tester)
cd device-kiosk-monitor-native/web && python3 -m http.server 5503

## Test
- **Ping:** curl -s http://127.0.0.1:5103 | jq .
- **Devices:** GET /api/devices\n- **Update:** POST /api/devices/:id/status

