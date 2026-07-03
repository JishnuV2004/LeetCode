func minMovesToSeat(seats []int, students []int) int {
    slices.Sort(seats)
    slices.Sort(students)
    count := 0
    for i:=0; i< len(students); i++ {
        if seats[i] > students[i] {
            count += seats[i] - students[i]
        } else {
             count += students[i] - seats[i]
        }
    }
    return count
}