
package main

import (
    "fmt"
    "log"
    "strings"
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "encoding/json"
)

type User struct {
    Username string
    Password string
}

var users = make(map[string]User)
var hasher = bcrypt.DefaultCost

func signUp(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Check if the username is already taken
    if _, exists := users[user.Username]; exists {
        http.Error(w, "Username already taken", http.StatusBadRequest)
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), hasher)
    if err != nil {
        http.Error(w, "Error while hashing the password", http.StatusInternalServerError)
        return
    }

    // Store the user with the hashed password
    user.Password = string(hashedPassword)
    users[user.Username] = user

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "User signed up successfully")
}

func signIn(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Check if the user exists
    storedUser, exists := users[user.Username]
    if !exists {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // Compare the provided password with the stored hashed password
    err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "User signed in successfully")
}
