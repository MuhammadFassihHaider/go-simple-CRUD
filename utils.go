package main

func addDummyData(movies []Movie) []Movie {
	movies = append(movies, Movie{
		Id:    "1",
		ISBN:  "978-0061120084",
		Title: "To Kill a Mockingbird",
		Director: &Director{
			Firstname: "Robert",
			Lastname:  "Mulligan",
		},
	})

	movies = append(movies, Movie{
		Id:    "2",
		ISBN:  "978-0679746041",
		Title: "One Hundred Years of Solitude",
		Director: &Director{
			Firstname: "Gabriel",
			Lastname:  "García Márquez",
		},
	})

	movies = append(movies, Movie{
		Id:    "3",
		ISBN:  "978-1451673319",
		Title: "The Great Gatsby",
		Director: &Director{
			Firstname: "Baz",
			Lastname:  "Luhrmann",
		},
	})

	movies = append(movies, Movie{
		Id:    "4",
		ISBN:  "978-0143110157",
		Title: "The Catcher in the Rye",
		Director: &Director{
			Firstname: "J.D.",
			Lastname:  "Salinger",
		},
	})

	movies = append(movies, Movie{
		Id:    "5",
		ISBN:  "978-1400031702",
		Title: "Lolita",
		Director: &Director{
			Firstname: "Stanley",
			Lastname:  "Kubrick",
		},
	})

	movies = append(movies, Movie{
		Id:    "6",
		ISBN:  "978-0385333849",
		Title: "Slaughterhouse-Five",
		Director: &Director{
			Firstname: "George",
			Lastname:  "Roy Hill",
		},
	})

	movies = append(movies, Movie{
		Id:    "7",
		ISBN:  "978-0307743662",
		Title: "The Help",
		Director: &Director{
			Firstname: "Tate",
			Lastname:  "Taylor",
		},
	})

	movies = append(movies, Movie{
		Id:    "8",
		ISBN:  "978-0307388138",
		Title: "The Road",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Hillcoat",
		},
	})

	movies = append(movies, Movie{
		Id:    "9",
		ISBN:  "978-0142437247",
		Title: "Moby-Dick",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Huston",
		},
	})

	movies = append(movies, Movie{
		Id:    "10",
		ISBN:  "978-0684801469",
		Title: "The Bell Jar",
		Director: &Director{
			Firstname: "Julie",
			Lastname:  "Taymor",
		},
	})

	return movies
}
