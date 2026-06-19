#!/bin/bash
set -e

SEED_NODE_ID="1efe4ede5860cd60a36d0161df60fc3e31c2a038"
SEED_IP="178.63.164.6"
CHAIN_ID="nebula-1"
DENOM="unebula"
BINARY="nebulad"

print_step() { echo -e "\n\e[1;34m>>> $1\e[0m"; }

print_step "Instalowanie Nebula binary..."
if ! command -v $BINARY &>/dev/null; then
    wget -q https://github.com/username/nebula/releases/latest/download/nebulad -O /usr/local/bin/nebulad
    chmod +x /usr/local/bin/nebulad
fi

print_step "Inicjalizacja node'a..."
$BINARY init "$(hostname)" --chain-id $CHAIN_ID

print_step "Konfiguracja persistent_peers..."
PEERS="${SEED_NODE_ID}@${SEED_IP}:26656"
sed -i "s/^persistent_peers = .*/persistent_peers = \"$PEERS\"/" ~/.nebula/config/config.toml

print_step "Konfiguracja seed_peers..."
sed -i "s/^seeds = .*/seeds = \"$PEERS\"/" ~/.nebula/config/config.toml

print_step "Konfiguracja app.toml (min gas price)..."
sed -i 's/^minimum-gas-prices = .*/minimum-gas-prices = "0.025unebula"/' ~/.nebula/config/app.toml

print_step "Systemd service..."
cat > /etc/systemd/system/nebulad.service <<EOF
[Unit]
Description=Nebula Node
After=network-online.target

[Service]
User=root
ExecStart=/usr/local/bin/nebulad start
Restart=on-failure
RestartSec=5
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable nebulad
systemctl start nebulad

print_step "Gotowe! Node synchronizuje sie z siecia Nebula."
echo "Sprawdz: journalctl -u nebulad -f"
echo "Status:  nebulad status"
