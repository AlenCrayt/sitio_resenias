package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

type resenia struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo_libro"`
	Parrafo     string `json:"resenia_parrafo"`
	Link_Imagen string `json:"link_portada"`
}

func main() {
	//Abrimos una conexión a la base de datos e ingresamos con un usuario
	sign_in := mysql.Config{
		User:                 "Test",
		Passwd:               "1234",
		AllowNativePasswords: true,
		DBName:               "sitio_resenias",
		Net:                  "tcp",
		Addr:                 "127.0.0.1",
	}
	base_datos, err := sql.Open("mysql", sign_in.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	error_conexion := base_datos.Ping()
	if error_conexion != nil {
		log.Fatal(error_conexion)
	}
	fmt.Println("Conexión con base de datos establecida")

	http.HandleFunc("/resenias-generales", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		leer_resenias(w, base_datos)
	})
	http.HandleFunc("/resenias-nuevas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//Usamos un condicional para ejecutar la funcion de agregar Reseñas solo si recibimos un pedido de metodo POST a esta URL
		if (r.Method == "POST") {
			agregar_resenia(w, r, base_datos)
		}
	})
	http.HandleFunc("/resenias-especificas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		buscar_resenia(w, r, base_datos)
	})

	//pagina := http.FileServer(http.Dir("./pagina_web"))
	//http.Handle("/", pagina)

	http.ListenAndServe(":8080", nil)
}

func leer_resenias(respuesta http.ResponseWriter, bd *sql.DB) {
	fmt.Println("se llama a leer reseñas")
	var lista_resenias []resenia
	resenias_bd, _ := bd.Query("SELECT * FROM resenias LIMIT 25")
	for resenias_bd.Next() {
		var temp_resenia resenia
		err := resenias_bd.Scan(&temp_resenia.ID, &temp_resenia.Titulo, &temp_resenia.Parrafo, &temp_resenia.Link_Imagen)
		if err != nil {
			fmt.Println(err)
		}
		lista_resenias = append(lista_resenias, temp_resenia)
	}
	respuesta.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respuesta).Encode(lista_resenias)
}

func agregar_resenia(respuesta http.ResponseWriter, pedido *http.Request, bd *sql.DB) {
	var resenia_subida resenia

	fmt.Println("se esta agregando una reseña")

	datos_recibidos, err := io.ReadAll(pedido.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer pedido.Body.Close()

	json.Unmarshal(datos_recibidos, &resenia_subida)
	id, err := bd.Exec("INSERT INTO resenias (titulo_libro, parrafo_resenia, link_img_portada) VALUES(?, ?, ?)", resenia_subida.Titulo, resenia_subida.Parrafo, resenia_subida.Link_Imagen)
	if err != nil {
		fmt.Println(err)
	}
	id_usable, _ := id.LastInsertId()
	texto_id := strconv.Itoa(int(id_usable))
	respuesta.Write([]byte(texto_id))
}

/*
Se va a necesitar una extension extensiva del código para poder manejar creación y autenticación de usuarios y vincular a esos usuarios
a reseñas especificas en las que van a tener permisos para actualizarlas o borrarlas
Va a haber que tener dos tablas vinculadas una de usuarios y una de reseñas
*/

func buscar_resenia(respuesta http.ResponseWriter, pedido *http.Request, bd *sql.DB) {
	var lista_resenias []resenia
	fmt.Println("buscando reseña")
	busqueda := pedido.URL.Query().Get("buscar-titulo")
	fmt.Println(busqueda)
	resultado, err := bd.Query("SELECT * FROM resenias WHERE titulo_libro LIKE ?", "%" + busqueda + "%")
	if err != nil {
		fmt.Println(err)
	}
	for resultado.Next() {
		var temporal_resenia resenia
		wrong := resultado.Scan(&temporal_resenia.ID, &temporal_resenia.Titulo, &temporal_resenia.Parrafo, &temporal_resenia.Link_Imagen)
		if wrong != nil {
			fmt.Println(wrong)
		}
		lista_resenias = append(lista_resenias, temporal_resenia)
	}
	respuesta.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respuesta).Encode(lista_resenias)
}
