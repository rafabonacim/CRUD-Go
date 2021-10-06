package main

import (
	//transformar a lista de struct em uma string de formato json
	"encoding/json"
	//para printar
	"fmt"
	//importado para logar o erro
	"log"
	"math/rand"
	"strconv"
	//para usar o pacote http
	"net/http"
	"github.com/gorilla/mux"
) 
	

type Customer struct {
	//json:"id" = codificar e decodificar em letra minusca
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

var customers []Customer
// 2 argumentos=um  escritor de reposta e um ponteiro de requisição
func getCustomers(w http.ResponseWriter, r *http.Request)  {
	// primeiro argumento é o w para escrever no escritor de resposta
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w) gera um encoder
	//Encode(customers) é o que passamos que seja codificado em json e passado p cliente
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _,item := range customers{
		if item.Id == params["id"]{
			//Encode é o que passamos que seja codificado em json e passado p cliente
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var customer Customer
	//_ignorar o resultado da func. para não checar o erro
	//decoder pega a string que veio da request e passa para objeto 
	_ = json.NewDecoder(r.Body).Decode(&customer)
	//atribui ao costumer id um numero 
	// int para string
	customer.Id = strconv.Itoa(rand.Intn(100000))
	//acrescenta o custumer criado ao objeto
	customers = append(customers, customer)
	//encoder pega o objeto  e passa para string p usuario visualizar 
	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	//pego o request
	params := mux.Vars(r)
	for index,item := range customers{
		//se p pametro = ao id da requisição
		if item.Id == params["id"]{
			//rafaela index=0,vazio,dadonas ate dadonas
			//dadonas index=1, rafa , vazio
			// exclui do array o index que quer editar
			customers = append(customers[:index],customers[index+1:]...)
			var customer Customer
			//decoder pega a string que veio da request e passa para objeto 
			_ = json.NewDecoder(r.Body).Decode(&customer)
			//atribui o id do params ao customer.id
			customer.Id =  params["id"]
			//inclui o customer editado a lista de customer
			customers = append(customers, customer)
			//encoder pega o objeto  e passa para string p usuario visualizar 
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range customers{
		if item.Id == params["id"]{
			//retiro o customer e já pausa a função
			customers = append(customers[:index],customers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(customers)
}

func main(){
	//servidor mux 
	router := mux.NewRouter()
	customers = append(customers, Customer{
		Id:"1",
		Name:"Rafaela Bonacim",
		Email:"rafaelabonacim@gmail.com",
	})
	customers = append(customers, Customer{Id:"2", Name:"Dadonas", Email:"dadonas@gmail.com"})
	
	//chama a função handlefunc que leva 2 argumentos
	router.HandleFunc("/listarclientes", getCustomers).Methods("GET")
	router.HandleFunc("/listarcliente/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/criarcliente", createCustomer).Methods("POST")
	router.HandleFunc("/atualizarcliente/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/deletarcliente/{id}",deleteCustomer).Methods("DELETE")
 	fmt.Println("Conectado na porta 8000")
	//para rodar o servidor de http usando go:"http.ListenAndServe"
	//router é o 2 argumento servidor mux
	log.Fatal(http.ListenAndServe(":8000",router))
}