#!/usr/bin/env bash

set -e

cd repo
git pull git@github.com:jaxfu/jacksonfurr.dev.git
cd ..
docker build -t portfolio_deploy:latest .
docker compose up -d
