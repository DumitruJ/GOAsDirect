package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtKey = []byte("your_secret_key") // Change this in production

// User struct for demonstration
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	// Connect to PostgreSQL
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=astrocms_dev sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Database connection failed")
			return
		}
		fmt.Fprintln(w, "OK")
	})

	// Register endpoint: create a new user with hashed password
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil || creds.Username == "" || creds.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Hash the password
		hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Insert user into DB
		_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", creds.Username, string(hash))
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintln(w, "Username already exists")
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "User registered successfully")
	})

	// JWT login endpoint (real auth)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Fetch user from DB
		var hash string
		err = db.QueryRow("SELECT password_hash FROM users WHERE username=$1", creds.Username).Scan(&hash)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Compare password
		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(creds.Password))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Create JWT token
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Username: creds.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"token": "%s"}`, tokenString)
	})

	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "AstroCMS Backend API")
	})

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
