#!/usr/bin/env bash

set -e

docker build -f orch/prod/Dockerfile.server \
	-t portfolio/server:latest .

docker compose -f orch/prod/docker-compose.yml up -d
