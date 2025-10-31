package main
import "fmt"

func main() {
	students := make(map[string]int)

	students["Alice"] = 85
	students["Bob"] = 90
	students["Charlie"] = 78
	students["Diana"] = 92
	fmt.Println("Students Map:", students)

	fmt.Println("Alice's Score:", students["Alice"])
	fmt.Println("Bob's Score:", students["Bob"])
	students["Alice"] = 88
	fmt.Println("Updated Alice's Score:", students["Alice"])


	delete(students,"Charlie")
	fmt.Println("Students Map after deletion:", students)

	fmt.Println("Socre of Nathan: ", students["Nathan"])

	Score, exists := students["Diana"]
	Score1, exists1 := students["Nathan"]

	fmt.Println("Diana's Score:", Score, "Exists:", exists)
	fmt.Println("Nathan's Score:", Score1, "Exists:", exists1)

	for name, score := range students {
		fmt.Printf("Name: %s, Score: %d\n", name, score)
	}
}