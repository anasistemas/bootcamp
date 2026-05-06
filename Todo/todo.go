package todo

import (
	"errors"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(tarea string) {
	*l = append(*l, item{
		Task:      tarea,
		CreatedAt: time.Now(),
	})
}

func (l *List) Complete(indice int) error {
	lista := *l
	if indice <= 0 || indice > len(lista) {
		return errors.New("índice inválido")
	}
	lista[indice-1].Done = true
	lista[indice-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(indice int) error {
	lista := *l
	if indice <= 0 || indice > len(lista) {
		return errors.New("índice inválido")
	}
	*l = append(lista[:indice-1], lista[indice:]...)
	return nil
}
