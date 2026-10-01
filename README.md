# 🎵 PlayGo Base - Bot de Música Simplificado

> **Versión base y simplificada** del bot de música PlayGo. Un reproductor de música de terminal que busca y reproduce canciones desde YouTube usando `ffplay`.

Un bot de música minimalista construido en **Go** que permite reproducir música desde la terminal mediante comandos simples. Esta es la versión base que sirve como punto de partida o demostración de la funcionalidad central.

---

## 🚀 Características

- 🎼 **Reproducción desde YouTube**: Busca y reproduce música en tiempo real usando `goutubedl`
- 🎛️ **Control por Comandos**: Interfaz simple basada en comandos de texto
- 🔊 **Reproducción Local**: Usa `ffplay` para reproducir audio directamente en tu terminal
- 📦 **Sin Dependencias Complejas**: Proyecto monolítico, todo en un solo archivo
- ⚡ **Ligero y Rápido**: Compilado como binario nativo de Go

---

## 📂 Estructura del Proyecto

```text
play-go-base/
├── main.go                    # Punto de entrada y lógica principal del bot
├── funcional_variante.go      # Código de prueba/variante experimental
├── go.mod                     # Dependencias de Go
├── go.sum                     # Checksum de dependencias
```

---

## 🛠️ Instalación

### Prerrequisitos

- **Go 1.21+** ([Descargar](https://golang.org/dl/))
- **FFmpeg/ffplay** instalado en el sistema

#### Instalar FFmpeg

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install ffmpeg
```

**macOS:**
```bash
brew install ffmpeg
```

**Windows:**
Descargar desde [ffmpeg.org](https://ffmpeg.org/download.html) y agregar al PATH.

### Pasos de Instalación

```bash
# 1. Clonar el repositorio
git clone https://github.com/MartinCiro/play-go-base.git
cd play-go-base

# 2. Instalar dependencias
go mod download

# 3. Compilar el binario
go build -o playgo main.go

# 4. Ejecutar
./playgo
```

---

## 🎮 Uso

Al ejecutar el programa, verás el menú de comandos disponibles:

```bash
$ ./playgo

🎵 Bot de Música Simplificado
==============================
Comandos disponibles:
!play [canción] - Añadir canción a la cola
!revoke - Revocar tu última canción
!skip - Saltar canción actual
!queue - Mostrar cola actual
!exit - Salir del programa
```

### Comandos Disponibles

| Comando | Descripción | Ejemplo |
|---------|-------------|---------|
| `!play [canción]` | Busca y añade una canción a la cola | `!play despacito` |
| `!skip` | Salta a la siguiente canción en la cola | `!skip` |
| `!queue` | Muestra la cola de reproducción actual | `!queue` |
| `!revoke` | Elimina la última canción añadida | `!revoke` |
| `!exit` | Cierra el programa | `!exit` |

### Ejemplo de Uso

```bash
> !play bad bunny
🎵 Buscando: bad bunny...
✅ Añadido: Bad Bunny - Tití Me Preguntó
🔊 Reproduciendo...

> !queue
📋 Cola de reproducción:
  1. Bad Bunny - Tití Me Preguntó (reproduciendo)

> !play shakira
🎵 Buscando: shakira...
✅ Añadido: Shakira - Waka Waka
🔊 En cola...

> !skip
⏭️ Saltando canción actual
🔊 Reproduciendo: Shakira - Waka Waka
```

---

## 🔧 Dependencias

El proyecto utiliza las siguientes librerías de Go:

- **[goutubedl](https://github.com/wader/goutubedl)**: Wrapper de Go para yt-dlp, usado para buscar y descargar streams de YouTube
- **[imcache](https://github.com/erni27/imcache)**: Caché en memoria para optimizar búsquedas repetidas

---

## 📝 Notas de Desarrollo

### Archivo `funcional_variante.go`

Este archivo contiene código experimental o variantes de la funcionalidad principal. Puede incluir:
- Pruebas de conceptos
- Implementaciones alternativas
- Código de debugging

Para ejecutar esta variante:
```bash
go run funcional_variante.go
```

### Archivo `b.txt`

Contiene notas de desarrollo, documentación técnica o referencias utilizadas durante la creación del proyecto.

---

## ⚠️ Limitaciones

Esta versión base tiene algunas limitaciones comparada con la versión completa:

- ❌ Sin integración con TikTok Live
- ❌ Sin arquitectura hexagonal (todo en un solo archivo)
- ❌ Sin sistema de logging avanzado
- ❌ Sin instalación automática de FFmpeg
- ❌ Sin soporte para múltiples usuarios

Para una versión más completa con estas características, consulta el repositorio principal: [play-go](https://github.com/MartinCiro/play-go)

---

## 🐛 Troubleshooting

### Error: "ffplay no encontrado"

**Solución**: Instala FFmpeg en tu sistema (ver sección de Prerrequisitos).

### Error: "go.mod file not found"

**Solución**: Ejecuta `go mod init play-go-base` y luego `go mod tidy`.

### La música no se reproduce

**Solución**: 
1. Verifica que `ffplay` esté instalado: `ffplay -version`
2. Asegúrate de tener conexión a internet
3. Verifica que el volumen del sistema no esté muteado

---

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Si deseas mejorar este proyecto:

1. Fork el repositorio
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

---

## 📄 Licencia

Este proyecto es de código abierto y está disponible bajo la licencia MIT.

---

## 👤 Autor

**Martin Ciro**  
[![GitHub](https://img.shields.io/badge/GitHub-MartinCiro-181717?style=flat&logo=github)](https://github.com/MartinCiro)

---

## 🔗 Enlaces Relacionados

- **Versión completa con arquitectura hexagonal**: [play-go](https://github.com/MartinCiro/play-go)
- **Otros proyectos**: [GitHub Profile](https://github.com/MartinCiro)

---

*Desarrollado con ❤️ en Go - Versión base simplificada*