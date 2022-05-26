package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSuccessResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

// Jwt key yang akan dipakai untuk membuat signature
var jwtKey = []byte("key")

// Struct claim digunakan sebagai object yang akan di encode oleh jwt
// jwt.StandardClaims ditambahkan sebagai embedded type untuk provide standart claim yang biasanya ada pada JWT
type Claims struct {
	Username string
	Role     string
	jwt.StandardClaims
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.usersRepo.Login(user.Username, user.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	userRole, err := api.usersRepo.GetUserRole(*res)

	// Task: 1. Deklarasi expiry time untuk token jwt
	//       2. Buat claim menggunakan variable yang sudah didefinisikan diatas
	//       3. expiry time menggunakan time millisecond

	// TODO: answer here
	expJWT := time.Now().Add(time.Minute * 20)
	claim := Claims{
		Username: *res,
		Role:     *userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expJWT.Unix(),
		},
	}

	// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai

	// TODO: answer here
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
	//       2. return internal error ketika ada kesalahan ketika pembuatan JWT string

	// TODO: answer here
	jwtString, err := myToken.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	// Task: Set token string kedalam cookie response

	// TODO: answer here
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtString,
		Expires: expJWT,
		Path:    "/",
	})

	// Task: Return response berupa username dan token JWT yang sudah login
	json.NewEncoder(w).Encode(LoginSuccessResponse{
		Username: *res,
		Token:    jwtString,
	})
	// json.NewEncoder(w).Encode(LoginSuccessResponse{Username: "", Token: ""}) // TODO: replace this
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	token, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// return unauthorized ketika token kosong
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// return bad request ketika field token tidak ada
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)

	// set status dan tampilkan pesan logout
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logged out"))
	// encoder.Encode(AuthErrorResponse{Error: ""}) // TODO: replace this
}
