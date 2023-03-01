# mysrtafes-backend

## Golang API

### First Build

```bash
make first
```

### Run Tests

```bash
make
or
make test
```

### Output API Spec

```bash
make mysrtafes-api.html
```

### Directory Structure

```bash
.
├── (cover)[Output Test Coverage File]
├── (spec)[Output API Spec File]
├── .image[Docker Image File]
├── documents
│   └── open-api[API Spec]
└── src
    ├── cmd[Entry Point]
    ├── handle[Request/Response Interface Source]
    │   └── http
    ├── pkg[Domain Source]
    └── repository[Driver Source]
        └── models
```
