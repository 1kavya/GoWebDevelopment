package main
import ("fmt"
	   "net/http"
	"github.com/gorilla/mux")
	   
func main(){
	//mux:=&http.ServeMux{}
	//mux.HandleFunc("/",handler)
	//http.HandleFunc("/index",index)
//	http.HandleFunc("/",home)
r:=mux.NewRouter()
r.HandleFunc("/",home)
r.HandleFunc("/index",index)
	http.ListenAndServe(":3000",r)		
	}

	func home(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"You are in home page.")
	}

	func index(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "text/html")	 
		fmt.Fprint(w,"You are in index page")
	}

			 	
	


	}

