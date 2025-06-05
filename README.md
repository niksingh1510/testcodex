# MDC Console Prototype

This repository contains a simple prototype for the **MDC Console** backend. The
backend is written in Go and exposes placeholder API endpoints for user
registration, tenant management, license activation, and provisioning of DataOS
instances and dataplanes. It stores data in memory for now and does not yet
integrate with an OIDC provider, but the structure is ready for integration with
[goauthentik/authentik](https://github.com/goauthentik/authentik).

```
POST /signup       - register a user
POST /license      - activate a license key
POST /tenants      - create a tenant ("MDC Account")
GET  /tenants      - list tenants
POST /instances    - create a DataOS instance
GET  /instances    - list created instances
POST /dataplanes   - create a dataplane
GET  /dataplanes   - list created dataplanes
```

Run the server:

```bash
cd backend
go run main.go
```

The `frontend` directory is currently empty and will contain the React/Ant
Design client in a separate repository.
