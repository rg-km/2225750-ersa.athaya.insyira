package main

import "fmt"

// Implementasikan MusicPlayer yang merupakan representasi dari pemutar musik.
// MusicPlayer memiliki atribut Playlist. Playlist merupakan daftar lagu yang tersimpan.
// Sebagai gambaran, seperti pada Spotify, playlist merepresentasikan lagu-lagu yang disimpan dalam satu playlist yang sama.
// Lagu direpresentasikan dalam objek Song yang memiliki dua atribut yaitu Singer dan Title.
// Di Playlist juga terdapat atribut IsRepeatable yang merepresentasikan apakah playlist tersebut akan diputar ulang jika lagu telah habis.

// Lengkapilah tiga method yang ada pada MusicPlayer yaitu AddSong, Play, dan Repeat.
// AddSong merupakan method yang menambahkan lagu pada playlist.
// Play merupakan method yang memutar lagu pada playlist. Pada method ini, cetaklah sebuah wording dengan format `Now playing [Singer] - [Title]`
// 		Sebagai contoh, jika lagu One Direction dengan judul Night Changes diputar, maka akan mengeluarkan output `Now playing One Direction - Night Changes`
// Repeat merupakan method yang memungkinkan playlist diputar ulang jika lagu telah habis.
// Jika playlist dilakukan repeat, maka ketika lagu telah diputar semua maka akan mengulang pada urutan semula.
type Song struct {
	Singer string
	Title  string
}

type Playlist struct {
	Songs        []Song
	IsRepeatable bool
}

func NewPlaylist() Playlist {
	return Playlist{
		Songs:        []Song{},
		IsRepeatable: false,
	}
}

func (pl Playlist) IsEmpty() bool {
	return len(pl.Songs) == 0
}

type MusicPlayer struct {
	Playlist Playlist
}

func NewMusicPlayer() MusicPlayer {
	return MusicPlayer{
		Playlist: NewPlaylist(),
	}
}

func (mp *MusicPlayer) AddSong(song Song) {
	// kita tambahkan musik ke playlist
	mp.Playlist.Songs = append(mp.Playlist.Songs, song)
}

func (mp *MusicPlayer) Play() string {
	// daftar lagu diplaylist
	songs := mp.Playlist.Songs

	// kita cek apakah ada lagu di playlist
	if mp.Playlist.IsEmpty() {
		return ""
	}

	// jika ada musik di playlist maka putar lagu dan perhatikan repateable atau tidak
	// jika tidak repeatable maka putar lagu pertama dan kosongkan playlist

	// jika non-repeatale
	if !mp.Playlist.IsRepeatable {

		// putar lagu pertama
		song := mp.Playlist.Songs[:1]
		notif := fmt.Sprintf("Now playing %s - %s", song[0].Singer, song[0].Title)

		// kosongkan playlist jika hanya ada satu lagu yang diputar
		if len(songs) == 1 {
			mp.Playlist.Songs = []Song{}
		} else {
			mp.Playlist.Songs = songs[1:]
		}
		return notif
	} else {
		// jika repeataable
		// putar lagu pertama
		song := mp.Playlist.Songs[:1]
		notif := fmt.Sprintf("Now playing %s - %s", song[0].Singer, song[0].Title)

		// sisakan lagu itu sendiri di playlist jika hanya ada 1 lagu yang diputar
		if len(songs) == 1 {
			mp.Playlist.Songs = mp.Playlist.Songs[:1]
		} else {
			// set len
			mp.Playlist.Songs = songs[:]
			// pindahkan lagu yang sedang diputar ke urutan terakhir
			mp.Playlist.Songs = append(mp.Playlist.Songs[1:], mp.Playlist.Songs[:1]...)
		}
		return notif
	}
}

func (p *Playlist) Repeat() {
	// ketika ini dipanggil maka ubah isRepeatable menjadi true
	p.IsRepeatable = true
}

func main() {
	// bikin inisiasi dari MusicPlayer
	var player MusicPlayer

	// tambhakan lagu
	player.AddSong(Song{
		Singer: "Tulus",
		Title:  "Hati-Hati di Jalan",
	})
	player.AddSong(Song{
		Singer: "TREASURE",
		Title:  "DARARI",
	})
	player.AddSong(Song{
		Singer: "NIKI",
		Title:  "Every Summertime",
	})
	player.AddSong(Song{
		Singer: "Pamungkas",
		Title:  "To The Bone",
	})

	// repeat playlist
	player.Playlist.Repeat()

	// play song
	a := player.Play()
	fmt.Println(a)

}
