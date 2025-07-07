<h1 align="center">HTTP Client TUI</h1>

![Captura de pantalla](./public/screenshot.png) 
[![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)

Un cliente HTTP moderno con **Interfaz de Usuario en Terminal (TUI)** construido con Go. Realiza peticiones HTTP de forma visual e intuitiva desde la terminal.

## ✨ Características

- 🔧 **Métodos HTTP**: GET, POST, PUT, DELETE
- 🎨 **Interfaz Visual**: Navegación intuitiva con colores
- 📝 **Gestión Completa**: Body, Headers, Query Params, Path Params
- 🌈 **Códigos de Estado Coloreados**: Verde (2xx), Amarillo (3xx), Rojo (4xx/5xx)
- 📄 **Resaltado JSON**: Respuestas formateadas automáticamente
- ⌨️ **Atajos de Teclado**: Control completo sin mouse
- 🖥️ **Modo Pantalla Completa**: Maximizar vista de respuesta

## 🚀 Instalación

```bash
# Clonar y ejecutar
git clone https://github.com/Lauta-dev/go-http_client-tui.git
cd go-http_client-tui
go run .

# O compilar
go build -o http-client .
./http-client
```

## 🎮 Controles

| Tecla | Acción |
|-------|--------|
| `F1` | Ayuda |
| `Alt+F` | Pantalla completa |
| `Alt+B` | Vista Body |
| `Alt+Q` | Vista Query Params |
| `Alt+P` | Vista Path Params |
| `Alt+H` | Vista Headers |
| `Alt+J` | Vista Variables |
| `Alt+D` | Enviar petición |
| `Alt+N` | Copiar respuesta al portapapeles (X11 y Wayland) |
| `Ctrl+C` | Salir |

## 📋 Uso Rápido

### Petición GET
1. Seleccionar `GET`
2. URL: `https://jsonplaceholder.typicode.com/posts/1`
3. `Alt+D` para enviar

### Petición POST
1. Seleccionar `POST`
2. URL: `https://jsonplaceholder.typicode.com/posts`
3. Body (`Alt+B`):
   ```json
   {
     "title": "test",
     "body": "content",
     "userId": 1
   }
   ```
4. Headers (`Alt+H`):
   ```
   Content-Type: application/json
   ```
5. `Alt+D` para enviar

### Formato de Headers
```
Content-Type: application/json
Authorization: Bearer token-here
x-api-key: your-key
# Esta línea será ignorada
```

## 🛠️ Tecnologías

- **[tview](https://github.com/rivo/tview)** - Framework TUI
- **[tcell](https://github.com/gdamore/tcell)** - Terminal handling
- **net/http** - Cliente HTTP nativo

## 🤝 Contribución

1. Fork el repositorio
2. Crea tu rama (`git checkout -b feature/nueva-caracteristica`)
3. Commit (`git commit -am 'Añade nueva característica'`)
4. Push (`git push origin feature/nueva-caracteristica`)
5. Abre un Pull Request

---

⭐ **¡Dale una estrella si te gusta el proyecto!**
