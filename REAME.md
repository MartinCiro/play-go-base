# 🎵 TikTok Music Bot

> **⚠️ NOTA: Este es solo un código base**  
> El repositorio completo y la versión en producción se mantienen de forma privada en:  
> **[https://github.com/MartinCiro/play-go](https://github.com/MartinCiro/play-go)**

Un bot que reproduce música controlado por el chat de TikTok Live.

## ✨ Características

- 🔴 **Conecta a TikTok Live**: Se conecta a cualquier livestream de TikTok
- 💬 **Control por chat**: Los espectadores controlan la música mediante comandos
- 🎶 **Reproducción de audio**: Reproduce música desde YouTube
- 📋 **Gestión de cola**: Sistema completo de cola de reproducción
- 👤 **Control por usuario**: Los usuarios pueden revocar sus propias canciones

## 🎮 Comandos Disponibles

### Comandos del Chat de TikTok
Los espectadores pueden usar estos comandos en el chat del livestream:

- `!play [nombre de canción]` - Añade una canción a la cola
- `!revoke` - Remueve tu última canción solicitada  
- `!skip` - Salta la canción actual (si hay suficientes votos)
- `!queue` - Muestra la cola actual de reproducción

## 🛠️ Instalación

### Prerrequisitos
- [Go](https://golang.org/dl/) 1.19 o superior
- [FFmpeg](https://ffmpeg.org/) con ffplay

### Instalar FFmpeg

**Windows:**
```bash
# Usando chocolatey
choco install ffmpeg

# O descargar desde https://ffmpeg.org/download.html
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt update && sudo apt install ffmpeg
```

**macOS:**
```bash
brew install ffmpeg
```

### Instalar y Ejecutar

```bash
# Clonar el repositorio
git clone <repository-url>
cd tiktok-music-bot

# Ejecutar el bot
go run main.go <username_tiktok>
```

## 📖 Uso

1. **Iniciar el bot:**
   ```bash
   go run main.go usernamedetiktok
   ```

2. **El bot se conectará** al livestream especificado y escuchará comandos

3. **Los espectadores** pueden usar los comandos en el chat para controlar la música

4. **El bot mostrará** en consola los comentarios y acciones realizadas

## 🎵 Funcionalidades de Música

- **Búsqueda automática**: Busca en YouTube la canción solicitada
- **Reproducción continua**: Reproduce automáticamente la siguiente canción
- **Gestión de cola**: Muestra y gestiona la cola de reproducción
- **Control de usuario**: Cada usuario puede revocar sus propias canciones

## ⚙️ Configuración

El bot funciona automáticamente con la configuración por defecto. Para personalizar el comportamiento, puedes modificar:

- Límites de canciones por usuario
- Tiempos máximos de reproducción
- Comandos personalizados
- Mensajes de respuesta

## 🚨 Notas Importantes

- ✅ El bot requiere que el livestream esté activo
- ✅ Los comandos solo funcionan durante transmisiones en vivo
- ✅ Se recomienda usar en cuentas dedicadas para streaming
- ✅ Verifica los términos de servicio de TikTok y YouTube

## 📄 Licencia

Este proyecto es de uso privado. Consulta los términos de licencia para más detalles.

---

**¿Problemas?** Revisa que FFmpeg esté instalado correctamente y que el username de TikTok sea válido.