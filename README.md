<h1 align="center">Cola</h1>

El contenido de este repositorio es la implementación de una cola.

Implementado mediante el uso de una cola enlazada.

---
## Contenido:
- [Requisitos previos](#requisitos-previos)
- [Pasos para levantar el repo](#pasos-para-levantar-el-repo)
- [Estructura del repo](#estructura-del-repo)
- [Enunciado del TP](#enunciado-del-tp)
- [Créditos](#créditos)
---

## Requisitos Previos:
- [Tener Git instalado](https://git-scm.com/book/es/v2/Inicio---Sobre-el-Control-de-Versiones-Instalaci%C3%B3n-de-Git)
- [El repo debe estar dentro de una carpeta llamada "tdas".](#estructura-del-repo)

## Pasos para levantar el repo:

### 1- Clona el repositorio:

```
git clone git@github.com:mikuu-montes/cola.git

``` 
o 

```
git clone https://github.com/mikuu-montes/cola.git
``` 
(te va a pedir que coloques tu contraseña)

### 2- Moverse a la carpeta del proyecto:
```
cd cola
```


## Estructura del repo:
Explicamos que contiene cada archivo:
- ` cola.go `: En este se declara la interfaz de la cola.
- ` cola_enlazada.go `: En este se implementan las primitivas de la interfaz antes declarada.
- ` cola_test.go `: En este se encuentran las pruebas que verifican que lo implementado, funcione efectivamente como una cola.

Para usarse deben tener esta estructura de archivos:

```
tdas/
└── cola/ ← (repositorio actual)
    ├── cola.go
    ├── cola_enlazada.go
    └── cola_test.go

```

## Enunciado del TP:
- [Link al enunciado de la cátedra.](https://algoritmos-rw.github.io/algoritmos/tps/cola/)

## Créditos:
Este trabajo fue realizado como tarea para la materia de **Algoritmos y Estructuras de Datos**, cátedra **Buchwald**, en **FIUBA**, por la alumna:
- **Montes Brisa Micaela**