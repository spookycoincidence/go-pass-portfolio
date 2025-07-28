# GO password generator

Generador de contraseñas en Go.  
---

## 🚀 Características

- Generación de contraseñas con:
  - Longitud configurable
  - Cantidad de dígitos y símbolos
  - Opción de desactivar mayúsculas
  - Permitir o no caracteres repetidos
- 🔁 Generación de múltiples contraseñas en un solo comando con el flag `-count`.
- Implementado en **Golang** con código limpio y modular.

---

## ⚙️ Instalación

Cloná el repositorio:

```bash
git clone https://github.com/spookycoincidence/go-password-enhanced.git
cd go-password-enhanced
go mod tidy
```

## 🛠️ Uso
Generá contraseñas ejecutando:

```bash
go run main.go [flags]
```

## 🔧 Flags disponibles:

-length	   Longitud total de la contraseña	  16
-digits	   Cantidad de dígitos numéricos	    4
-symbols	 Cantidad de símbolos	              2
-no-upper	 Desactiva el uso de mayúsculas	    false
-repeat	   Permite caracteres repetidos	      false
-count	   Cantidad de contraseñas a generar	1


📌 Ejemplo:
```bash
go run main.go -length=12 -digits=3 -symbols=1 -count=5
```

Esto generará 5 contraseñas de 12 caracteres cada una, con 3 dígitos, 1 símbolo.

## 🚀 Deploy en Railway
Podés desplegar este CLI como un servicio en Railway o cualquier entorno que soporte Go, o bien extenderlo a una API para consumir por HTTP.

🏷 Créditos
Basado en la librería oficial:
👉 sethvargo/go-password


## Desarrollado con ❤️ por spookycoincidence
