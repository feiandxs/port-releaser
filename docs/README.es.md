# Port Releaser

Una herramienta simple para liberar puertos ocupados.

## Instalación

```bash
go install github.com/codenamev/port-releaser@latest
```

## Uso

```bash
port-releaser -p 3000
```

## Parámetros

- `-p`: El número de puerto que se desea liberar

## Cómo funciona

1. Encuentra el proceso que está usando el puerto especificado usando el comando `netstat`
2. Termina el proceso usando el comando `taskkill`
3. Si el proceso no puede ser terminado normalmente, usa el parámetro `-f` para forzar su terminación

## Precauciones

- Se requieren privilegios de administrador para terminar ciertos procesos
- Usar con precaución, asegurándose de que el proceso a terminar no sea un proceso crítico del sistema

## Licencia

Este proyecto está bajo la Licencia Creative Commons Atribución-NoComercial 4.0 Internacional (CC BY-NC 4.0).

Usted es libre de:
- Compartir — copiar y redistribuir el material en cualquier medio o formato
- Adaptar — remezclar, transformar y construir a partir del material

Bajo los siguientes términos:
- Atribución — Debe dar crédito de manera adecuada, brindar un enlace a la licencia, e indicar si se han realizado cambios
- NoComercial — No puede utilizar el material con propósitos comerciales sin el permiso explícito del autor

Más información: https://creativecommons.org/licenses/by-nc/4.0/deed.es
