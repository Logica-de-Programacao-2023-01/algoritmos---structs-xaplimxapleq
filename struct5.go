
func main() {
// Criar playlists
playlists := []Playlist{
{
nome: "playlist 1",
musicas: []Musica{
{
titulo:  "musica 1",
artista: "artista 1",
duracao: 10,
},
},
},
{
nome: "playlist 2",
musicas: []Musica{
{
titulo:  "musica 2",
artista: "artista 2",
duracao: 15,
},
{
titulo:  "musica 3",
artista: "artista 3",
duracao: 20,
},
},
},
}

// Encontrar músicas em playlists
resultado := encontrarMusica(playlists, "musica 1")
fmt.Println("Músicas encontradas:")
for _, playlist := range resultado {
for _, musica := range playlist.musicas {
fmt.Printf("Título: %s, Artista: %s, Duração: %d\n", musica.titulo, musica.artista, musica.duracao)
}
}
}
