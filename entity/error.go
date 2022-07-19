package entity

import "errors"

// Geral
var ErrParametroInvalido = errors.New("Parâmetro(s) inválido(s)")

// Video
var ErrVideoNotFound = errors.New("Vídeo não encontrado")
var ErrVideoInexistente = errors.New("Vídeo Inexistente")
var ErrIdPathEBodyDiferentes = errors.New("ID informado no path é diferente do ID informado no request body")

// Categoria
var ErrCategoriaNotFound = errors.New("Categoria não encontrado")
var ErrCategoriaInexistente = errors.New("Categoria Inexistente")
