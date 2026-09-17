package sprint

package main
import "fmt"

type Student struct {
	Name string
	Grades []int
}

func Mean(scores []int) float32 {
	sum := 0
	count := 0

	for _, v := range scores {
		sum += v
		count ++
	}

	return float32(sum/count)
}

func TopStudent(students []Student) Student {
	top_index := int(0)
	top_score := float32(0)
	var s Student

	if len(students) == 0 {
		return s
	}

	for i, v := range students {
		student_score := Mean(v.Grades)
		if top_score < student_score {
			top_score = student_score
			top_index = i
		}
	}

	return students[top_index]
}


func main() {
    //var s []Student
	students := []Student{
    	{Name: "Alice", Grades: []int{80, 90, 85}},
    	{Name: "Bob",   Grades: []int{95, 92, 98}},
    	{Name: "Carol", Grades: []int{70, 75, 72}},
	}
	fmt.Println(TopStudent(students))

	fmt.Println([]Student{})
}

