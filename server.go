package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func endpoint(w http.ResponseWriter, r *http.Request) {
	var request map[string]any

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&request); err != nil {
		w.Write([]byte(err.Error()))
	}

	// body
	for _, body := range engine.Bodies {
		v, ok := request[body.Name]
		if !ok {
			fmt.Fprintf(w, "body tidak valid")
			return
		}

		switch body.Datatype {
		case "str":
			types := reflect.TypeOf(v).Name()
			if types != "string" {
				fmt.Fprintf(w, "bukan str")
				return
			}

		case "number":
			types := reflect.TypeOf(v).Name()
			if types != "float64" {
				fmt.Fprintf(w, "bukan number")
				return
			}

			toFloat64, _ := v.(float64)
			request[body.Name] = uint64(toFloat64)

		case "boolean":
			types := reflect.TypeOf(v).Name()
			if types != "bool" {
				fmt.Fprintf(w, "bukan boolean")
				return
			}
		}
	}

	for _, rules := range engine.Rules {
		valid := true

		for j, dict := range engine.Dictionaries {
			r := request[dict.Attribute]
			if rules.Value[j] == nil {
				continue
			}

			switch r.(type) {
			case uint64:
				r := r.(uint64)
				v := rules.Value[j].(uint64)

				switch dict.Operator {
				case ">":
					if r < v {
						valid = false
					}
				case "<":
					if r > v {
						valid = false
					}
				case "=":
					if r != v {
						valid = false
					}
				case ">=":
					if r < v {
						valid = false
					}
				case "<=":
					if r > v {
						valid = false
					}
				}
			case string:
				r := r.(string)
				v := rules.Value[j].(string)

				if r != v {
					valid = false
				}
			case bool:
				r := r.(bool)
				v := rules.Value[j].(bool)

				if r != v {
					valid = false
				}
			}
		}

		if valid {
			res := rules.Action.(uint64)
			w.Write([]byte(fmt.Sprint(res)))
			return
		}
	}

	w.Write([]byte("semua rules gagal"))
}
