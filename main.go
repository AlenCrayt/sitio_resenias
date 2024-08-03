// Código escrito por Agustin Santisteban, Usuario de Github: AlenCrayt
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

//El nombre resenia se usa por cuestiones de posibles incompatibilidades con programas que no procesan la ñ
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
	//Hacemos un Ping() a la base de datos para comprobar que la conexión funciona correctamente
	error_conexion := base_datos.Ping()
	if error_conexion != nil {
		log.Fatal(error_conexion)
	}
	fmt.Println("Conexión con base de datos establecida")

	/*Tenemos tres funciones que manejan pedidos HTTP de diferentes operaciones para URLs especificas, las funciones regulares del
	Paquete HTTP tienen adentro suyo a otra tres funciones propias que permiten agregar el puntero a la base de datos como parámetro*/
	http.HandleFunc("/resenias-generales", func(w http.ResponseWriter, r *http.Request) {
		//Los pedidos de otros orígenes solo están activados durante el desarrollo de la pagina
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		leer_resenias(w, r, base_datos)
	})
	http.HandleFunc("/resenias-nuevas", func(w http.ResponseWriter, r *http.Request) {
		/*w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")*/
		//Usamos un condicional para ejecutar la función de agregar Reseñas solo si recibimos un pedido de método POST a esta URL
		if (r.Method == "POST") {
			agregar_resenia(r, base_datos)
		}
	})
	http.HandleFunc("/resenias-especificas", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		buscar_resenia(w, r, base_datos)
	})

	pagina := http.FileServer(http.Dir("./pagina_web"))
	http.Handle("/", pagina)

	//Arrancamos el servidor y escuchamos en el puerto :8080
	http.ListenAndServe(":8080", nil)
}

func leer_resenias(respuesta http.ResponseWriter, pedido *http.Request , bd *sql.DB) {
	fmt.Println("se llama a leer reseñas")

	//Checkeamos la cantidad de entradas en la tabla resenias de la base de datos
	var cantidad_de_entradas int
	resultado_inicial := bd.QueryRow("SELECT COUNT(*) FROM resenias")
	resultado_inicial.Scan(&cantidad_de_entradas)

	//extraemos el parámetro de indice pasado a traves de la URL y lo convertimos en un entero
	indice_string := pedido.URL.Query().Get("indice")
	indice_numerico, _ := strconv.Atoi(indice_string)

	//mediante un condicional determinamos si el indice que se provee esta afuera del rango del numero de entradas en la tabla
	if indice_numerico < 0 ||  indice_numerico >= cantidad_de_entradas{
		fmt.Println("El indice provisto esta fuera del rango de entradas en la tabla resenias")
		//Enviamos un código de estado 404 en la respuesta
		respuesta.WriteHeader(http.StatusNotFound)
		return
	}

	//Creamos un slice de structs resenia para poder manipular los datos obtenidos del pedido a la base de datos
	var lista_resenias []resenia
	//Recibimos las entradas a base de un pedido SQL a la tabla resenias
	resenias_bd, _ := bd.Query("SELECT * FROM resenias ORDER BY ID_resenia DESC LIMIT 5 OFFSET ?", indice_numerico)
	//Usamos for para iterar a traves de las entradas que recibimos de nuestro pedido SQL
	for resenias_bd.Next() {
		//Creamos una variable del tipo de struct resenia para temporalmente guardar los datos que sacamos de la lista de entradas de la Base de datos
		var temp_resenia resenia
		err := resenias_bd.Scan(&temp_resenia.ID, &temp_resenia.Titulo, &temp_resenia.Parrafo, &temp_resenia.Link_Imagen)
		if err != nil {
			fmt.Println(err)
		}
		//Agregamos esta variable al final del slice de resenias
		lista_resenias = append(lista_resenias, temp_resenia)
	}
	//Codificamos la respuesta del servidor como JSON y lo enviamos al cliente
	respuesta.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respuesta).Encode(lista_resenias)
}

func agregar_resenia(pedido *http.Request, bd *sql.DB) {
	fmt.Println("se esta agregando una reseña")

	//Usamos una variable de tipo reseña para tener temporalmente los datos decodificados del paquete JSON enviado por el cliente
	var resenia_subida resenia
	datos_recibidos, err := io.ReadAll(pedido.Body)
	if err != nil {
		fmt.Println(err)
	}
	//Evitamos cerrar el cuerpo del pedido HTTP hasta que se haya terminado de ejecutar la función
	defer pedido.Body.Close()

	json.Unmarshal(datos_recibidos, &resenia_subida)
	//Ejecutamos un pedido SQL para Insertar una nueva entrada con los datos de la nueva reseña
	_, err = bd.Exec("INSERT INTO resenias (titulo_libro, parrafo_resenia, link_img_portada) VALUES(?, ?, ?)", resenia_subida.Titulo, resenia_subida.Parrafo, resenia_subida.Link_Imagen)
	if err != nil {
		fmt.Println(err)
	}
}

func buscar_resenia(respuesta http.ResponseWriter, pedido *http.Request, bd *sql.DB) {
	fmt.Println("buscando reseña")

	var lista_resenias []resenia
	busqueda := pedido.URL.Query().Get("buscar-titulo")
	//Ejecutamos una búsqueda en la tabla resenias a base del termino enviado por el cliente
	resultado, err := bd.Query("SELECT * FROM resenias WHERE titulo_libro LIKE ?", "%" + busqueda + "%")
	if err != nil {
		fmt.Println(err)
	}
	//Realizamos el mismo proceso que en leer_resenias() simplemente con un diferente resultado obtenido a base de la búsqueda
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
