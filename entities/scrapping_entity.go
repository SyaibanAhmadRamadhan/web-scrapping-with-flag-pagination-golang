package entities

type Scrapping struct {
	Judul          string `bson:"judul"`
	NamaPenulis    string `bson:"nama_penulis"`
	TanggalArtikel string `bson:"tanggal_artikel"`
	Gambar         string `bson:"gambar"`
	IsiArtikel     string `bson:"isi_artikel"`
}
