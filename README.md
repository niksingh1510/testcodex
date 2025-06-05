# MDC Console Prototype

This repository contains a simple prototype for the **MDC Console** backend. The
backend is written in Go and exposes placeholder API endpoints for user
registration, license activation, and provisioning of DataOS instances and
dataplanes. It does not yet integrate with an OIDC provider, but the structure
is ready for integration with [goauthentik/authentik](https://github.com/goauthentik/authentik).

```
POST /signup       - register a user
POST /license      - activate a license key
POST /instances    - create a DataOS instance (placeholder)
GET  /instances    - list created instances
POST /dataplanes   - create a dataplane (placeholder)
GET  /dataplanes   - list created dataplanes
```

Run the server:

```bash
cd backend
go run main.go
```

The `frontend` directory is currently empty and will contain the React/Ant
Design client in a separate repository.
