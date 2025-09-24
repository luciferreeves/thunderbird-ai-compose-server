# thunderbird-ai-compose-server
Server for **Thunderbird AI Compose** plugin

This server should be run before using the [Thunderbird AI Compose](https://github.com/luciferreeves/thunderbird-ai-compose) Thunderbird plugin. You can either selfhost this server on a VPS you own or you can run it locally on your machine. 

> [!IMPORTANT]
> I **do not** provide a public instance of this server for you to use nor do I provide support for selfhosting this server.

## Running the server
You can either run the server directly or download a prebuilt binary from the [releases page](https://github.com/luciferreeves/thunderbird-ai-compose-server/releases).

### Running from source
1. Make sure you have [Go](https://go.dev/dl/) installed (version 1.25 or higher).
2. Clone this repository:
    ```bash
    git clone https://github.com/luciferreeves/thunderbird-ai-compose-server.git
    cd thunderbird-ai-compose-server
    ```
3. Install dependencies:
    ```bash
    go mod download
    ```
4. Build the server:
    ```bash
    go build -o thunderbird-ai-compose-server
    ```
5. Run the server:
    Make sure to set up the environment variables as described in the [Setting up the server](#setting-up-the-server) section before running the server.

    ```bash
    ./thunderbird-ai-compose-server
    ```

### Running from a prebuilt binary
1. Download the latest release for your operating system from the [releases page](https://github.com/luciferreeves/thunderbird-ai-compose-server/releases).
2. Make sure to set up the environment variables as described in the [Setting up the server](#setting-up-the-server) section before running the server.
3. Run the binary:
    ```bash
    ./thunderbird-ai-compose-server
    ```

## Setting up the server
In order to setup the server, copy the [`.env.example`](./.env.example) file to `.env` and fill in the required environment variables. See the following table for details on environment variables:

| Variable | Description | Required | Default Value |
|----------|-------------|----------|---------------|
| `PORT` | The port the server will run on | ❌ | `3000` |
| `PROVIDER` | The AI provider to use. See [Supported Providers](#supported-providers) | ❌ | Gemini |
| `Model` | The model to use for the selected provider. See [Supported Models](#supported-models) | ❌ | `gemini-2.5-flash` |
| `API_KEY` | The API key for the selected provider. | ✅ | N/A |

### Supported Providers
| Provider | Environment Variable Value | Available |
|----------|----------------------------|-----------|
| Gemini | `gemini` | ✅ |
| OpenAI | `openai` | ❌ |

> [!NOTE]
> Models marked with a ❌ are not yet implemented, but are planned for future releases.

### Supported Models
| Provider | Supported Models |
|----------|------------------|
| Gemini | See [here](https://ai.google.dev/gemini-api/docs/models#model-variations) for available models. The default is `gemini-2.5-flash`. |
| OpenAI | Unimplemented |

## Running the Server as a Service
You can run the server as a systemd service for easier management. Create a file named `thunderbird-ai-compose-server.service` in `/etc/systemd/system/` with the following content:

```ini
[Unit]
Description=Thunderbird AI Compose Server
After=network.target
Wants=network.target
StartLimitIntervalSec=0
StartLimitBurst=3
[Service]
Type=simple
Restart=always
RestartSec=5
User=your-username
WorkingDirectory=/path/to/thunderbird-ai-compose-server
ExecStart=/path/to/thunderbird-ai-compose-server/thunderbird-ai-compose-server
EnvironmentFile=/path/to/thunderbird-ai-compose-server/.env
[Install]
WantedBy=multi-user.target
```

Make sure to replace `your-username` with your actual username and update the paths accordingly.

Then, enable and start the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable thunderbird-ai-compose-server
sudo systemctl start thunderbird-ai-compose-server
```

