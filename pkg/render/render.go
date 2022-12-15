package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*func RenderTemplatesTEst(w http.ResponseWriter, tmpl string) {
	parsed, _ := template.ParseFiles("./templates/" + tmpl,"./templates/base.layout.tmpl")
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("There is error", err)
		return
	}
}*/

//To stote the template in the cache so that each time we render our template it does not go to disk to render.

var tc=make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error
	_, inMap:=tc[t]
	if !inMap{
		log.Println("Template is not present in the cache")
		err=createTemplateCache(t)
		if err!=nil{
			fmt.Println(err)
			return
		}
	} else{
		log.Println("Template is present in the cache and it is used form the cache")

	}
	tmpl=tc[t];
	err=tmpl.Execute(w,nil)
	if err!=nil{
		fmt.Println(err)
		return;
	}
		
	
}

func createTemplateCache(t string) error{
	templates:=[]string{
		fmt.Sprintf("./templates/%s",t),
		"./templates/base.layout.tmpl",
	}
	tmpl,err:=template.ParseFiles(templates...)
	if err!=nil{
		fmt.Println(err)
		return err

	}
	tc[t]=tmpl
	return nil

}