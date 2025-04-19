#!/bin/bash

echo "=== Iniciando Loop 1: sem API_KEY ==="
for i in {1..200}; do
  echo "Request #$i [Loop 1 - sem API_KEY]"
  curl -i http://localhost:8080/hello
  echo ""
done

echo "=== Iniciando Loop 2: com API_KEY ==="
for i in {1..200}; do
  echo "Request #$i [Loop 2 - com API_KEY]"
  curl -H "API_KEY: TOKENTEST" -i http://localhost:8080/hello
  echo ""
done
