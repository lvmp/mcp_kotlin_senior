$REPO = "lvmp7/mcp-kotlin-senior"

Write-Host "🔄 Tagueando a imagem antiga (TypeScript) como 1.0.0..." -ForegroundColor Yellow
docker tag "${REPO}:latest" "${REPO}:1.0.0"
docker push "${REPO}:1.0.0"

Write-Host "🚀 Iniciando build da nova imagem (Go) v2.0.0..." -ForegroundColor Cyan
docker build -t "${REPO}:2.0.0" -t "${REPO}:latest" .

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Erro no build da imagem!" -ForegroundColor Red
    exit $LASTEXITCODE
}

Write-Host "✅ Build v2.0.0 concluído com sucesso!" -ForegroundColor Green

# Push the images
Write-Host "📤 Enviando imagens v2.0.0 e latest para o Docker Hub..." -ForegroundColor Cyan
docker push "${REPO}:2.0.0"
docker push "${REPO}:latest"

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Erro ao enviar a imagem! Verifique se você está logado (docker login)." -ForegroundColor Red
    exit $LASTEXITCODE
}

Write-Host "🎉 Imagens enviadas com sucesso para ${REPO}" -ForegroundColor Green

# Show final overview
Write-Host "`n📊 Resumo das imagens no registro:" -ForegroundColor Yellow
docker images --format "{{.Repository}}:{{.Tag}} -> {{.Size}}" | Select-String ${REPO}


