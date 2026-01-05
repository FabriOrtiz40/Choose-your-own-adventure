# 🧭 Choose Your Own Adventure (Gophercise #3)

Este proyecto es una aplicación web escrita en Go que recrea la experiencia clásica de los libros de “Choose Your Own Adventure” (Elige tu propia aventura), donde los usuarios toman decisiones y eligen entre diferentes caminos narrativos.

---

## 🚀 ¿Cómo funciona?

- El programa carga una historia en formato JSON.
- Cada capítulo tiene un título, uno o más párrafos de texto, y una lista de opciones que apuntan a otros capítulos.
- El servidor renderiza cada capítulo como una página web.
- Los usuarios hacen clic en las opciones para navegar por diferentes rutas de la historia.

---

## 📦 Estructura del Proyecto

```
.
├── cmd/web/             # Punto de entrada principal (main.go)
│
├── story/               # Lógica del dominio (story.go)
│   ├── story.go         # Tipos, parsers, handler, templates
│
├── gopher.json          # Archivo JSON con la historia
├── template.html        # Template HTML base (estilo caricatura urbana)
├── go.mod
├── go.sum
└── README.md
```

---

## 🧩 Conceptos Aplicados en Go

### ✅ Decodificación de JSON
Usamos `encoding/json` para convertir el archivo `gopher.json` a estructuras Go como:

```go
type Story map[string]Chapter
type Chapter struct {
  Title      string
  Paragraphs []string
  Options    []Option
}
```

### ✅ Templates en HTML
Usamos `html/template` para renderizar cada capítulo dinámicamente en el navegador. También aplicamos estilos personalizados.

### ✅ http.Handler personalizado
Creamos un handler que:
- Determina la ruta del capítulo (usando la URL).
- Busca el capítulo en la historia.
- Ejecuta el template con los datos de ese capítulo.

### ✅ Funciones como opciones (Functional Options)
Configuramos el handler usando opciones como `WithTemplate` y `WithPathFunc`:

```go
h := story.NewHandler(storyData,
  story.WithTemplate(myTemplate),
  story.WithPathFunc(myPathParser),
)
```

Esto permite inyectar comportamientos personalizados sin modificar el handler base.

### ✅ Custom Paths
Los paths por defecto comienzan en `/`, pero también agregamos soporte para rutas como `/story/intro` usando una función `pathFn`.

```go
func pathFn(r *http.Request) string {
  path := strings.TrimSpace(r.URL.Path)
  path = strings.TrimPrefix(path, "/story/")
  return path
}
```

### ✅ Manejo de errores
- Si un capítulo no existe: `http.StatusNotFound`.
- Si falla el renderizado del template: `http.StatusInternalServerError`.

---

## 🌐 Cómo correr el proyecto

```bash
go run cmd/web/main.go
```

También podés pasar parámetros:

```bash
go run cmd/web/main.go -file=gopher.json -port=8080
```

Abrí en el navegador: `http://localhost:8080`

---

## ✨ Mejoras y Bonus Implementados

- ✅ Estilo HTML personalizado (caricatura urbana).
- ✅ Soporte para rutas `/story/{chapter}`.
- ✅ Flags para elegir archivo JSON y puerto.
- ✅ Organización modular (`cmd/`, `story/`).
- ✅ Uso de Functional Options para inyección flexible de dependencias.

---

## 📘 Ejemplo de historia (resumen)

- Comienza en `intro`.
- Elegís entre visitar Denver o Nueva York.
- En Nueva York podés ver una conferencia extraña.
- En Denver vas a esquiar.
- Algunos caminos terminan en casa. Otros, en sorpresas.

---

## 🛠 Posibles mejoras futuras

- Guardar progreso del jugador (cookies / sesiones).
- Historial de decisiones.
- Soporte para múltiples archivos de historia.
- Versión CLI del juego.

---

## 📚 Aprendizajes Clave

- Cómo usar `http.Handler` como interfaces reutilizables.
- Inyección de dependencias mediante opciones funcionales.
- Manejo de JSON dinámico con `map[string]T`.
- Organización profesional de código en Go.

---

## 🤘 Autor

Hecho por Fabrizio Ortiz. Parte del curso [Gophercises by Jon Calhoun](https://gophercises.com/).
