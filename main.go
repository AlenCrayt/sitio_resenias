package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

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
		agregar_resenia(w, r, base_datos)
	})
	http.HandleFunc("/resenias-especificas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		buscar_resenia(w, r, base_datos)
	})
	http.HandleFunc("/usuarios-nuevos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		registrar_usuario(r, base_datos)
	})

	//pagina := http.FileServer(http.Dir("./pagina_web"))
	//http.Handle("/", pagina)

	http.ListenAndServe(":8080", nil)
}

func leer_resenias(respuesta http.ResponseWriter, bd *sql.DB) {
	fmt.Println("se llama a leer reseñas")
	type resenia struct {
		ID          int    `json:"id"`
		Titulo      string `json:"titulo"`
		Parrafo     string `json:"parrafo"`
		Link_Imagen string `json:"link_imagen"`
	}
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
	fmt.Println("se esta agregando una reseña")
	titulo := pedido.FormValue("titulo")
	parrafo := pedido.FormValue("parrafo")
	link_imagen := pedido.FormValue("link_img")
	id, err := bd.Exec("INSERT INTO resenias (titulo_libro, parrafo, link_imagen) VALUES(?, ?, ?)", titulo, parrafo, link_imagen)
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
func actualizar_resenia(pedido *http.Request, bd *sql.DB) {
	//Como obtener el ID de una reseña en el frontend?
	fmt.Println("se esta actualizando una reseña")
	titulo := pedido.FormValue("titulo")
	parrafo := pedido.FormValue("parrafo")
	link_imagen := pedido.FormValue("link_img")
	//Necesito saber que partes de la reseña estan siendo actualizada? O podria recibir todos los datos incluso aquellos que no son actualizados
	_, err := bd.Exec("UPDATE resenias SET Titulo = ?, Parrafo = ?, Imagen = ? WHERE ID = ?", titulo, parrafo, link_imagen)
	if err != nil {
		fmt.Println(err)
	}
}

func borrar_resenia(pedido *http.Request, bd *sql.DB) {
	fmt.Println("se esta eliminando una reseña")
}

func buscar_resenia(respuesta http.ResponseWriter, pedido *http.Request, bd *sql.DB) {
	type resenia struct {
		ID          int    `json:"id"`
		Titulo      string `json:"titulo"`
		Parrafo     string `json:"parrafo"`
		Link_Imagen string `json:"link_imagen"`
	}
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

func registrar_usuario(pedido *http.Request, bd *sql.DB) {
	fmt.Println("Registrando usuario nuevo")
	nombre_a_registrar := pedido.FormValue("registro_nombre")
	contrasenia_a_registrar := pedido.FormValue("registro_contrasenia")
	_, err := bd.Exec("INSERT INTO usuarios (nombre_usuario, contrasenia) VALUES(?, ?)", nombre_a_registrar, contrasenia_a_registrar)
	if err != nil {
		fmt.Println(err)
	}
}

func ingresar_usuario(pedido *http.Request, bd *sql.DB) {
	fmt.Println("Un usuario esta ingresando")
}