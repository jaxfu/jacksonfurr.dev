#!/usr/bin/env bash

set -e

docker compose -f orch/prod/docker-compose.yml down -v
