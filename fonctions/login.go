package fonctions
import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
	"golang.org/x/crypto/bcrypt"
)
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func login(w http.ResponseWriter, rec *http.Request){
	if rec.Method == http.MethodPost {
		username := rec.FormValue("username")
		password := rec.FormValue("password")
		var motDePasse string
		err := db.DB.QueryRow("SELECT password FROM utilisateurs WHERE username = ?", username).Scan(&hashedPassword)
		tpl.ExecuteTemplate(w, "login.gohtml", nil)

		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
			http.SetCookie(w, &http.Cookie{
				Name:  "username",
				Value: username,
				Path:  "/",

			})
		}
	}
}
