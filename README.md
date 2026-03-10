# stackctl

A CLI tool for managing development stacks using Docker Compose. Stackctl provides a simple interface to start, stop, and monitor your development environment.

## Features

- Quick stack management with simple commands
- Service status monitoring
- System health checks
- Log viewing
- Service restart capabilities
- Environment validation

## Prerequisites

- Go 1.25.0 or later
- Docker and Docker Compose installed
- docker-compose.yml file in your project directory

## Installation

### From Source

```bash
git clone <repository-url>
cd stackctl
go build -o stackctl
```

### Install Globally

```bash
go install
```

## Usage

### Basic Commands

```bash
# Start the development stack
stackctl up

# Stop the development stack
stackctl down

# Stop with orphan cleanup
stackctl down --clean

# Check service status
stackctl status

# View logs
stackctl logs

# Restart services
stackctl restart
```

### System Health

```bash
# Run system diagnostics
stackctl doctor
```

The doctor command checks for:
- Go installation
- Docker installation
- Docker daemon status
- Docker Compose availability
- .env file presence

### Advanced Commands

```bash
# Deploy configuration
stackctl deploy

# Execute commands in services
stackctl exec <service> <command>

# Check configuration
stackctl config

# Show version
stackctl version
```

## Configuration

Stackctl expects a `docker-compose.yml` file in the current directory. Here's an example configuration for testing:

```yaml
version: "3"

services:
  api:
    image: nginx
    ports:
      - "8080:80"
```

**Note:** Replace this configuration with your actual service definitions for production use.

### Environment Variables

Create a `.env` file in your project directory for environment-specific configuration. The `stackctl doctor` command will verify its presence.

## Development

### Project Structure

```
stackctl/
├── cmd/                 # CLI command definitions
│   ├── root.go         # Root command
│   ├── up.go           # Start services
│   ├── down.go         # Stop services
│   ├── status.go       # Service status
│   ├── doctor.go       # System checks
│   └── ...
├── internal/            # Internal packages
│   ├── docker/         # Docker operations
│   ├── doctor/         # Health checks
│   └── preflight/      # Pre-flight checks
├── docker-compose.yml  # Example configuration
└── main.go            # Application entry point
```

### Building

```bash
# Build for current platform
go build -o stackctl

# Build with version info
go build -ldflags "-X main.version=1.0.0" -o stackctl

# Cross-compile
GOOS=linux GOARCH=amd64 go build -o stackctl-linux
GOOS=windows GOARCH=amd64 go build -o stackctl.exe
```

## Commands Reference

| Command | Description | Flags |
|---------|-------------|-------|
| `up` | Start the development stack | - |
| `down` | Stop the development stack | `--clean` Remove orphan containers |
| `status` | Show service status | - |
| `logs` | View service logs | - |
| `restart` | Restart services | - |
| `doctor` | Run system health checks | - |
| `exec` | Execute command in service | service, command |
| `deploy` | Deploy configuration | - |
| `config` | Show configuration | - |
| `version` | Show version information | - |

## Troubleshooting

### Common Issues

1. **"docker-compose.yml not found"**
   - Ensure you have a docker-compose.yml file in the current directory
   - Run `stackctl doctor` to verify your setup

2. **"Docker daemon not running"**
   - Start Docker Desktop or Docker daemon
   - Run `stackctl doctor` to check daemon status

3. **Permission denied errors**
   - Ensure your user has Docker permissions
   - On Linux/macOS, you may need to add your user to the docker group

### Debug Mode

For detailed error information, run commands with verbose output:

```bash
stackctl up --verbose
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions:

1. Run `stackctl doctor` to check your system setup
2. Check the troubleshooting section above
3. Open an issue on the repository

---

**stackctl** - Simplify your development stack management.
