package seed

import (
	"log"

	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	superAdmin := &entities.Role{Name: "Admin"}
	if err := db.Create(superAdmin).Error; err != nil {
		log.Fatalf("error insert role %s", err.Error())
	}
	userSuperAdmin := &entities.User{Username: "admin", Password: "admin", Active: true, IsOnline: false, RoleUUID: superAdmin.UUID}
	if err := db.Create(userSuperAdmin).Error; err != nil {
		log.Fatalf("error insert user %s", err.Error())
	}
	provinsi := &entities.Provinsi{
		Provinsi: "Nusa Tenggara Barat", Kode: 52,
	}
	if err := db.Create(provinsi).Error; err != nil {
		log.Fatalf("error insert provinsi %s", err.Error())
	}
	kabupaten := []*entities.Kabupaten{
		{Kabupaten: "Lombok Tengah", Kode: 5202, ProvinsiKode: 52},
	}
	if err := db.Save(kabupaten).Error; err != nil {
		log.Fatalf("error insert kabupaten %s", err.Error())
	}
	kecamatan := []*entities.Kecamatan{
		{Kecamatan: "Pujut", Kode: 520204, KabupatenKode: 5202},
		{Kecamatan: "Pringgarata", Kode: 520208, KabupatenKode: 5202},
		{Kecamatan: "Praya Timur", Kode: 520206, KabupatenKode: 5202},
		{Kecamatan: "Praya Tengah", Kode: 520210, KabupatenKode: 5202},
		{Kecamatan: "Praya Barat Daya", Kode: 520211, KabupatenKode: 5202},
		{Kecamatan: "Praya Barat", Kode: 520205, KabupatenKode: 5202},
		{Kecamatan: "Praya", Kode: 520201, KabupatenKode: 5202},
		{Kecamatan: "Kopang", Kode: 520209, KabupatenKode: 5202},
		{Kecamatan: "Jonggat", Kode: 520202, KabupatenKode: 5202},
		{Kecamatan: "Janapria", Kode: 520207, KabupatenKode: 5202},
		{Kecamatan: "Batukliang Utara", Kode: 520212, KabupatenKode: 5202},
		{Kecamatan: "Batukliang", Kode: 520203, KabupatenKode: 5202},
	}
	if err := db.Save(kecamatan).Error; err != nil {
		log.Fatalf("error insert %s", err.Error())
	}
	desa := []*entities.Desa{}
	jonggat := []string{
		"Barejulat",
		"Batutulis",
		"Bonjeruk",
		"Bunkate",
		"Gemel",
		"Jelantik",
		"Labulia",
		"Nyerot",
		"Pengenjek",
		"Perina",
		"Puyung",
		"Sukarara",
		"Ubung",
	}
	for _, data := range jonggat {
		desa = append(desa, &entities.Desa{Desa: data, KecamatanKode: 520202})
	}

	janapria := []string{
		"Bakan",
		"Durian",
		"Janapria",
		"Jango",
		"Kerembong",
		"Langko",
		"Lekor",
		"Loangmaka",
		"Pendem",
		"Saba",
		"Selebung",
		"Rembiga",
		"Setuta",
	}
	for _, data := range janapria {
		desa = append(desa, &entities.Desa{Desa: data, KecamatanKode: 520207})
	}

	batut := []string{
		"Aik Berik",
		"Aik Bukak",
		"Karang Sidemen",
		"Lantan",
		"Mas Mas",
		"Setiling",
		"Tanak Beak",
		"Teratak",
	}
	for _, data := range batut {
		desa = append(desa, &entities.Desa{Desa: data, KecamatanKode: 520212})
	}

	batukliang := []string{
		"Aik Darek",
		"Barabali",
		"Beber",
		"Bujak",
		"Mantang",
		"Mekar Bersatu",
		"Pagutan",
		"Peresak",
		"Selebung",
		"Tampak Siring",
	}

	for _, data := range batukliang {
		desa = append(desa, &entities.Desa{Desa: data, KecamatanKode: 520203})
	}

	if err := db.Save(desa).Error; err != nil {
		log.Fatalf("error insert %s", err.Error())
	}

}
