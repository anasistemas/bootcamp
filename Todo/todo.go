package todo

import (
	"encoding/json"
	"errors"
	"os"
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

func (l *List) Save(archivo string) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(archivo, data, 0644)
}

func (l *List) Get(archivo string) error {
	data, err := os.ReadFile(archivo)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("archivo no encontrado")
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, l)
}
