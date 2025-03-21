#!/bin/sh
set -e  # Faz com que o script pare em caso de erro

echo "🚀 Executando migração do banco de dados..."
/root/migrate

echo "✅ Migração concluída! Iniciando a API..."
exec /root/go-api
