#!/bin/bash
set -e

CHAIN_ID="heya-1"
DENOM="uheya"
BINARY="heyad"
BINARY_PATH="/usr/local/bin/heyad"
SCRIPT_DIR="$(dirname "$(realpath "$0")")"
GENESIS_FILE="$SCRIPT_DIR/genesis.json"

R="\e[31m" G="\e[32m" Y="\e[33m" B="\e[34m" C="\e[36m" M="\e[35m" W="\e[37m" N="\e[0m"
BOLD="\e[1m"

TOTAL_STEPS=9
CURRENT_STEP=0

progress_bar() {
    local pct=$((CURRENT_STEP * 100 / TOTAL_STEPS))
    local filled=$((pct / 4))
    local empty=$((25 - filled))
    printf "  ${B}[${G}"
    printf '█%.0s' $(seq 1 $filled)
    printf "${W}░%.0s" $(seq 1 $empty)
    printf "${N}${B}] ${C}%3d%%${N}\n" "$pct"
}

print_banner() {
    clear 2>/dev/null || true
    echo ""
    echo -e "  ${BOLD}${C}╔══════════════════════════════════════╗${N}"
    echo -e "  ${BOLD}${C}║      ${Y}HEYA NETWORK — NODE INSTALL${C}      ║${N}"
    echo -e "  ${BOLD}${C}║        ${W}${CHAIN_ID}${C}               ║${N}"
    echo -e "  ${BOLD}${C}╚══════════════════════════════════════╝${N}"
    echo ""
}

print_step() {
    CURRENT_STEP=$((CURRENT_STEP + 1))
    echo ""
    echo -e "  ${BOLD}${B}┌──────────────────────────────────────────┐${N}"
    echo -e "  ${BOLD}${B}│ ${G}●${N}  Step ${C}${CURRENT_STEP}/${TOTAL_STEPS}${N}  ${W}$1${N}"
    echo -e "  ${BOLD}${B}└──────────────────────────────────────────┘${N}"
    progress_bar
}

print_ok()   { echo -e "  ${G}✓${N} $1"; }
print_warn() { echo -e "  ${Y}⚠${N} $1"; }
print_info() { echo -e "  ${C}ℹ${N} $1"; }

run_with_spinner() {
    local label="$1" ; shift
    local tmpout="/tmp/heya_spinner_out.$$"
    local tmperr="/tmp/heya_spinner_err.$$"
    >"$tmpout" >"$tmperr"
    ("$@") >"$tmpout" 2>"$tmperr" &
    local pid=$!
    local spin='⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏'
    local i=0
    while kill -0 "$pid" 2>/dev/null; do
        printf "\r  ${C}%s${N} ${W}%s${N}" "${spin:$((i++ % 10)):1}" "$label"
        sleep 0.1
    done
    wait "$pid"
    local rc=$?
    if [ $rc -eq 0 ]; then
        printf "\r  ${G}✓${N} ${W}%s${N}\n" "$label"
        cat "$tmpout"
    else
        printf "\r  ${R}✗${N} ${W}%s${N}\n" "$label"
        cat "$tmperr" >&2
    fi
    rm -f "$tmpout" "$tmperr"
    return $rc
}

print_banner

BUILD_MODE="download"
if [ "$#" -eq 0 ]; then
    echo -e "  ${BOLD}${W}Select installation method:${N}"
    echo -e "    ${G}1)${N} Download pre-built binary ${W}(default)${N}"
    echo -e "    ${Y}2)${N} Build from source"
    echo ""
    read -r -p "  Choice [1/2]: " CHOICE
    case "$CHOICE" in
        2|build) BUILD_MODE="build" ;;
        *)       BUILD_MODE="download" ;;
    esac
else
    for arg in "$@"; do
        case "$arg" in
            --build) BUILD_MODE="build" ;;
            --download) BUILD_MODE="download" ;;
        esac
    done
fi

ARCH="$(uname -m)"
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$ARCH" in
    x86_64)  ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64)   ARCH="arm64" ;;
    *)       echo -e "  ${R}✗ Unsupported architecture: $ARCH${N}"; exit 1 ;;
esac

print_step "Fetching latest release info from GitHub..."
API_URL="https://api.github.com/repos/heyanetwork/heya/releases/latest"
run_with_spinner "Fetching release info" bash -c "curl -sS '$API_URL' > /tmp/heya_api.json"
TAG=$(grep '"tag_name"' /tmp/heya_api.json | head -1 | sed 's/.*"tag_name": "\(.*\)",/\1/')
VERSION="${TAG#v}"
print_ok "Latest version: ${G}$TAG${N}"

if [ "$BUILD_MODE" = "build" ]; then
    print_step "Building Heya binary from source (${TAG})..."
    if ! command -v go &>/dev/null; then
        echo -e "  ${R}✗ Go is not installed. Install it first: https://go.dev/doc/install${N}"
        exit 1
    fi
    SRC_URL="https://github.com/heyanetwork/heya/archive/refs/tags/${TAG}.tar.gz"
    run_with_spinner "Downloading source" curl -sSL "$SRC_URL" -o /tmp/heya-src.tar.gz
    run_with_spinner "Extracting source" tar -xzf /tmp/heya-src.tar.gz -C /tmp/
    cd "/tmp/heya-${VERSION}"
    run_with_spinner "Building binary (this may take a while)" bash -c "CGO_ENABLED=1 go build -trimpath -ldflags \"-s -w -X github.com/cosmos/cosmos-sdk/version.Name=heya -X github.com/cosmos/cosmos-sdk/version.AppName=heyad -X github.com/cosmos/cosmos-sdk/version.Version=${TAG} -X github.com/cosmos/cosmos-sdk/version.Commit=$(git rev-parse HEAD 2>/dev/null || echo 'unknown') -X 'github.com/cosmos/cosmos-sdk/version.BuildTags=cosmwasm wasm'\" -o \"$BINARY_PATH\" ./cmd/heyad/"
    rm -f /tmp/heya-src.tar.gz
    rm -rf "/tmp/heya-${VERSION}"
    print_ok "Built ${BOLD}${W}$BINARY${N} ${G}${TAG}${N} from source to ${C}$BINARY_PATH${N}"
else
    print_step "Downloading latest Heya release from GitHub..."
    FILENAME="heya-${VERSION}-${OS}-${ARCH}.tar.gz"
    DOWNLOAD_URL="https://github.com/heyanetwork/heya/releases/download/${TAG}/${FILENAME}"
    run_with_spinner "Downloading ${FILENAME}" curl -sSL "$DOWNLOAD_URL" -o /tmp/heya.tar.gz
    run_with_spinner "Extracting binary" tar -xzf /tmp/heya.tar.gz -C /tmp/
    cp "/tmp/heya-${VERSION}-${OS}-${ARCH}/heyad" "$BINARY_PATH"
    chmod +x "$BINARY_PATH"
    rm -f /tmp/heya.tar.gz
    rm -rf "/tmp/heya-${VERSION}-${OS}-${ARCH}"
    print_ok "Downloaded ${BOLD}${W}$BINARY${N} ${G}v$VERSION${N} to ${C}$BINARY_PATH${N}"
fi

print_step "Checking for existing home directory..."
HEYA_HOME="$HOME/.heya"
if [ -d "$HEYA_HOME/config" ]; then
    print_warn "Existing config found, backing up keys..."
    [ -f "$HEYA_HOME/config/priv_validator_key.json" ] && cp "$HEYA_HOME/config/priv_validator_key.json" "$HEYA_HOME/config/priv_validator_key.json.bak" && print_ok "priv_validator_key.json backed up"
    [ -f "$HEYA_HOME/config/node_key.json" ] && cp "$HEYA_HOME/config/node_key.json" "$HEYA_HOME/config/node_key.json.bak" && print_ok "node_key.json backed up"
else
    print_info "No existing config found"
fi

print_step "Initializing node (generating config files)..."
"$BINARY_PATH" init "$(hostname)" --chain-id $CHAIN_ID --overwrite 2>/dev/null
print_ok "Node initialized with chain-id ${C}$CHAIN_ID${N}"

print_step "Restoring keys and genesis from repository..."
if [ -f "$HEYA_HOME/config/priv_validator_key.json.bak" ]; then
    mv "$HEYA_HOME/config/priv_validator_key.json.bak" "$HEYA_HOME/config/priv_validator_key.json"
    print_ok "priv_validator_key.json restored"
fi
if [ -f "$HEYA_HOME/config/node_key.json.bak" ]; then
    mv "$HEYA_HOME/config/node_key.json.bak" "$HEYA_HOME/config/node_key.json"
    print_ok "node_key.json restored"
fi
if [ ! -f "$GENESIS_FILE" ]; then
    echo -e "  ${R}✗ Error: $GENESIS_FILE not found!${N}"
    exit 1
fi
cp "$GENESIS_FILE" "$HEYA_HOME/config/genesis.json"
print_ok "Genesis copied from repository"

print_step "Peer and gas configuration..."
print_info "Skipping automatic peer configuration — set seeds/persistent_peers manually:"
echo -e "  ${W}  Edit: ~/.heya/config/config.toml${N}"
sed -i 's/^minimum-gas-prices = .*/minimum-gas-prices = "0.025uheya"/' ~/.heya/config/app.toml
print_ok "min gas price set to ${C}0.025uheya${N}"

print_step "Checking priv_validator_state.json..."
if [ ! -f "$HEYA_HOME/data/priv_validator_state.json" ]; then
    mkdir -p "$HEYA_HOME/data"
    echo '{"height":"0","round":0,"step":0}' > "$HEYA_HOME/data/priv_validator_state.json"
    print_ok "Created priv_validator_state.json (height=0)"
else
    print_ok "priv_validator_state.json already exists"
fi

print_step "Setting up systemd service..."
cat > /etc/systemd/system/heyad.service <<EOF
[Unit]
Description=Heya Node
After=network-online.target

[Service]
User=$(whoami)
ExecStart=$BINARY_PATH start
Restart=on-failure
RestartSec=5
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
EOF

run_with_spinner "Reloading systemd" systemctl daemon-reload
run_with_spinner "Enabling service" systemctl enable heyad
run_with_spinner "Starting node" systemctl start heyad
print_ok "Systemd service ${G}enabled${N} and ${G}started${N}"

echo ""
echo -e "  ${BOLD}${G}╔══════════════════════════════════════╗${N}"
echo -e "  ${BOLD}${G}║         ${W}INSTALLATION COMPLETE${G}         ║${N}"
echo -e "  ${BOLD}${G}╚══════════════════════════════════════╝${N}"
echo ""
echo -e "  ${C}ℹ${N} Check logs:  ${W}journalctl -u heyad -f${N}"
echo -e "  ${C}ℹ${N} Node status: ${W}$BINARY status${N}"
echo ""
echo -e "  ${B}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${N}"
echo ""
