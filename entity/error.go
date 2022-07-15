package entity

import "errors"

var ErrNotFound = errors.New("Vídeo não encontrado")
var ErrParametroInvalido = errors.New("Parâmetro(s) inválido(s)")
var ErrVideoInexistente = errors.New("Vídeo Inexistente")
var ErrIdPathEBodyDiferentes = errors.New("ID informado no path é diferente do ID informado no request body")
