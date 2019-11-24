package main

//import "fmt"
import (
  "fmt"
  "database/sql"
_ "github.com/go-sql-driver/mysql"
  "io/ioutil"
  "os"
  "time"
  "math/rand"
  "strconv"
  "math"
)


func obtenerConexiones()  [5]*sql.DB {
	var conexiones [5]*sql.DB
	var ips = [5]string{"127.0.0.1","127.0.0.1","127.0.0.1","127.0.0.1","127.0.0.1"};
	for i:=0; i < 5; i++ {
		db, err := sql.Open("mysql", "root@tcp(" + ips[i] +":3306)/pruebas_inv")
		if  err != nil {
			panic("Fallo una de las conexiones");
		} else {
			conexiones[i] = db
		}
	}
	return conexiones;
}

func determinarRango(i float64) int {
	if i >= 0 &&  i < 1 {
		return 1
	} else if i >= 1 && i < 2 {
		return 2
	} else if i >= 2 && i < 3 {
		return 3
	} else if i >= 3 && i < 4 {
		return 4
	} else if i >= 4 && i <= 5 {
		return 5
	}
	return 0;
}

func main() {
	dat, err := ioutil.ReadFile("comentario_prueba.txt")
	tiempos, err := os.Create("tiempos.txt")
	distribucion := [5]int32 {0,0,0,0,0}
	conexiones := obtenerConexiones();
	if err != nil {
		panic("El comentario de prueba no esta")
	}
	texto := string(dat)
	nodo_seleccionado := 0;
	for i:= 0; i < 5000; i++ {
		nodo_seleccionado = int(math.Floor(5 * rand.Float64()))
		fmt.Println(nodo_seleccionado)
		node_id := "0" + strconv.Itoa(nodo_seleccionado + 1)
		distribucion[nodo_seleccionado]++;
		start := time.Now()
		row, errSQL := conexiones[nodo_seleccionado].Query("INSERT INTO post VALUES(null, ?, '0123456789012345678901234', ?, 0, '2017-07-20')", node_id, texto)
		tiempos.WriteString("" + fmt.Sprintf("%f", float64(time.Since(start) / time.Millisecond)) + "\n")
		if errSQL != nil {
			panic(errSQL.Error())
		}
		row.Close()
	}
	defer tiempos.Close()
	fmt.Println("Terminado")
	fmt.Println("Distribucion:")
	fmt.Println(distribucion)
}