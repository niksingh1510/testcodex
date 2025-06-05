# MDC Console Prototype

This repository contains a simple prototype for the **MDC Console** backend. The
backend is written in Go and exposes placeholder API endpoints for license
activation and provisioning of DataOS instances and dataplanes.

```
POST /license      - activate a license key
POST /instances    - create a DataOS instance (placeholder)
GET  /instances    - list created instances
POST /dataplanes   - create a dataplane (placeholder)
GET  /dataplanes   - list created dataplanes
```

Run the server:

```bash
cd backend
go build -mod=vendor -o mdc-console
./mdc-console
```

If your environment uses an internal proxy, set `GOPROXY` accordingly:

```bash
export GOPROXY=<internal proxy URL>
go build -o mdc-console
./mdc-console
```

The `frontend` directory is currently empty and will contain the React/Ant
Design client in a separate repository.
