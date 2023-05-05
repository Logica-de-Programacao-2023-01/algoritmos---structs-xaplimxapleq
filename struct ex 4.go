package main

type playlists struct {
	musicas musicas

}
type musicas struct {
	titulo  string
	artista string
	duracao int
}

func main() {
	m := musicas{
	titulo: "holiday",
	artista: "lil nas x"
	duracao: 350
}
